package test

import (
	"fmt"
	"testing"

	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/products"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
)

var productTest models.Product

func TestAddProduct(t *testing.T) {
	Setup()
	productList := [...]models.Product{
		{SKU: "120P90", Name: "Google Home", Price: 49.99, Inventory_Qty: 10},
		{SKU: "43N23P", Name: "MacBook Pro", Price: 5399.99, Inventory_Qty: 5},
		{SKU: "A304SD", Name: "Alexa Speaker", Price: 109.5, Inventory_Qty: 10},
		{SKU: "234234", Name: "Raspberry Pi B", Price: 30, Inventory_Qty: 2},
	}

	product := productList[0]
	s := persistence.GetProductRepository()
	if err := s.Add(&product); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	productTest = product
}

func TestGetAllProducts(t *testing.T) {
	s := persistence.GetProductRepository()
	if _, err := s.All(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetProductById(t *testing.T) {
	db.SetupDB()
	db.SetupDB()
	s := persistence.GetProductRepository()
	if _, err := s.Get(fmt.Sprint(productTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
