package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello Struct
type Hello struct {
	l *log.Logger
}

// NewHello Function
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//log.Println("Hello World")
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s\n", d)
}
