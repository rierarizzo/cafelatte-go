package order

import (
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Repository struct {
	db *sqlx.DB
}

func (r *Repository) InsertPurchaseOrder(order domain.Order) (int, *domain.AppError) {
	rollbackAndError := func(tx *sqlx.Tx, err error) *domain.AppError {
		_ = tx.Rollback()
		logrus.WithField(misc.RequestIDKey, request.ID()).Error(err)

		return domain.NewAppError(err, domain.RepositoryError)
	}

	tx, _ := r.db.Beginx()

	model := fromOrderToModel(order)
	query := `insert into PurchaseOrder (UserID, ShippingAddressID, PaymentMethodID, 
        Notes, OrderedAt) values (?,?,?,?,?)`

	res, err := tx.Exec(query, model.UserID, model.ShippingAddressID,
		model.PaymentMethodID, model.Notes.String, time.Now())
	if err != nil {
		return 0, rollbackAndError(tx, err)
	}
	orderID, _ := res.LastInsertId()

	query = "insert into ProductInOrder (OrderID, ProductID, Quantity) values (?,?,?)"

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
		return 0, rollbackAndError(tx, err)
	}

	err = updateOrderAmount(tx, int(orderID))
	if err != nil {
		return 0, rollbackAndError(tx, err)
	}

	return int(orderID), nil
}

func updateOrderAmount(tx *sqlx.Tx, orderID int) error {
	var total float64

	query := `select sum(pp.Quantity * p.Price) from ProductInOrder pp inner
				join Product p on pp.ProductID = p.ID where OrderID=?`
	err := tx.Get(&total, query, orderID)
	if err != nil {
		return err
	}

	query = "update PurchaseOrder set TotalAmount=? where ID=?"
	_, err = tx.Exec(query, total, orderID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func NewOrderRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}
