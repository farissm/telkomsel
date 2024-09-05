package handler

import (
	"fmt"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"telkomsel/product"
	"telkomsel/helper"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (p *productHandler) CreateProduct(c *gin.Context) {
	var product product.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		logrus.Errorf("FAILED bind json input product data %v", err)

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	

	createProduct, err := p.productService.CreateProduct(product)
	if err != nil {
		logrus.Errorf("FAILED create product at handler %v", err)

		response := helper.APIResponse(http.StatusInternalServerError, "failed create product. Please call customer service")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, createProduct)
}

func (p *productHandler) GetProduct(c *gin.Context) {
	page := c.Query("page")
	pageInt := 1
	if page != "" {
		pageIntParse, err := strconv.Atoi(page)
		if err != nil {
			logrus.Errorf("Error when converting product page value. Error: %s", err.Error())
	
			response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		pageInt = pageIntParse
	}
	

	products, err := p.productService.GetProducts(pageInt)
	if err != nil {
		logrus.Errorf("FAILED get product at handler %v", err)

		response := helper.APIResponse(http.StatusInternalServerError, "failed fetch data")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (p *productHandler) GetProductById(c *gin.Context) {
	productId := c.Query("id")

	if productId == "" {
		response := helper.APIResponse(http.StatusBadRequest, "missing product id")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		logrus.Errorf("Error when converting id. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	

	product, err := p.productService.GetProductById(productIdInt)
	if err != nil {
		logrus.Errorf("FAILED get product by id at handler %v", err)

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *productHandler) UpdateProduct(c *gin.Context) {
	productId := c.Query("id")
	
	if productId == "" {
		response := helper.APIResponse(http.StatusBadRequest, "missing product id")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		logrus.Errorf("Error when converting id. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var product product.Product

	err = c.ShouldBindJSON(&product)
	if err != nil {
		logrus.Errorf("FAILED bind json input product update data %v", err)

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	

	if err := p.productService.UpdateProduct(productIdInt, product); err != nil {
		logrus.Errorf("FAILED update product at handler %v", err)

		response := helper.APIResponse(http.StatusBadRequest, "failed update product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	response := helper.APIResponse(http.StatusOK, fmt.Sprintf("Product with id %s updated", productId))
	c.JSON(http.StatusOK, response)
}

func (p *productHandler) DeleteProduct(c *gin.Context) {
	productId := c.Query("id")
	
	if productId == "" {
		response := helper.APIResponse(http.StatusBadRequest, "missing product id")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		logrus.Errorf("Error when converting id. Error: %s", err.Error())

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	
	if err := p.productService.DeleteProduct(productIdInt); err != nil {
		logrus.Errorf("FAILED delete product at handler %v", err)

		response := helper.APIResponse(http.StatusInternalServerError, "internal server error")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, fmt.Sprintf("Product with id %s deleted", productId))
	c.JSON(http.StatusOK, response)
}
