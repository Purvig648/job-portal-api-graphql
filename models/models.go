package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required"`
	HashPassword string `json:"hash_password" validate:"required"`
}
type Company struct {
	gorm.Model
	Name     string `json:"companyname" validate:"required"`
	Location string `json:"companylocation" validate:"required"`
}
type Job struct {
	gorm.Model
	Cid    string `json:"cid"`
	Role   string `json:"role"`
	Salary string `json:"salary"`
}
