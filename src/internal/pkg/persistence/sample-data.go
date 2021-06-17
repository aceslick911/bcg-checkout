package persistence

import (
	"fmt"
	"time"

	"github.com/aceslick911/bcg-checkout/internal/pkg/models"
	basket_items "github.com/aceslick911/bcg-checkout/internal/pkg/models/basket_items"
	baskets "github.com/aceslick911/bcg-checkout/internal/pkg/models/baskets"
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
			Name:  "Alexa SpeakerFFFFFFFF",
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
	Baskets = [...]baskets.Basket{

		{Model: newModel(),
			User:      0,
			Total:     100,
			Discounts: 10,
		},
	}
	Basket_Items = [...]basket_items.Basket_Item{

		{Model: newModel(),
			Basket_ID:   0,
			Discount_ID: 0,

			Item_Type:   "FREE_ITEM_QTY",
			Product_SKU: "120P90",

			Cost:           100,
			After_Discount: 0,
		},
	}
)

type SampleDatabaseType struct {
	Products     []products.Product
	Discounts    []discounts.Discount
	Baskets      []baskets.Basket
	Basket_Items []basket_items.Basket_Item
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

	var baskets = func() []baskets.Basket {
		var Output = make([]baskets.Basket, len(Baskets))
		copier.Copy(Baskets, &Output)
		return Output
	}

	var basket_items = func() []basket_items.Basket_Item {
		var Output = make([]basket_items.Basket_Item, len(Basket_Items))
		copier.Copy(Basket_Items, &Output)
		return Output
	}

	return SampleDatabaseType{
		Products:     products(),
		Discounts:    discounts(),
		Baskets:      baskets(),
		Basket_Items: basket_items(),
	}
}

var sampleData = SampleDatabase()

func HydrateDatabase() {
	products := sampleData.Products
	s := GetProductRepository()

	for _, prod := range products {
		if err := s.Add(&prod); err != nil {
			fmt.Println("ERR", err)
		}
	}

	//productTest = product
}
