package usecase

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/repository"
)

// Interface UseCase
type AccountUseCase interface {
	RegisterAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error)
	UpdateAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error)
	ActivatedAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error)
	DeactivatedAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error)
	GetAllAccount() ([]model.CustomerAccounts, error)
	GetAccountById(id string) (model.CustomerAccounts, error)
	DeleteAccount(id string) error
}

// Make Struct connected to Repo
type accountUseCase struct {
	repo repository.AccountRepository
}

// Function Usecase Register
func (ac *accountUseCase) RegisterAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error) {
	a, err := ac.repo.RegisterAccount(accounts)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Function Usecase Update
func (ac *accountUseCase) UpdateAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error) {
	a, err := ac.repo.UpdateAccount(accounts)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Function Usecase Get All
func (ac *accountUseCase) GetAllAccount() ([]model.CustomerAccounts, error) {
	return ac.repo.GetAllAccount()
}

// Function Usecase Get ID
func (ac *accountUseCase) GetAccountById(id string) (model.CustomerAccounts, error) {
	return ac.repo.GetAccountById(id)
}

// Function Usecase Activated by ID
func (ac *accountUseCase) ActivatedAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error) {
	a, err := ac.repo.ActivatedAccount(accounts)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Function Usecase Deactivate by ID
func (ac *accountUseCase) DeactivatedAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error) {
	a, err := ac.repo.DeactivatedAccount(accounts)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Function Usecase Delete
func (ac *accountUseCase) DeleteAccount(id string) error {
	return ac.repo.DeleteAccount(id)
}

// Repo for Usecase
func NewAccountUseCase(repoParm repository.AccountRepository) AccountUseCase {
	aUseCase := new(accountUseCase)
	aUseCase.repo = repoParm
	return aUseCase
}
