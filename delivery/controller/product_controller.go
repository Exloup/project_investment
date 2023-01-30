package controller

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router  *gin.Engine
	usecase usecase.ProductUseCase
}

// Function Usecase Insert
func (pc *ProductController) InsertProduct(ctx *gin.Context) {
	var products model.Products
	err := ctx.ShouldBindJSON(&products)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	p, err := pc.usecase.InsertProduct(&products)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, p)
}

// Function Usecase Update by ID
func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	var products model.Products
	err := ctx.ShouldBindJSON(&products)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	p, err := pc.usecase.UpdateProduct(&products)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, p)
}

// Function Usecase Get All
func (pc *ProductController) GetAllProduct(ctx *gin.Context) {
	products, err := pc.usecase.GetAllProduct()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// Function Usecase Get by ID
func (pc *ProductController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	products, err := pc.usecase.GetProductById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// Function Usecase Delete by ID
func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	var products model.Products
	id := ctx.Param("id")
	err := pc.usecase.DeleteProduct(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// Store Product Controller
func NewProdctController(router *gin.Engine, productUseCase usecase.ProductUseCase) *ProductController {
	newProductController := ProductController{
		router:  router,
		usecase: productUseCase,
	}
	router.POST("/product/insert", newProductController.InsertProduct)
	router.GET("/product/getall", newProductController.GetAllProduct)
	router.GET("/product/getall/:id", newProductController.GetProductById)
	router.PUT("/product/update/:id", newProductController.UpdateProduct)
	router.DELETE("/product/delete/:id", newProductController.DeleteProduct)
	return &newProductController
}
