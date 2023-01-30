package delivery

import (
	"final_project_coins_investment/config"
	"final_project_coins_investment/delivery/controller"
	"final_project_coins_investment/delivery/middleware"
	"final_project_coins_investment/manager"
	"final_project_coins_investment/usecase"
	"final_project_coins_investment/utils/authenticator"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Struct for app server
type appServer struct {
	acases       usecase.AccountUseCase
	pcase        usecase.ProductUseCase
	prcase       usecase.PriceUseCase
	porcase      usecase.PortofolioUseCase
	trcase       usecase.TransactionProductUseCase
	trxcase      usecase.TransactionDetailUseCase
	authcase     usecase.AuthUseCase
	tokenService authenticator.AccessToken
	engine       *gin.Engine
	host         string
}

// Function for server
func ServerAccount() *appServer {
	g := gin.Default()
	config := config.NewConfig()
	admin := manager.NewAdminManager(config)
	pr := manager.NewRepositoryManager(admin)
	pus := manager.NewUseCaseManager(pr)
	ts := authenticator.NewAccessToken(config.TokenConfig)
	ac := usecase.NewAuthUseCase(ts)
	if config.ApiHost == "" || config.ApiPort == "" {
		panic("No Host or Port")
	}
	host := fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort)
	return &appServer{
		acases:       pus.AccountUseCase(),
		pcase:        pus.ProductUseCase(),
		prcase:       pus.PriceUseCase(),
		porcase:      pus.PortofolioUseCase(),
		trcase:       pus.TransactionProductUseCase(),
		trxcase:      pus.TransactionDetailUseCase(),
		authcase:     ac,
		tokenService: ts,
		engine:       g,
		host:         host,
	}
}

// Function handler for server
func (as *appServer) initHandlerAccount() {
	controller.NewAccountController(as.engine, as.acases)
	controller.NewProdctController(as.engine, as.pcase)
	controller.NewPriceController(as.engine, as.prcase)
	controller.NewPortofolioController(as.engine, as.porcase)
	controller.NewTransactionProductController(as.engine, as.trcase)
	controller.NewTransactionDetailController(as.engine, as.trxcase)
	publicRoute := as.engine.Group("/login")
	tokenMdw := middleware.NewTokenValidator(as.tokenService)
	controller.NewAppController(publicRoute, as.authcase, tokenMdw)
}

func (as *appServer) Run() {
	as.initHandlerAccount()
	err := as.engine.Run(as.host)
	if err != nil {
		panic(err)
	}
}
