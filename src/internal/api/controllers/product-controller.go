package controllers

import (
	"errors"
	"log"
	"net/http"

	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/products"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
	http_err "github.com/aceslick911/bcg-checkout/pkg/http-err"
	"github.com/gin-gonic/gin"
)

// GetProductById godoc
// @Summary Retrieves product based on given ID
// @Description get Product by ID
// @Produce json
// @Param id path integer true "Product ID"
// @Success 200 {object} products.Product
// @Router /api/products/{id} [get]
// @Security Authorization Token
func GetProductById(c *gin.Context) {
	s := persistence.GetProductRepository()
	id := c.Param("id")
	if product, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("product not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// GetProducts godoc
// @Summary Retrieves products based on query
// @Description Get Products
// @Produce json
// @Param sku query string false "SKU"
// @Param name query string false "Name"
// @Param price query string false "Price"
// @Param inventory_qty query string false "Inventory_Qty"
// @Success 200 {array} []products.Product
// @Router /api/products [get]
// @Security Authorization Token
func GetProducts(c *gin.Context) {
	s := persistence.GetProductRepository()
	var q models.Product
	_ = c.Bind(&q)
	if products, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("products not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

func CreateProduct(c *gin.Context) {
	s := persistence.GetProductRepository()
	var productInput models.Product
	_ = c.BindJSON(&productInput)
	if err := s.Add(&productInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, productInput)
	}
}

func UpdateProduct(c *gin.Context) {
	s := persistence.GetProductRepository()
	id := c.Params.ByName("id")
	var productInput models.Product
	_ = c.BindJSON(&productInput)
	if _, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("product not found"))
		log.Println(err)
	} else {
		if err := s.Update(&productInput); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, productInput)
		}
	}
}

func DeleteProduct(c *gin.Context) {
	s := persistence.GetProductRepository()
	id := c.Params.ByName("id")
	var productInput models.Product
	_ = c.BindJSON(&productInput)
	if product, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("product not found"))
		log.Println(err)
	} else {
		if err := s.Delete(product); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
