package test

import (
	"fmt"
	"testing"

	"github.com/aceslick911/bcg-checkout/internal/pkg/config"
	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/basket_items"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
)

var basket_itemTest models.Basket_Item

func SetupBasket_Items() {
	config.Setup("./config.yml")
	db.SetupDB()
	db.GetDB().Exec("DELETE FROM basket_items")
}

func TestAddBasket_Item(t *testing.T) {
	SetupBasket_Items()

	fmt.Printf("%+v\n", sampleData)
	basket_item := sampleData.Basket_Items[0]
	s := persistence.GetBasket_ItemRepository()
	if err := s.Add(&basket_item); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	basket_itemTest = basket_item
}

func TestGetAllBasket_Items(t *testing.T) {
	s := persistence.GetBasket_ItemRepository()
	if _, err := s.All(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetBasket_ItemById(t *testing.T) {
	db.SetupDB()
	s := persistence.GetBasket_ItemRepository()
	if _, err := s.Get(fmt.Sprint(basket_itemTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
