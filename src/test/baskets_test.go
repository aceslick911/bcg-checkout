package test

import (
	"fmt"
	"testing"

	"github.com/aceslick911/bcg-checkout/internal/pkg/config"
	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	models "github.com/aceslick911/bcg-checkout/internal/pkg/models/baskets"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
)

var basketTest models.Basket

func SetupBaskets() {
	config.Setup("./config.yml")
	db.SetupDB()
	db.GetDB().Exec("DELETE FROM baskets")
}

func TestAddBasket(t *testing.T) {
	SetupBaskets()

	fmt.Printf("%+v\n", sampleData)
	basket := sampleData.Baskets[0]
	s := persistence.GetBasketRepository()
	if err := s.Add(&basket); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	basketTest = basket
}

func TestGetAllBaskets(t *testing.T) {
	s := persistence.GetBasketRepository()
	if _, err := s.All(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetBasketById(t *testing.T) {
	db.SetupDB()
	s := persistence.GetBasketRepository()
	if _, err := s.Get(fmt.Sprint(basketTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
