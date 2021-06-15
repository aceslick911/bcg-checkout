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
			Condition_qty:  1,
			Condition_item: "43N23P",
			Discount_type:  "FREE_ITEM_QTY",
			Discount_item:  "234234",
			Discount_value: 1.0,
		},
		{
			Condition_qty:  3,
			Condition_item: "120P90",
			Discount_type:  "FREE_ITEM_QTY",
			Discount_item:  "120P90",
			Discount_value: 1.0,
		},
		{
			Condition_qty:  3,
			Condition_item: "A304SD",
			Discount_type:  "ITEM_DISCOUNT_PERCENTAGE",
			Discount_item:  "A304SD",
			Discount_value: 10.0,
		},
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
