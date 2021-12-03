package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int64   `json:"quantity"`
}

type Products []*Product

func GetProducts() Products {
	return productList
}

func (p *Products) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Product 1",
		Description: "Description 1",
		Price:       132.4,
		Quantity:    53,
	},
	{
		ID:          2,
		Name:        "Product 2",
		Description: "Description 2",
		Price:       45.4,
		Quantity:    2,
	},
}
