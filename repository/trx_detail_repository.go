package repository

import (
	"database/sql"
	"final_project_coins_investment/model"
	"final_project_coins_investment/utils"
)

// Interface Transaction Product Repo
type TransactionDetailRepository interface {
	CreateTransactionDetail(transaction *model.TransactionDetails) (*model.TransactionDetails, error)
	UpdateTransactionDetail(transaction *model.TransactionDetails) (*model.TransactionDetails, error)
	GetAllTransactionDetail() ([]model.TransactionDetails, error)
	GetTransactionDetailById(id string) (model.TransactionDetails, error)
	DeleteTransactionDetail(id string) error
}

// Make Struct Connect to DB
type transactiondetailRepository struct {
	db *sql.DB
}

// Function For Create Transaction Product
func (tx *transactiondetailRepository) CreateTransactionDetail(trx *model.TransactionDetails) (*model.TransactionDetails, error) {
	_, err := tx.db.Exec(utils.INSERT_TRX_DETAIL, trx.Id,trx.PriceId, trx.TrxProductId, trx.Quantity)
	if err != nil {
		return nil, err
	}
	return trx, nil
}

// Function For List Transaction Product
func (tx *transactiondetailRepository) GetAllTransactionDetail() ([]model.TransactionDetails, error) {
	rows, err := tx.db.Query(utils.VIEW_TRX_DETAIL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var listtrxs []model.TransactionDetails
	for rows.Next() {
		var listtrx model.TransactionDetails
		err := rows.Scan(&listtrx.Id, &listtrx.PriceId, &listtrx.TrxProductId, &listtrx.Quantity)
		if err != nil {
			panic(err)
		}
		listtrxs = append(listtrxs, listtrx)
	}
	return listtrxs, nil
}

// Function For List Transaction Product by ID
func (tx *transactiondetailRepository) GetTransactionDetailById(id string) (model.TransactionDetails, error) {
	var trd model.TransactionDetails
	err := tx.db.QueryRow(utils.SELECT_TRX_DETAIL_IDS, id).Scan(&trd.Id, &trd.PriceId, &trd.TrxProductId, &trd.Quantity)
	if err != nil {
		return model.TransactionDetails{}, err
	}
	return trd, nil
}

// Function For Update Transaction Product by ID
func (tx *transactiondetailRepository) UpdateTransactionDetail(trx *model.TransactionDetails) (*model.TransactionDetails, error) {
	_, err := tx.db.Exec(utils.UPDATE_TRX_DETAIL_IDS, trx.PriceId, trx.TrxProductId, trx.Quantity, trx.Id)
	if err != nil {
		return nil, err
	}
	return trx, nil
}

// Function For Delete Transaction Product by ID
func (tx *transactiondetailRepository) DeleteTransactionDetail(id string) error {
	_, err := tx.db.Exec(utils.DELETE_TRX_DETAIL, id)
	if err != nil {
		return err
	}
	return nil
}

// Create Repo for product Customers
func NewTransactionDetailRepository(db *sql.DB) TransactionDetailRepository {
	trxRepo := new(transactiondetailRepository)
	trxRepo.db = db
	return trxRepo
}
