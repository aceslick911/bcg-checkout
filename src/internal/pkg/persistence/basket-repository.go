package persistence

import (
	"strconv"

	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/baskets"
)

type BasketRepository struct{}

var basketRepository *BasketRepository

func GetBasketRepository() *BasketRepository {
	if basketRepository == nil {
		basketRepository = &BasketRepository{}
	}
	return basketRepository
}

func (r *BasketRepository) Get(id string) (*models.Basket, error) {
	var basket models.Basket
	where := models.Basket{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &basket, []string{}) //, []string{"User"})
	if err != nil {
		return nil, err
	}
	return &basket, err
}

func (r *BasketRepository) All() (*[]models.Basket, error) {
	var baskets []models.Basket
	err := Find(&models.Basket{}, &baskets, []string{}) //, []string{"User"}, "id asc")
	return &baskets, err
}

func (r *BasketRepository) Query(q *models.Basket) (*[]models.Basket, error) {
	var baskets []models.Basket
	err := Find(&q, &baskets, []string{}) //, []string{"User"}, "id asc")
	return &baskets, err
}

func (r *BasketRepository) Add(basket *models.Basket) error {
	err := Create(&basket)
	err = Save(&basket)
	return err
}

func (r *BasketRepository) Update(basket *models.Basket) error {
	return db.GetDB().Save(&basket).Error
	//.Omit("User")

}

func (r *BasketRepository) Delete(basket *models.Basket) error {
	return db.GetDB().Unscoped().Delete(&basket).Error
}
