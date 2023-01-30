package usecase

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/repository"
)

// Interface UseCase
type ProductUseCase interface {
	InsertProduct(products *model.Products) (*model.Products, error)
	UpdateProduct(products *model.Products) (*model.Products, error)
	GetAllProduct() ([]model.Products, error)
	GetProductById(id string) (model.Products, error)
	DeleteProduct(id string) error
}

// Repo Struct
type productUseCase struct {
	repo repository.ProductRepository
}

// Function Usecase Insert
func (pu *productUseCase) InsertProduct(products *model.Products) (*model.Products, error) {
	p, err := pu.repo.InsertProduct(products)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Function Usecase Update by ID
func (pu *productUseCase) UpdateProduct(products *model.Products) (*model.Products, error) {
	p, err := pu.repo.UpdateProduct(products)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Function Usecase Get All
func (pu *productUseCase) GetAllProduct() ([]model.Products, error) {
	return pu.repo.GetAllProduct()
}

// Function Usecase Get by ID
func (pu *productUseCase) GetProductById(id string) (model.Products, error) {
	return pu.repo.GetProductById(id)
}

// Function Usecase Delete by ID
func (pu *productUseCase) DeleteProduct(id string) error {
	return pu.repo.DeleteProduct(id)
}

// Store for Usecase repo
func NewProductUseCase(repoParm repository.ProductRepository) ProductUseCase {
	pUseCase := new(productUseCase)
	pUseCase.repo = repoParm
	return pUseCase
}
