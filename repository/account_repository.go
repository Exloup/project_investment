package repository

import (
	"database/sql"
	"final_project_coins_investment/model"
	"final_project_coins_investment/utils"
	"log"
)

// Interface Account repository
type AccountRepository interface {
	RegisterAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error)
	UpdateAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error)
	ActivatedAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error)
	DeactivatedAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error)
	GetAllAccount() ([]model.CustomerAccounts, error)
	GetAccountById(id string) (model.CustomerAccounts, error)
	DeleteAccount(id string) error
}

// Make Struct connected to DB
type accountRepository struct {
	db *sql.DB
}

// Function Register
func (ar *accountRepository) RegisterAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error) {
	_, err := ar.db.Exec(utils.CREATE_ACCOUNTS, accounts.Email, accounts.Password, accounts.Name, accounts.Phone)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	return accounts, nil
}

// Function List all
func (ar *accountRepository) GetAllAccount() ([]model.CustomerAccounts, error) {
	rows, err := ar.db.Query(utils.VIEW_ACCOUNTS)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var listaccounts []model.CustomerAccounts
	for rows.Next() {
		var listaccount model.CustomerAccounts
		err := rows.Scan(&listaccount.Id, &listaccount.Email, &listaccount.Password, &listaccount.Name, &listaccount.Phone, &listaccount.CreateAt, &listaccount.Active)
		if err != nil {
			panic(err)
		}
		listaccounts = append(listaccounts, listaccount)
	}
	return listaccounts, nil
}

// Function Get ID
func (ar *accountRepository) GetAccountById(id string) (model.CustomerAccounts, error) {
	var account model.CustomerAccounts
	err := ar.db.QueryRow(utils.SELECT_ACCOUNT_IDS, id).Scan(&account.Id, &account.Email, &account.Password, &account.Name, &account.Phone, &account.CreateAt, &account.CreateAt)
	if err != nil {
		return model.CustomerAccounts{}, err
	}
	return account, nil
}

// Function Update Data by ID
func (ar *accountRepository) UpdateAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error) {
	_, err := ar.db.Exec(utils.UPDATE_ACCOUNT_IDS, accounts.Name, accounts.Phone, accounts.Id)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// Function Status Account by ID
// Activated
func (ar *accountRepository) ActivatedAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error) {
	_, err := ar.db.Exec(utils.ACTIVE_ACCOUNT_IDS, accounts.Id)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// Deactivated
func (ar *accountRepository) DeactivatedAccount(accounts *model.CustomerAccounts) (*model.CustomerAccounts, error) {
	_, err := ar.db.Exec(utils.DEACTIVE_ACCOUNT_IDS, accounts.Id)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// Function Deleted Account by ID
func (ar *accountRepository) DeleteAccount(id string) error {
	_, err := ar.db.Exec(utils.DELETE_ACCOUNT, id)
	if err != nil {
		return err
	}
	return nil
}

// Create Repo for Account Customers
func NewAccountRepository(db *sql.DB) AccountRepository {
	accountRepo := new(accountRepository)
	accountRepo.db = db
	return accountRepo
}
