package products

import (
	"github.com/aceslick911/bcg-checkout/internal/pkg/models"
	// "github.com/aceslick911/bcg-checkout/internal/pkg/models/users"
	"time"
)

type Product struct {
	models.Model
	SKU           string  `gorm:"column:sku;not null;" json:"sku" form:"text"`
	Name          string  `gorm:"column:name;not null;" json:"name" form:"name"`
	Price         float64 `gorm:"column:price;not null;" json:"price" form:"price"`
	Inventory_Qty uint64  `gorm:"column:inventory_qty;not null;" json:"inventory_qty" form:"inventory_qty"`
	// UserID uint64     `gorm:"column:user_id;unique_index:user_id;not null;" json:"user_id" form:"user_id"`
	// User   users.User `json:"user"`
}

func (m *Product) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Product) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
