package controllers

import (
	"errors"
	"log"
	"net/http"

	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/basket_items"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
	http_err "github.com/aceslick911/bcg-checkout/pkg/http-err"
	"github.com/gin-gonic/gin"
)

// GetBasket_ItemById godoc
// @Summary Retrieves basket_item based on given ID
// @Description get Basket_Item by ID
// @Produce json
// @Param id path integer true "Basket_Item ID"
// @Success 200 {object} basket_items.Basket_Item
// @Router /api/basket_items/{id} [get]
// @Security Authorization Token
func GetBasket_ItemById(c *gin.Context) {
	s := persistence.GetBasket_ItemRepository()
	id := c.Param("id")
	if basket_item, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("basket_item not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, basket_item)
	}
}

// GetBasket_Items godoc
// @Summary Retrieves basket_items based on query
// @Description Get Basket_Items
// @Produce json
// @Param basket_id query string false "Basket_ID"
// @Param discount_id query string false "Discount_ID"
// @Param item_type query string false "Item_Type"
// @Param product_sku query string false "Product_SKU"
// @Param cost query string false "Cost"
// @Param after_discount query string false "After_Discount"

// @Success 200 {array} []basket_items.Basket_Item
// @Router /api/basket_items [get]
// @Security Authorization Token
func GetBasket_Items(c *gin.Context) {
	s := persistence.GetBasket_ItemRepository()
	var q models.Basket_Item
	_ = c.Bind(&q)
	if basket_items, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("basket_items not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, basket_items)
	}
}

func CreateBasket_Item(c *gin.Context) {
	s := persistence.GetBasket_ItemRepository()
	var basket_itemInput models.Basket_Item
	_ = c.BindJSON(&basket_itemInput)
	if err := s.Add(&basket_itemInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, basket_itemInput)
	}
}

func UpdateBasket_Item(c *gin.Context) {
	s := persistence.GetBasket_ItemRepository()
	id := c.Params.ByName("id")
	var basket_itemInput models.Basket_Item
	_ = c.BindJSON(&basket_itemInput)
	if _, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("basket_item not found"))
		log.Println(err)
	} else {
		if err := s.Update(&basket_itemInput); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, basket_itemInput)
		}
	}
}

func DeleteBasket_Item(c *gin.Context) {
	s := persistence.GetBasket_ItemRepository()
	id := c.Params.ByName("id")
	var basket_itemInput models.Basket_Item
	_ = c.BindJSON(&basket_itemInput)
	if basket_item, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("basket_item not found"))
		log.Println(err)
	} else {
		if err := s.Delete(basket_item); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
