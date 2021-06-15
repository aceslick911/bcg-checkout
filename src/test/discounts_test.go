package test

import (
	"fmt"
	"testing"

	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/discounts"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
)

var discountTest models.Discount

func TestAddDiscount(t *testing.T) {
	Setup()
	discountList := [...]models.Discount{
		{
			"condition_qty": 1,
			"condition_item": "43N23P",
			"discount_type": "FREE_ITEM_QTY",
			"discount_item": "234234",
			"discount_value": 1,
		},
		{
			"condition_qty": 3,
			"condition_item": "120P90",
			"discount_type": "FREE_ITEM_QTY",
			"discount_item": "120P90",
			"discount_value": 1,
		},
		{
			"condition_qty": 3,
			"condition_item": "A304SD",
			"discount_type": "ITEM_DISCOUNT_PERCENTAGE",
			"discount_item": "A304SD",
			"discount_value": 10,
		}
	}

	discount := discountList[0]
	s := persistence.GetDiscountRepository()
	if err := s.Add(&discount); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	discountTest = discount
}

func TestGetAllDiscounts(t *testing.T) {
	s := persistence.GetDiscountRepository()
	if _, err := s.All(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetDiscountById(t *testing.T) {
	db.SetupDB()
	db.SetupDB()
	s := persistence.GetDiscountRepository()
	if _, err := s.Get(fmt.Sprint(discountTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
