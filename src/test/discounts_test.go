package test

import (
	"fmt"
	"testing"

	"github.com/aceslick911/bcg-checkout/internal/pkg/config"
	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/discounts"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
)

var discountTest models.Discount

func SetupDiscounts() {
	config.Setup("./config.yml")
	db.SetupDB()
	db.GetDB().Exec("DELETE FROM discounts")
}

func TestAddDiscount(t *testing.T) {
	SetupDiscounts()

	fmt.Printf("%+v\n", sampleData)
	discount := sampleData.Discounts[0]
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
