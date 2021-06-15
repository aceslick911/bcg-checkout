package persistence

import (
	"strconv"

	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/discounts"
)

type DiscountRepository struct{}

var discountRepository *DiscountRepository

func GetDiscountRepository() *DiscountRepository {
	if discountRepository == nil {
		discountRepository = &DiscountRepository{}
	}
	return discountRepository
}

func (r *DiscountRepository) Get(id string) (*models.Discount, error) {
	var discount models.Discount
	where := models.Discount{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &discount, []string{}) //, []string{"User"})
	if err != nil {
		return nil, err
	}
	return &discount, err
}

func (r *DiscountRepository) All() (*[]models.Discount, error) {
	var discounts []models.Discount
	err := Find(&models.Discount{}, &discounts, []string{}) //, []string{"User"}, "id asc")
	return &discounts, err
}

func (r *DiscountRepository) Query(q *models.Discount) (*[]models.Discount, error) {
	var discounts []models.Discount
	err := Find(&q, &discounts, []string{}) //, []string{"User"}, "id asc")
	return &discounts, err
}

func (r *DiscountRepository) Add(discount *models.Discount) error {
	err := Create(&discount)
	err = Save(&discount)
	return err
}

func (r *DiscountRepository) Update(discount *models.Discount) error {
	return db.GetDB().Save(&discount).Error
	//.Omit("User")

}

func (r *DiscountRepository) Delete(discount *models.Discount) error {
	return db.GetDB().Unscoped().Delete(&discount).Error
}
