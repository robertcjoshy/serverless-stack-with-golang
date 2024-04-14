package models

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "database-robert.chsu6q8eqfk8.ap-south-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "daisyjoshy"
	dbname   = "robert_first"
)

var Database *gorm.DB

func init() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s "+"password=%s", host, port, user, dbname, password)

	Database, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Database.AutoMigrate(&User{})
}

func (user *User) SaveEmail() (*User, error) {
	//var startup S

	err := Database.Model(&user).Create(&user).Error
	//err := Database.Create(&startup)
	if err != nil {
		log.Println("error in creating to databse")
		return nil, err
	}
	return user, err
}

func GetEmail(search_mail string) (*User, error) {
	var startup User
	err := Database.Where("email = ?", search_mail).First(&startup).Error
	if err != nil {
		return nil, err
	}
	return &startup, nil
}
