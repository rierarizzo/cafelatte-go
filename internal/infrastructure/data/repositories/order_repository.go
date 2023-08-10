package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/constants"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type OrderRepository struct {
	db *sqlx.DB
}

func (r *OrderRepository) InsertPurchaseOrder(order entities.PurchaseOrder) (int, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	rollbackTxAndReturnErr := func(tx *sqlx.Tx,
		err error) (int, *domain.AppError) {
		_ = tx.Rollback()
		log.Error(err)
		return 0, domain.NewAppError(err, domain.RepositoryError)
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return rollbackTxAndReturnErr(tx, err)
	}

	orderModel := mappers.OrderToModel(order)

	// todo: check why the "notes" field is not saved
	result, err := tx.Exec(`insert into PurchaseOrder (
                           UserID, 
                           ShippingAddressID, 
                           PaymentMethodID, 
                           Notes,  
                           OrderedAt) VALUES (?,?,?,?,?)`, orderModel.UserID,
		orderModel.ShippingAddressID, orderModel.PaymentMethodID,
		orderModel.Notes, time.Now())
	if err != nil {
		return rollbackTxAndReturnErr(tx, err)
	}

	orderID, _ := result.LastInsertId()
	insertProductStmnt, err := tx.Prepare(`insert into ProductPurchased (
                              OrderID, 
                              ProductID, 
                              Quantity) values (?,?,?)`)
	if err != nil {
		return rollbackTxAndReturnErr(tx, err)
	}

	var sem = make(chan struct{}, 5)
	var errCh = make(chan error, len(order.PurchasedProducts))
	var wg sync.WaitGroup

	for _, v := range order.PurchasedProducts {
		wg.Add(1)
		sem <- struct{}{}

		go func(entity entities.PurchasedProduct) {
			defer func() {
				wg.Done()
				<-sem
			}()

			product := mappers.PurchasedProductToModel(entity)

			_, err := insertProductStmnt.Exec(orderID, product.ProductID,
				product.Quantity)
			if err != nil {
				errCh <- err
				return
			}

		}(v)
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		return rollbackTxAndReturnErr(tx, err)
	}

	// todo: update the "TotalAmount" field

	err = tx.Commit()
	if err != nil {
		return rollbackTxAndReturnErr(tx, err)
	}

	return int(orderID), nil
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db}
}
