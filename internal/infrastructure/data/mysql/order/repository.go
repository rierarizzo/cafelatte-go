package order

import (
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db *sqlx.DB
}

func (repository *Repository) InsertPurchaseOrder(order domain.Order) (
	int,
	*domain.AppError,
) {
	tx, _ := repository.db.Beginx()

	orderId, appErr := generatePurchaseOrderId(tx, order)
	if appErr != nil {
		return 0, appErr
	}

	query := "insert into ProductInOrder (OrderId, ProductId, Quantity) values (?,?,?)"

	insertProductStmnt, err := tx.Prepare(query)
	if err != nil {
		return 0, rollbackAndError(tx, err)
	}

	var sem = make(chan struct{}, 5)
	var errCh = make(chan error, len(order.Products))
	var wg sync.WaitGroup

	for _, v := range order.Products {
		wg.Add(1)
		sem <- struct{}{}

		go func(entity domain.ProductInOrder) {
			defer func() {
				wg.Done()
				<-sem
			}()

			product := fromProductInOrderToModel(entity)
			_, err = insertProductStmnt.Exec(orderId, product.ProductId,
				product.Quantity)
			if err != nil {
				errCh <- err
				return
			}
		}(v)
	}

	wg.Wait()
	close(errCh)
	for err = range errCh {
		return 0, rollbackAndError(tx, err)
	}

	appErr = updateOrderAmount(tx, orderId)
	if appErr != nil {
		return 0, appErr
	}

	return orderId, nil
}

func generatePurchaseOrderId(tx *sqlx.Tx, order domain.Order) (
	int,
	*domain.AppError,
) {
	model := fromOrderToModel(order)
	query := `insert into PurchaseOrder (UserId, ShippingAddressId, PaymentMethodId, 
        Notes, OrderedAt) values (?,?,?,?,?)`

	res, err := tx.Exec(query, model.UserId, model.ShippingAddressId,
		model.PaymentMethodId, model.Notes.String, time.Now())
	if err != nil {
		return 0, rollbackAndError(tx, err)
	}
	orderId, _ := res.LastInsertId()

	return int(orderId), nil
}

func updateOrderAmount(tx *sqlx.Tx, orderId int) *domain.AppError {
	var total float64

	query := `select sum(pp.Quantity * p.Price) from ProductInOrder pp 
    	inner join Product p on pp.ProductId = p.Id where OrderId=?`
	err := tx.Get(&total, query, orderId)
	if err != nil {
		return rollbackAndError(tx, err)
	}

	query = "update PurchaseOrder set TotalAmount=? where Id=?"
	_, err = tx.Exec(query, total, orderId)
	if err != nil {
		return rollbackAndError(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return rollbackAndError(tx, err)
	}

	return nil
}

func rollbackAndError(tx *sqlx.Tx, err error) *domain.AppError {
	_ = tx.Rollback()
	logrus.WithField(misc.RequestIdKey, request.Id()).Error(err)

	return domain.NewAppError(err, domain.RepositoryError)
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}
