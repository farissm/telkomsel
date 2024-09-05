package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"

	"telkomsel/config"
	"telkomsel/product"
	"telkomsel/handler"
)

func main() {
	fmt.Println("Starting Project ...")

	rl, logErr := rotatelogs.New("./storage/logs/telkomsel-%Y-%m-%d.log")
	if logErr != nil {
		panic(logErr)
	}

	log.SetOutput(rl)
	log.Info("logging to file success")

	db := config.ConnectDB()
	fmt.Println(db)
	fmt.Println("Connected to database")

	productRepository := product.NewRepository(db)

	productService := product.NewService(productRepository)

	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Telkomsel Test RESTful API")
	})

	router.POST("/create-product", productHandler.CreateProduct)
	router.GET("/products", productHandler.GetProduct)
	router.GET("/product", productHandler.GetProductById)
	router.PUT("/product", productHandler.UpdateProduct)
	router.DELETE("/product", productHandler.DeleteProduct)

	// router.Run()
	router.Run(":"+os.Getenv("APP_PORT"))
}

