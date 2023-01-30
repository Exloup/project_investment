package repository

import (
	"database/sql"
	"final_project_coins_investment/model"
	"final_project_coins_investment/utils"
)

// Interface Account repository
type ProductRepository interface {
	InsertProduct(products *model.Products) (*model.Products, error)
	UpdateProduct(products *model.Products) (*model.Products, error)
	GetAllProduct() ([]model.Products, error)
	GetProductById(id string) (model.Products, error)
	DeleteProduct(id string) error
}

// Make Struct connected to DB
type productRepository struct {
	db *sql.DB
}

// Function for Insert Product
func (pr *productRepository) InsertProduct(products *model.Products) (*model.Products, error) {
	_, err := pr.db.Exec(utils.INSERT_PRODUCT, products.ProductName, products.ProductDescription)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Function List all
func (pr *productRepository) GetAllProduct() ([]model.Products, error) {
	rows, err := pr.db.Query(utils.VIEW_PRODUCT)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var listproducts []model.Products
	for rows.Next() {
		var listproduct model.Products
		err := rows.Scan(&listproduct.Id, &listproduct.ProductName, &listproduct.ProductDescription)
		if err != nil {
			panic(err)
		}
		listproducts = append(listproducts, listproduct)
	}
	return listproducts, nil
}

// Function Get ID
func (pr *productRepository) GetProductById(id string) (model.Products, error) {
	var pro model.Products
	err := pr.db.QueryRow(utils.SELECT_PRODUCT_IDS, id).Scan(&pro.Id, &pro.ProductName, &pro.ProductDescription)
	if err != nil {
		return model.Products{}, err
	}
	return pro, nil
}

// Function Update Data by ID
func (pr *productRepository) UpdateProduct(products *model.Products) (*model.Products, error) {
	_, err := pr.db.Exec(utils.UPDATE_PRODUCT_IDS, products.ProductName, products.ProductDescription, products.Id)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Function Deleted Product by ID
func (pr *productRepository) DeleteProduct(id string) error {
	_, err := pr.db.Exec(utils.DELETE_PRODUCT, id)
	if err != nil {
		return err
	}
	return nil
}

// Create Repo for product
func NewProductRepository(db *sql.DB) ProductRepository {
	productRepo := new(productRepository)
	productRepo.db = db
	return productRepo
}
