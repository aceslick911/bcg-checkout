package persistence

import (
	"time"

	"github.com/aceslick911/bcg-checkout/internal/pkg/models"
	discounts "github.com/aceslick911/bcg-checkout/internal/pkg/models/discounts"
	products "github.com/aceslick911/bcg-checkout/internal/pkg/models/products"
	users "github.com/aceslick911/bcg-checkout/internal/pkg/models/users"
	"github.com/jinzhu/copier"
)

var creationDate = time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC)

func newModel() models.Model {
	return models.Model{
		ID:        0,
		CreatedAt: creationDate,
		UpdatedAt: time.Now(),
	}
}

var (
	Products = [...]products.Product{
		{
			Model: newModel(),
			SKU:   "120P90",
			Name:  "Google Home",
			Price: 49.99, Inventory_Qty: 10},
		{
			Model: newModel(),
			SKU:   "43N23P",
			Name:  "MacBook Pro",
			Price: 5399.99, Inventory_Qty: 5},
		{
			Model: newModel(),
			SKU:   "A304SD",
			Name:  "Alexa Speaker",
			Price: 109.5, Inventory_Qty: 10},
		{
			Model: newModel(),
			SKU:   "234234",
			Name:  "Raspberry Pi B",
			Price: 30, Inventory_Qty: 2},
	}
	Users = [...]users.User{
		{
			Model:     newModel(),
			Username:  "aceslick911",
			Firstname: "Angelo",
			Lastname:  "Perera",
			Hash:      "ABC123",
		},
	}
	Discounts = [...]discounts.Discount{

		{Model: newModel(),
			Condition_qty:  1,
			Condition_item: "43N23P",
			Discount_type:  "FREE_ITEM_QTY",
			Discount_item:  "234234",
			Discount_value: 1,
		},
		{Model: newModel(),
			Condition_qty:  3,
			Condition_item: "120P90",
			Discount_type:  "FREE_ITEM_QTY",
			Discount_item:  "120P90",
			Discount_value: 1,
		},
		{Model: newModel(),
			Condition_qty:  3,
			Condition_item: "A304SD",
			Discount_type:  "ITEM_DISCOUNT_PERCENTAGE",
			Discount_value: 10,
		},
	}
)

type SampleDatabaseType struct {
	Products  []products.Product
	Discounts []discounts.Discount
}

func SampleDatabase() SampleDatabaseType {

	var products = func() []products.Product {
		var Output = make([]products.Product, len(Products))
		copier.Copy(Products, &Output)
		return Output
	}

	var discounts = func() []discounts.Discount {
		var Output = make([]discounts.Discount, len(Discounts))
		copier.Copy(Discounts, &Output)
		return Output
	}

	return SampleDatabaseType{
		Products:  products(),
		Discounts: discounts(),
	}
}
