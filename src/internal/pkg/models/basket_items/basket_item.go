package basket_items

import (
	"github.com/aceslick911/bcg-checkout/internal/pkg/models"
	// "github.com/aceslick911/bcg-checkout/internal/pkg/models/users"
	"time"
)

type Basket_Item struct {
	models.Model

	Basket_ID   uint64 `gorm:"column:basket_id;not null;" json:"basket_id" form:"basket_id"`
	Discount_ID uint64 `gorm:"column:discount_id;not null;" json:"discount_id" form:"discount_id"`

	Item_Type   string `gorm:"column:item_type;not null;" json:"item_type" form:"item_type"`
	Product_SKU string `gorm:"column:product_sku;not null;" json:"product_sku" form:"product_sku"`

	Cost           float64 `gorm:"column:cost;not null;" json:"cost" form:"cost"`
	After_Discount float64 `gorm:"column:after_discount;not null;" json:"after_discount" form:"after_discount"`
}

func (m *Basket_Item) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Basket_Item) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
