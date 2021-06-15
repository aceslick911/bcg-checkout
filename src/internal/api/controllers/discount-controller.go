package controllers

import (
	"errors"
	"log"
	"net/http"

	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/discounts"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
	http_err "github.com/aceslick911/bcg-checkout/pkg/http-err"
	"github.com/gin-gonic/gin"
)

// GetDiscountById godoc
// @Summary Retrieves discount based on given ID
// @Description get Discount by ID
// @Produce json
// @Param id path integer true "Discount ID"
// @Success 200 {object} discounts.Discount
// @Router /api/discounts/{id} [get]
// @Security Authorization Token
func GetDiscountById(c *gin.Context) {
	s := persistence.GetDiscountRepository()
	id := c.Param("id")
	if discount, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("discount not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, discount)
	}
}

// GetDiscounts godoc
// @Summary Retrieves discounts based on query
// @Description Get Discounts
// @Produce json
// @Param discountname query string false "Discountname"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []discounts.Discount
// @Router /api/discounts [get]
// @Security Authorization Token
func GetDiscounts(c *gin.Context) {
	s := persistence.GetDiscountRepository()
	var q models.Discount
	_ = c.Bind(&q)
	if discounts, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("discounts not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, discounts)
	}
}

func CreateDiscount(c *gin.Context) {
	s := persistence.GetDiscountRepository()
	var discountInput models.Discount
	_ = c.BindJSON(&discountInput)
	if err := s.Add(&discountInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, discountInput)
	}
}

func UpdateDiscount(c *gin.Context) {
	s := persistence.GetDiscountRepository()
	id := c.Params.ByName("id")
	var discountInput models.Discount
	_ = c.BindJSON(&discountInput)
	if _, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("discount not found"))
		log.Println(err)
	} else {
		if err := s.Update(&discountInput); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, discountInput)
		}
	}
}

func DeleteDiscount(c *gin.Context) {
	s := persistence.GetDiscountRepository()
	id := c.Params.ByName("id")
	var discountInput models.Discount
	_ = c.BindJSON(&discountInput)
	if discount, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("discount not found"))
		log.Println(err)
	} else {
		if err := s.Delete(discount); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
