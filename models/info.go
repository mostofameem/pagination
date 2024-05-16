package models

type Product struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Category string `json:"category" db:"category"`
	Price    string `json:"price" db:"price"`
	Quantity string `json:"quantity" db:"quantity"`
}
