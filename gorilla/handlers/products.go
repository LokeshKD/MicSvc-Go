// Package handlers classification of Product API
//
// Documentation for Product API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../data"
	"github.com/gorilla/mux"
)

//A list of products returned in response
//swagger:response productResponse
type productResponse struct {
	//All Products in the system
	//in: body
	Body []data.Product
}

// Products Struct
type Products struct {
	l *log.Logger
}

// NewProducts Func
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//AddProduct func
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

//UpdateProduct func
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, "Unable to convert Id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product ", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

// swagger:route GET / listProducts
// Returns a list of products
// responses:
// 	200: productResponse

//GetProducts func
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshall json", http.StatusInternalServerError)
		return
	}
}

//KeyProduct Struct
type KeyProduct struct{}

//MiddlewareProductValidator func
func (p *Products) MiddlewareProductValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)
			return
		}

		//Validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] Validation Product", err)
			http.Error(rw,
				fmt.Sprintf("Validation failed %s", err),
				http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
