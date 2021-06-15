package discounts

import (
	"github.com/aceslick911/bcg-checkout/internal/pkg/models"
	// "github.com/aceslick911/bcg-checkout/internal/pkg/models/users"
	"time"
)

type Discount struct {
	models.Model
	condition_qty  uint64  `gorm:"column:inventory_qty;not null;" json:"inventory_qty" form:"inventory_qty"`
	condition_item string  `gorm:"column:sku;not null;" json:"sku" form:"text"`
	discount_type  string  `gorm:"column:sku;not null;" json:"sku" form:"text"`
	discount_item  string  `gorm:"column:sku;not null;" json:"sku" form:"text"`
	discount_value float64 `gorm:"column:price;not null;" json:"price" form:"price"`
	// UserID uint64     `gorm:"column:user_id;unique_index:user_id;not null;" json:"user_id" form:"user_id"`
	// User   users.User `json:"user"`
}

func (m *Discount) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Discount) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
