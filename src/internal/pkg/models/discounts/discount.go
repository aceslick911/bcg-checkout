package discounts

import (
	"github.com/aceslick911/bcg-checkout/internal/pkg/models"
	// "github.com/aceslick911/bcg-checkout/internal/pkg/models/users"
	"time"
)

type Discount struct {
	models.Model
	Condition_qty  uint64  `gorm:"column:condition_qty;not null;" json:"condition_qty" form:"condition_qty"`
	Condition_item string  `gorm:"column:condition_item;not null;" json:"condition_items" form:"condition_items"`
	Discount_type  string  `gorm:"column:discount_type;not null;" json:"discount_type" form:"discount_type"`
	Discount_item  string  `gorm:"column:discount_item;" json:"discount_item" form:"discount_item"`
	Discount_value float64 `gorm:"column:discount_value;" json:"discount_value" form:"discount_value"`
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
