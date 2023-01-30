package manager

import (
	"database/sql"
	"final_project_coins_investment/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type AdminManager interface {
	DbConn() *sql.DB
}

type adminManager struct {
	db   *sql.DB
	conf config.Config
}

// Initial DB
func (am *adminManager) initDB() {
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", am.conf.Host, am.conf.User, am.conf.Password, am.conf.Name, am.conf.Port)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application Failed to Run", err)
		}
	}()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	am.db = db
	fmt.Println("DB is Connected")
}

// Connection DB
func (am *adminManager) DbConn() *sql.DB {
	return am.db
}

// function for Repo
func NewAdminManager(configParam config.Config) AdminManager {
	admin := adminManager{
		conf: configParam,
	}
	admin.initDB()
	return &admin
}
