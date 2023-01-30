package usecase

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/repository"
)

// Interface Usecase
type TransactionProductUseCase interface {
	CreateTransactionProduct(transaction *model.TransactionProducts) (*model.TransactionProducts, error)
	UpdateTransactionProduct(transaction *model.TransactionProducts) (*model.TransactionProducts, error)
	GetAllTransactionProduct() ([]model.TransactionProducts, error)
	GetTransactionProductById(id string) (model.TransactionProducts, error)
	DeleteTransactionProduct(id string) error
}

// Repo Struct
type transactionproductUseCase struct {
	repo repository.TransactionProductRepository
}

// Function Usecase Insert / Create
func (trp *transactionproductUseCase) CreateTransactionProduct(transaction *model.TransactionProducts) (*model.TransactionProducts, error) {
	p, err := trp.repo.CreateTransactionProduct(transaction)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Function Usecase Get All
func (trp *transactionproductUseCase) GetAllTransactionProduct() ([]model.TransactionProducts, error) {
	return trp.repo.GetAllTransactionProduct()
}

// Function Usecase Get By ID
func (trp *transactionproductUseCase) GetTransactionProductById(id string) (model.TransactionProducts, error) {
	return trp.repo.GetTransactionProductById(id)
}

// Function Usecase Update By ID
func (trp *transactionproductUseCase) UpdateTransactionProduct(transaction *model.TransactionProducts) (*model.TransactionProducts, error) {
	p, err := trp.repo.UpdateTransactionProduct(transaction)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Function Usecase Delete by ID
func (trp *transactionproductUseCase) DeleteTransactionProduct(id string) error {
	return trp.repo.DeleteTransactionProduct(id)
}

// Store For Usecase Repo
func NewTransactionProductUseCase(repoParm repository.TransactionProductRepository) TransactionProductUseCase {
	trpUseCase := new(transactionproductUseCase)
	trpUseCase.repo = repoParm
	return trpUseCase
}
