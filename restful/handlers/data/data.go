package data

import "time"

//Product Struct
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// GetProducts List
func GetProducts() []*Product {
	return ProductList
}

// ProductList is a list of products
var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy Milk Coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and Strong Coffee without Milk",
		Price:       1.99,
		SKU:         "xyz789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}
