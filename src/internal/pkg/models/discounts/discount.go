package discounts

import (
	"github.com/aceslick911/bcg-checkout/internal/pkg/models"
	// "github.com/aceslick911/bcg-checkout/internal/pkg/models/users"
	"time"
)

type Discount struct {
	models.Model
	Condition_qty  uint64  `gorm:"column:condition_qty;not null;" json:"condition_qty" form:"condition_qty"`
	Condition_item string  `gorm:"column:Condition_item;not null;" json:"Condition_items" form:"Condition_items"`
	Discount_type  string  `gorm:"column:Discount_type;not null;" json:"Discount_type" form:"Discount_type"`
	Discount_item  string  `gorm:"column:Discount_item;not null;" json:"Discount_item" form:"Discount_item"`
	Discount_value float64 `gorm:"column:Discount_value;not null;" json:"Discount_value" form:"Discount_value"`
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
