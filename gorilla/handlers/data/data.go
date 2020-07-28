package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

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
func GetProducts() Products {
	return productList
}

//AddProduct - adds a product to the list
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

//UpdateProduct - updates an existing product
func UpdateProduct(id int, p *Product) error {
	fp, pos, err := findProductByID(id)
	if err != nil {
		return err
	}
	fp.ID = id
	productList[pos] = fp

	return nil
}

// ErrProductNotFound error
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProductByID(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// FromJSON to get product
func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

// Products list
type Products []*Product

// ToJSON conversion of list of products
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// ProductList is a list of products
var productList = []*Product{
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
