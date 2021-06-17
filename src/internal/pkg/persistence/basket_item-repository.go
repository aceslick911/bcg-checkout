package persistence

import (
	"strconv"

	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/basket_items"
)

type Basket_ItemRepository struct{}

var basket_itemRepository *Basket_ItemRepository

func GetBasket_ItemRepository() *Basket_ItemRepository {
	if basket_itemRepository == nil {
		basket_itemRepository = &Basket_ItemRepository{}
	}
	return basket_itemRepository
}

func (r *Basket_ItemRepository) Get(id string) (*models.Basket_Item, error) {
	var basket_item models.Basket_Item
	where := models.Basket_Item{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &basket_item, []string{}) //, []string{"User"})
	if err != nil {
		return nil, err
	}
	return &basket_item, err
}

func (r *Basket_ItemRepository) All() (*[]models.Basket_Item, error) {
	var basket_items []models.Basket_Item
	err := Find(&models.Basket_Item{}, &basket_items, []string{}) //, []string{"User"}, "id asc")
	return &basket_items, err
}

func (r *Basket_ItemRepository) Query(q *models.Basket_Item) (*[]models.Basket_Item, error) {
	var basket_items []models.Basket_Item
	err := Find(&q, &basket_items, []string{}) //, []string{"User"}, "id asc")
	return &basket_items, err
}

func (r *Basket_ItemRepository) Add(basket_item *models.Basket_Item) error {
	err := Create(&basket_item)
	err = Save(&basket_item)
	return err
}

func (r *Basket_ItemRepository) Update(basket_item *models.Basket_Item) error {
	return db.GetDB().Save(&basket_item).Error
	//.Omit("User")

}

func (r *Basket_ItemRepository) Delete(basket_item *models.Basket_Item) error {
	return db.GetDB().Unscoped().Delete(&basket_item).Error
}
