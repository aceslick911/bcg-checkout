package persistence

import (
	"strconv"

	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/products"
)

type ProductRepository struct{}

var productRepository *ProductRepository

func GetProductRepository() *ProductRepository {
	if productRepository == nil {
		productRepository = &ProductRepository{}
	}
	return productRepository
}

func (r *ProductRepository) Get(id string) (*models.Product, error) {
	var product models.Product
	where := models.Product{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &product, []string{}) //, []string{"User"})
	if err != nil {
		return nil, err
	}
	return &product, err
}

func (r *ProductRepository) All() (*[]models.Product, error) {
	var products []models.Product
	err := Find(&models.Product{}, &products, []string{}) //, []string{"User"}, "id asc")
	return &products, err
}

func (r *ProductRepository) Query(q *models.Product) (*[]models.Product, error) {
	var products []models.Product
	err := Find(&q, &products, []string{}) //, []string{"User"}, "id asc")
	return &products, err
}

func (r *ProductRepository) Add(product *models.Product) error {
	err := Create(&product)
	err = Save(&product)
	return err
}

func (r *ProductRepository) Update(product *models.Product) error {
	return db.GetDB().Save(&product).Error
	//.Omit("User")

}

func (r *ProductRepository) Delete(product *models.Product) error {
	return db.GetDB().Unscoped().Delete(&product).Error
}
