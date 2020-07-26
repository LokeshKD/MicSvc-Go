package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"./data"
)

// Products Struct
type Products struct {
	l *log.Logger
}

// NewProducts Func
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to Marshall json", http.StatusInternalServerError)
		return
	}
	rw.Write(d)
}
