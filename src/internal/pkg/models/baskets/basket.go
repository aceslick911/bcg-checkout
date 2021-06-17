package baskets

import (
	"github.com/aceslick911/bcg-checkout/internal/pkg/models"
	// "github.com/aceslick911/bcg-checkout/internal/pkg/models/users"
	"time"
)

type Basket struct {
	models.Model
	User      int     `gorm:"column:user;" json:"user" form:"user"`
	Total     float64 `gorm:"column:total;" json:"total" form:"total"`
	Discounts float64 `gorm:"column:discounts;" json:"discounts" form:"discounts"`
}

func (m *Basket) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Basket) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
