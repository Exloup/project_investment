package manager

import (
	"final_project_coins_investment/repository"
)

type RepositoryManager interface {
	AccountRepo() repository.AccountRepository
	ProductRepo() repository.ProductRepository
	PriceRepo() repository.PriceRepository
	PortofolioRepo() repository.PortofolioRepository
	TransactionProductRepo() repository.TransactionProductRepository
	TransactionDetailRepo() repository.TransactionDetailRepository
	// AuthRepo() repository.
}

type repositoryManager struct {
	admin AdminManager
}

func (rm *repositoryManager) AccountRepo() repository.AccountRepository {
	return repository.NewAccountRepository(rm.admin.DbConn())
}

func (rm *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(rm.admin.DbConn())
}

func (rm *repositoryManager) PriceRepo() repository.PriceRepository {
	return repository.NewPriceRepository(rm.admin.DbConn())
}
func (rm *repositoryManager) PortofolioRepo() repository.PortofolioRepository {
	return repository.NewPortofolioRepository(rm.admin.DbConn())
}
func (rm *repositoryManager) TransactionProductRepo() repository.TransactionProductRepository {
	return repository.NewTransactionProductRepository(rm.admin.DbConn())
}
func (rm *repositoryManager) TransactionDetailRepo() repository.TransactionDetailRepository {
	return repository.NewTransactionDetailRepository(rm.admin.DbConn())
}

func NewRepositoryManager(admin AdminManager) RepositoryManager {
	return &repositoryManager{admin: admin}
}
