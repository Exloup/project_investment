package controller

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct Controller
type PortofolioController struct {
	router  *gin.Engine
	usecase usecase.PortofolioUseCase
}

// Function Usecase Insert / Create
func (por *PortofolioController) CreatePortofolio(ctx *gin.Context) {
	var porto model.Portofolios
	err := ctx.ShouldBindJSON(&porto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	p, err := por.usecase.CreatePortofolio(&porto)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, p)
}

// Function Usecase Get All
func (por *PortofolioController) GetAllPortofolio(ctx *gin.Context) {
	port, err := por.usecase.GetAllPortofolio()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, port)
}

// Function Usecase Get by ID
func (por *PortofolioController) GetPortofolioById(ctx *gin.Context) {
	id := ctx.Param("id")
	port, err := por.usecase.GetPortofolioById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, port)
}

// Function Usecase Update by ID
func (por *PortofolioController) UpdatePortofolio(ctx *gin.Context) {
	var porto model.Portofolios
	err := ctx.ShouldBindJSON(&porto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
}

// Function Usecase Delete by ID
func (por *PortofolioController) DeletePortofolio(ctx *gin.Context) {
	var porto model.Portofolios
	id := ctx.Param("id")
	err := por.usecase.DeletePortofolio(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, porto)
}

// Store Controller
func NewPortofolioController(router *gin.Engine, portofolioUseCase usecase.PortofolioUseCase) *PortofolioController {
	newPortofolioController := PortofolioController{
		router:  router,
		usecase: portofolioUseCase,
	}
	router.POST("/portofolio/insert", newPortofolioController.CreatePortofolio)
	router.GET("/portofolio/getall", newPortofolioController.GetAllPortofolio)
	router.GET("/portofolio/getall/:id", newPortofolioController.GetPortofolioById)
	router.PUT("/portofolio/update/:id", newPortofolioController.UpdatePortofolio)
	router.DELETE("/portofolio/delete/:id", newPortofolioController.DeletePortofolio)
	return &newPortofolioController
}
