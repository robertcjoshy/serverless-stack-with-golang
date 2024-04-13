package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email   string `json:"email" bson:"email"`
	Name    string `json:"name" bson:"name"`
	Country string `json:"country" bson:"country"`
}
