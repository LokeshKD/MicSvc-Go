package handlers

import (
	"log"
	"net/http"
)

// Goodbye Struct
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye Func
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye")
	rw.Write([]byte("Byeee!\n"))
}
