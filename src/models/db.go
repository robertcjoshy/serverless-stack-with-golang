package models

import "gorm.io/gorm"

type Startup struct {
	gorm.Model
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
