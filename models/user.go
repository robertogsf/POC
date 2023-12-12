package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `jason:"email"`
	Password string `jason:"password"`
	Rol      string `jason:"rol"`
}
