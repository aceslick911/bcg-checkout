package controllers

import (
	"errors"
	"log"
	"net/http"

	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/baskets"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
	http_err "github.com/aceslick911/bcg-checkout/pkg/http-err"
	"github.com/gin-gonic/gin"
)

// GetBasketById godoc
// @Summary Retrieves basket based on given ID
// @Description get Basket by ID
// @Produce json
// @Param id path integer true "Basket ID"
// @Success 200 {object} baskets.Basket
// @Router /api/baskets/{id} [get]
// @Security Authorization Token
func GetBasketById(c *gin.Context) {
	s := persistence.GetBasketRepository()
	id := c.Param("id")
	if basket, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("basket not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, basket)
	}
}

// GetBaskets godoc
// @Summary Retrieves baskets based on query
// @Description Get Baskets
// @Produce json
// @Param basketname query string false "Basketname"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []baskets.Basket
// @Router /api/baskets [get]
// @Security Authorization Token
func GetBaskets(c *gin.Context) {
	s := persistence.GetBasketRepository()
	var q models.Basket
	_ = c.Bind(&q)
	if baskets, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("baskets not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, baskets)
	}
}

func CreateBasket(c *gin.Context) {
	s := persistence.GetBasketRepository()
	var basketInput models.Basket
	_ = c.BindJSON(&basketInput)
	if err := s.Add(&basketInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, basketInput)
	}
}

func UpdateBasket(c *gin.Context) {
	s := persistence.GetBasketRepository()
	id := c.Params.ByName("id")
	var basketInput models.Basket
	_ = c.BindJSON(&basketInput)
	if _, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("basket not found"))
		log.Println(err)
	} else {
		if err := s.Update(&basketInput); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, basketInput)
		}
	}
}

func DeleteBasket(c *gin.Context) {
	s := persistence.GetBasketRepository()
	id := c.Params.ByName("id")
	var basketInput models.Basket
	_ = c.BindJSON(&basketInput)
	if basket, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("basket not found"))
		log.Println(err)
	} else {
		if err := s.Delete(basket); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
