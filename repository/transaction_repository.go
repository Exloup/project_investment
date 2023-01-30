package repository

import (
	"database/sql"
	"final_project_coins_investment/model"
	"final_project_coins_investment/utils"
)

// Interface Transaction Product Repo
type TransactionProductRepository interface {
	CreateTransactionProduct(transaction *model.TransactionProducts) (*model.TransactionProducts, error)
	UpdateTransactionProduct(transaction *model.TransactionProducts) (*model.TransactionProducts, error)
	GetAllTransactionProduct() ([]model.TransactionProducts, error)
	GetTransactionProductById(id string) (model.TransactionProducts, error)
	DeleteTransactionProduct(id string) error
}

// Make Struct Connect to DB
type transactionproductRepository struct {
	db *sql.DB
}

// Function For Create Transaction Product
func (tp *transactionproductRepository) CreateTransactionProduct(transaction *model.TransactionProducts) (*model.TransactionProducts, error) {
	_, err := tp.db.Exec(utils.INSERT_TRX_PRODUCT, transaction.TrxCode, transaction.DateTransaction, transaction.PortofolioId, transaction.ProductId, transaction.StatusDescription)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// Function For List Transaction Product
func (tp *transactionproductRepository) GetAllTransactionProduct() ([]model.TransactionProducts, error) {
	rows, err := tp.db.Query(utils.VIEW_TRX_PRODUCT)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var listtrxs []model.TransactionProducts
	for rows.Next() {
		var listtrx model.TransactionProducts
		err := rows.Scan(&listtrx.Id, &listtrx.TrxCode, &listtrx.DateTransaction, &listtrx.PortofolioId, &listtrx.ProductId, &listtrx.StatusDescription)
		if err != nil {
			panic(err)
		}
		listtrxs = append(listtrxs, listtrx)
	}
	return listtrxs, nil
}

// Function For List Transaction Product by ID
func (tp *transactionproductRepository) GetTransactionProductById(id string) (model.TransactionProducts, error) {
	var trx model.TransactionProducts
	err := tp.db.QueryRow(utils.SELECT_TRX_PRODUCT_IDS, id).Scan(&trx.Id, &trx.TrxCode, &trx.DateTransaction, &trx.PortofolioId, &trx.ProductId, &trx.StatusDescription)
	if err != nil {
		return model.TransactionProducts{}, err
	}
	return trx, nil
}

// Function For Update Transaction Product by ID
func (tp *transactionproductRepository) UpdateTransactionProduct(transaction *model.TransactionProducts) (*model.TransactionProducts, error) {
	_, err := tp.db.Exec(utils.UPDATE_TRX_PRODUCT_IDS, transaction.TrxCode, transaction.PortofolioId, transaction.ProductId, transaction.StatusDescription, transaction.Id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// Function For Delete Transaction Product by ID
func (tp *transactionproductRepository) DeleteTransactionProduct(id string) error {
	_, err := tp.db.Exec(utils.DELETE_TRX_PRODUCT, id)
	if err != nil {
		return err
	}
	return nil
}

// Create Repo for product Customers
func NewTransactionProductRepository(db *sql.DB) TransactionProductRepository {
	transactionRepo := new(transactionproductRepository)
	transactionRepo.db = db
	return transactionRepo
}
