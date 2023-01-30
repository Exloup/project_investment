package manager

import "final_project_coins_investment/usecase"

type UseCaseManager interface {
	AccountUseCase() usecase.AccountUseCase
	ProductUseCase() usecase.ProductUseCase
	PriceUseCase() usecase.PriceUseCase
	PortofolioUseCase() usecase.PortofolioUseCase
	TransactionProductUseCase() usecase.TransactionProductUseCase
	TransactionDetailUseCase() usecase.TransactionDetailUseCase
	// UserAuthenUseCase() usecase.AuthUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (um *useCaseManager) AccountUseCase() usecase.AccountUseCase {
	return usecase.NewAccountUseCase(um.repoManager.AccountRepo())
}

func (um *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(um.repoManager.ProductRepo())
}

func (um *useCaseManager) PriceUseCase() usecase.PriceUseCase {
	return usecase.NewPriceUseCase(um.repoManager.PriceRepo())
}
func (um *useCaseManager) PortofolioUseCase() usecase.PortofolioUseCase {
	return usecase.NewPortofolioUseCase(um.repoManager.PortofolioRepo())
}
func (um *useCaseManager) TransactionProductUseCase() usecase.TransactionProductUseCase {
	return usecase.NewTransactionProductUseCase(um.repoManager.TransactionProductRepo())
}
func (um *useCaseManager) TransactionDetailUseCase() usecase.TransactionDetailUseCase {
	return usecase.NewTransactionDetailUseCase(um.repoManager.TransactionDetailRepo())
}
// func (um *useCaseManager)UserAuthenUseCase() usecase.AuthUseCase{
// 	return usecase.NewAuthUseCase(um.repoManager.)
// }
func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
