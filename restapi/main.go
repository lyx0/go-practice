package main

import (
	"net/http"

	"github.com/lyx0/go-practice/restapi/handlers"
)

func main() {
	ph := handlers.NewProducts()

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	http.ListenAndServe(":8080", sm)

}
