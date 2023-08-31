package order

import (
	sqlUtil "github.com/rierarizzo/cafelatte/pkg/utils/sql"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) InsertPurchaseOrder(order domain.Order) (int, *domain.AppError) {
	tx, appErr := sqlUtil.StartTransaction(r.db)
	if appErr != nil {
		return 0, appErr
	}

	defer sqlUtil.RollbackIfPanic(tx)

	orderId, appErr := generatePurchaseOrderId(tx, order)
	if appErr != nil {
		return 0, appErr
	}

	query := `
		INSERT INTO ProductInOrder (OrderId, ProductId, Quantity) VALUES (?,?,?)
	`
	insertProductStmnt, err := tx.Prepare(query)
	if err != nil {
		appErr = domain.NewAppError(err, domain.RepositoryError)
		return 0, appErr
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
			_, err = insertProductStmnt.Exec(orderId,
				product.ProductId,
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
		appErr = domain.NewAppError(err, domain.RepositoryError)
		return 0, appErr
	}

	if appErr = updateOrderAmount(tx, orderId); appErr != nil {
		return 0, appErr
	}

	if appErr = sqlUtil.CommitTransaction(tx); appErr != nil {
		return 0, appErr
	}

	return orderId, nil
}

func generatePurchaseOrderId(tx *sqlx.Tx,
	order domain.Order) (int, *domain.AppError) {
	model := fromOrderToModel(order)

	query := `
		INSERT INTO PurchaseOrder 
		    (UserId, ShippingAddressId, PaymentMethodId, Notes, OrderedAt) 
		VALUES (?,?,?,?,?)
	`
	result, appErr := sqlUtil.ExecWithTransaction(tx,
		query,
		model.UserId,
		model.ShippingAddressId,
		model.PaymentMethodId,
		model.Notes.String,
		time.Now())
	if appErr != nil {
		return 0, appErr
	}

	orderId, appErr := sqlUtil.GetLastInsertedId(result)
	if appErr != nil {
		return 0, appErr
	}

	return orderId, nil
}

func updateOrderAmount(tx *sqlx.Tx, orderId int) *domain.AppError {
	var total float64

	query := `
		SELECT SUM(pp.Quantity * p.Price) FROM ProductInOrder pp 
    	INNER JOIN Product p ON pp.ProductId = p.Id WHERE OrderId=?
	`
	err := tx.Get(&total, query, orderId)
	if err != nil {
		appErr := domain.NewAppError(err, domain.RepositoryError)
		return appErr
	}

	query = `
		UPDATE PurchaseOrder SET TotalAmount=? WHERE Id=?
	`
	_, err = tx.Exec(query, total, orderId)
	if err != nil {
		appErr := domain.NewAppError(err, domain.RepositoryError)
		return appErr
	}

	return nil
}
