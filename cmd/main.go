package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mkdigo/go-product-api.git/controller"
	"github.com/mkdigo/go-product-api.git/db"
	"github.com/mkdigo/go-product-api.git/repository"
	"github.com/mkdigo/go-product-api.git/usecase"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	//Camada Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Camada useCase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/products", ProductController.CreateProduct)
	server.GET("/products/:productId", ProductController.GetProductById)

	server.Run(":80")
}
