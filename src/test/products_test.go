package test

import (
	"fmt"
	"testing"

	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/products"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
)

var productTest models.Product

var sampleData = persistence.SampleDatabase()

func TestAddProduct(t *testing.T) {
	Setup()
	// fmt.Printf("%+v\n", sampleData)
	product := sampleData.Products[0]
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
