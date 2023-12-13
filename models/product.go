package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `jason:"name"`
	Quantity string `jason:"quantity"`
	Price    string `jason:"price"`
}
