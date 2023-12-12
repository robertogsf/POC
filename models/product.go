package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string  `jason:"name"`
	Quantity int     `jason:"quantity"`
	Price    float32 `jason:"price"`
}
