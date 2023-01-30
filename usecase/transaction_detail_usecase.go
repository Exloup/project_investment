package usecase

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/repository"
)

// Interface Usecase
type TransactionDetailUseCase interface {
	CreateTransactionDetail(transaction *model.TransactionDetails) (*model.TransactionDetails, error)
	UpdateTransactionDetail(transaction *model.TransactionDetails) (*model.TransactionDetails, error)
	GetAllTransactionDetail() ([]model.TransactionDetails, error)
	GetTransactionDetailById(id string) (model.TransactionDetails, error)
	DeleteTransactionDetail(id string) error
}

// Repo Struct
type transactiondetailUseCase struct {
	repo repository.TransactionDetailRepository
}

// Function Usecase Insert / Create
func (trd *transactiondetailUseCase) CreateTransactionDetail(transaction *model.TransactionDetails) (*model.TransactionDetails, error) {
	p, err := trd.repo.CreateTransactionDetail(transaction)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Function Usecase Get All
func (trd *transactiondetailUseCase) GetAllTransactionDetail() ([]model.TransactionDetails, error) {
	return trd.repo.GetAllTransactionDetail()
}

// Function Usecase Get By ID
func (trd *transactiondetailUseCase) GetTransactionDetailById(id string) (model.TransactionDetails, error) {
	return trd.repo.GetTransactionDetailById(id)
}

// Function Usecase Update By ID
func (trd *transactiondetailUseCase) UpdateTransactionDetail(transaction *model.TransactionDetails) (*model.TransactionDetails, error) {
	p, err := trd.repo.UpdateTransactionDetail(transaction)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Function Usecase Delete by ID
func (trd *transactiondetailUseCase) DeleteTransactionDetail(id string) error {
	return trd.repo.DeleteTransactionDetail(id)
}

// Store For Usecase Repo
func NewTransactionDetailUseCase(repoParm repository.TransactionDetailRepository) TransactionDetailUseCase {
	tdUseCase := new(transactiondetailUseCase)
	tdUseCase.repo = repoParm
	return tdUseCase
}
