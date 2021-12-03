package handlers

import "net/http"

type Products struct{}

func NewProducts() *Products {
	return &Products{}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Product"))
}
