package usecase

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/repository"
)

// Interface Usecase
type PortofolioUseCase interface {
	CreatePortofolio(portofolios *model.Portofolios) (*model.Portofolios, error)
	UpdatePortofolio(portofolios *model.Portofolios) (*model.Portofolios, error)
	GetAllPortofolio() ([]model.Portofolios, error)
	GetPortofolioById(id string) (model.Portofolios, error)
	DeletePortofolio(id string) error
}

// Repo Struct
type portofolioUsecase struct {
	repo repository.PortofolioRepository
}

// Function Usecase Insert / Create
func (por *portofolioUsecase) CreatePortofolio(portofolios *model.Portofolios) (*model.Portofolios, error) {
	p, err := por.repo.CreatePortofolio(portofolios)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Function Usecase Get All
func (por *portofolioUsecase) GetAllPortofolio() ([]model.Portofolios, error) {
	return por.repo.GetAllPortofolio()
}

// Function Usecase Get By ID
func (por *portofolioUsecase) GetPortofolioById(id string) (model.Portofolios, error) {
	return por.repo.GetPortofolioById(id)
}

// Function Usecase Update By ID
func (por *portofolioUsecase) UpdatePortofolio(portofolios *model.Portofolios) (*model.Portofolios, error) {
	p, err := por.repo.UpdatePortofolio(portofolios)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Function Usecase Delete by ID
func (por *portofolioUsecase) DeletePortofolio(id string) error {
	return por.repo.DeletePortofolio(id)
}

// Store For Usecase Repo
func NewPortofolioUseCase(repoParm repository.PortofolioRepository) PortofolioUseCase {
	porUseCase := new(portofolioUsecase)
	porUseCase.repo = repoParm
	return porUseCase
}
