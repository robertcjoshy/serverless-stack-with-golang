package models

import "gorm.io/gorm"

type Startup struct {
	gorm.Model
	Name             string  `json:"name" bson:"name"`
	CEO              string  `json:"ceo" bson:"ceo"`
	ValueProposition string  `json:"value_proposition" bson:"value_proposition"`
	Industry         string  `json:"industry" bson:"industry"`
	Founder          Founder `json:"founder" bson:"founder" gorm:"embedded"`
}

type Founder struct {
	gorm.Model
	Name       string `json:"name" bson:"name"`
	Profession string `json:"profession" bson:"profession"`
	Age        int    `json:"age" bson:"age"`
}
