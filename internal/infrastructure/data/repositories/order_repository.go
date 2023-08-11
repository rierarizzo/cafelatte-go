package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/pkg/constants"
	"github.com/rierarizzo/cafelatte/pkg/params"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type OrderRepository struct {
	db *sqlx.DB
}

func (r *OrderRepository) InsertPurchaseOrder(order entities.PurchaseOrder) (int, *domain.AppError) {
	rollbackTxAndReturnZeroAndErr := func(tx *sqlx.Tx,
		err error) (int, *domain.AppError) {
		log := logrus.WithField(constants.RequestIDKey, params.RequestID())

		_ = tx.Rollback()
		log.Error(err)
		return 0, domain.NewAppError(err, domain.RepositoryError)
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return rollbackTxAndReturnZeroAndErr(tx, err)
	}

	orderModel := mappers.OrderToModel(order)

	result, err := tx.Exec(`insert into PurchaseOrder (
                           UserID, 
                           ShippingAddressID, 
                           PaymentMethodID, 
                           Notes,  
                           OrderedAt) values (?,?,?,?,?)`, orderModel.UserID,
		orderModel.ShippingAddressID, orderModel.PaymentMethodID,
		orderModel.Notes.String, time.Now())
	if err != nil {
		return rollbackTxAndReturnZeroAndErr(tx, err)
	}

	orderID, _ := result.LastInsertId()
	insertProductStmnt, err := tx.Prepare(`insert into PurchasedProduct (
                              OrderID, 
                              ProductID, 
                              Quantity) values (?,?,?)`)
	if err != nil {
		return rollbackTxAndReturnZeroAndErr(tx, err)
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
		return rollbackTxAndReturnZeroAndErr(tx, err)
	}

	var query = `select sum(pp.Quantity * p.Price) from PurchasedProduct pp 
				inner join Product p on pp.ProductID = p.ID where OrderID=?`

	var totalAmount float64
	err = tx.Get(&totalAmount, query, orderID)
	if err != nil {
		return rollbackTxAndReturnZeroAndErr(tx, err)
	}

	_, err = tx.Exec("update PurchaseOrder set TotalAmount=? where ID=?",
		totalAmount, orderID)
	if err != nil {
		return rollbackTxAndReturnZeroAndErr(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return rollbackTxAndReturnZeroAndErr(tx, err)
	}

	return int(orderID), nil
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db}
}
