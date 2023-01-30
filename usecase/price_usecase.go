package usecase

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/repository"
)

// Interface Usecase
type PriceUseCase interface {
	InsertPrice(prices *model.Prices) (*model.Prices, error)
	UpdatePrice(prices *model.Prices) (*model.Prices, error)
	GetAllPrice() ([]model.Prices, error)
	GetPriceById(id string) (model.Prices, error)
	DeletePrice(id string) error
}

// Repo Struct
type priceUseCase struct {
	repo repository.PriceRepository
}

// Function Usecase Insert / Create
func (pru *priceUseCase) InsertPrice(prices *model.Prices) (*model.Prices, error) {
	p, err := pru.repo.InsertPrice(prices)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Function Usecase Get All
func (pru *priceUseCase) GetAllPrice() ([]model.Prices, error) {
	return pru.repo.GetAllPrice()
}

// Function Usecase Get By ID
func (pru *priceUseCase) GetPriceById(id string) (model.Prices, error) {
	return pru.repo.GetPriceById(id)
}

// Function Usecase Update By ID
func (pru *priceUseCase) UpdatePrice(prices *model.Prices) (*model.Prices, error) {
	p, err := pru.repo.UpdatePrice(prices)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Function Usecase Delete by ID
func (pru *priceUseCase) DeletePrice(id string) error {
	return pru.repo.DeletePrice(id)
}

// Store For Usecase Repo
func NewPriceUseCase(repoParm repository.PriceRepository) PriceUseCase {
	prUseCase := new(priceUseCase)
	prUseCase.repo = repoParm
	return prUseCase
}
