package repository

import (
	"database/sql"
	"final_project_coins_investment/model"
	"final_project_coins_investment/utils"
)

// Interface Portofolio Repository
type PortofolioRepository interface {
	CreatePortofolio(portofolios *model.Portofolios) (*model.Portofolios, error)
	UpdatePortofolio(portofolios *model.Portofolios) (*model.Portofolios, error)
	GetAllPortofolio() ([]model.Portofolios, error)
	GetPortofolioById(id string) (model.Portofolios, error)
	DeletePortofolio(id string) error
}

// Make Struct connected to DB
type portofolioRepository struct {
	db *sql.DB
}

// Function for Create Portofolio
func (po *portofolioRepository) CreatePortofolio(portofolios *model.Portofolios) (*model.Portofolios, error) {
	_, err := po.db.Exec(utils.INSERT_PORTOFOLIO,portofolios.AccountId, portofolios.TotalQuantity, portofolios.Profit)
	if err != nil {
		return nil, err
	}
	return portofolios, nil
}

// Function for List All
func (po *portofolioRepository) GetAllPortofolio() ([]model.Portofolios, error) {
	rows, err := po.db.Query(utils.VIEW_PORTOFOLIO)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var listportos []model.Portofolios
	for rows.Next() {
		var listporto model.Portofolios
		err := rows.Scan(&listporto.Id, &listporto.AccountId, &listporto.TotalQuantity, &listporto.Profit)
		if err != nil {
			panic(err)
		}
		listportos = append(listportos, listporto)
	}
	return listportos, nil
}

// Function for List by ID
func (po *portofolioRepository) GetPortofolioById(id string) (model.Portofolios, error) {
	var por model.Portofolios
	err := po.db.QueryRow(utils.SELECT_PORTOFOLIO_IDS, id).Scan(&por.Id, &por.AccountId, &por.TotalQuantity, &por.Profit)
	if err != nil {
		return model.Portofolios{}, err
	}
	return por, nil
}

// Function for Update Portofolio
func (po *portofolioRepository) UpdatePortofolio(portofolios *model.Portofolios) (*model.Portofolios, error) {
	_, err := po.db.Exec(utils.UPDATE_PORTOFOLIO_IDS, portofolios.Id, portofolios.AccountId, portofolios.TotalQuantity, portofolios.Profit)
	if err != nil {
		return nil, err
	}
	return portofolios, nil
}

// Function for Delete Portofolio
func (po *portofolioRepository) DeletePortofolio(id string) error {
	_, err := po.db.Exec(utils.DELETE_PORTOFOLIO, id)
	if err != nil {
		return err
	}
	return nil
}

// Create Repo for Portofolio
func NewPortofolioRepository(db *sql.DB) PortofolioRepository {
	portoRepo := new(portofolioRepository)
	portoRepo.db = db
	return portoRepo
}
