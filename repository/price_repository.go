package repository

import (
	"database/sql"
	"final_project_coins_investment/model"
	"final_project_coins_investment/utils"
)

// Interfacce Price Repo
type PriceRepository interface {
	InsertPrice(prices *model.Prices) (*model.Prices, error)
	UpdatePrice(prices *model.Prices) (*model.Prices, error)
	GetAllPrice() ([]model.Prices, error)
	GetPriceById(id string) (model.Prices, error)
	DeletePrice(id string) error
}

// Struct Connection to DB
type priceRepository struct {
	db *sql.DB
}

// Function For Insert Price
func (pp *priceRepository) InsertPrice(prices *model.Prices) (*model.Prices, error) {
	_, err := pp.db.Exec(utils.INSERT_PRICE, prices.Price, prices.Date, prices.ProductId, prices.Performa)
	if err != nil {
		return nil, err
	}
	return prices, nil
}

// Function List All Price
func (pp *priceRepository) GetAllPrice() ([]model.Prices, error) {
	rows, err := pp.db.Query(utils.VIEW_PRICE)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var listprices []model.Prices
	for rows.Next() {
		var listprice model.Prices
		err := rows.Scan(&listprice.Id, &listprice.Price, &listprice.Date, &listprice.ProductId, &listprice.Performa)
		if err != nil {
			panic(err)
		}
		listprices = append(listprices, listprice)
	}
	return listprices, nil
}

// Function List Price by ID
func (pp *priceRepository) GetPriceById(id string) (model.Prices, error) {
	var pri model.Prices
	err := pp.db.QueryRow(utils.SELECT_PRICE_IDS, id).Scan(&pri.Id, &pri.Price, &pri.Date, &pri.ProductId, &pri.Performa)
	if err != nil {
		return model.Prices{}, err
	}
	return pri, nil
}

// Function Update Price By ID
func (pp *priceRepository) UpdatePrice(prices *model.Prices) (*model.Prices, error) {
	_, err := pp.db.Exec(utils.UPDATE_PRICE_IDS, prices.Price, prices.Performa, prices.Id)
	if err != nil {
		return nil, err
	}
	return prices, nil
}

// Function Delete Price By ID
func (pp *priceRepository) DeletePrice(id string) error {
	_, err := pp.db.Exec(utils.DELETE_PRICE, id)
	if err != nil {
		return err
	}
	return nil
}

// Store Repo
func NewPriceRepository(db *sql.DB) PriceRepository {
	priceRepo := new(priceRepository)
	priceRepo.db = db
	return priceRepo
}
