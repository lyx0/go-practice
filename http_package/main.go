package main

import (
	"net/http"
)

type Index struct {
	text string
}

func (i *Index) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Index"))
}

func newIndex(text string) *Index {
	return &Index{text}
}

func main() {
	i := newIndex("test")

	sm := http.NewServeMux()
	sm.Handle("/", i)

	http.ListenAndServe(":8080", sm)
}
