package models

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "database-robert.chsu6q8eqfk8.ap-south-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "daisyjoshy"
	dbname   = "robertfirst"
)

var Database *gorm.DB

func init() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s", host, port, user, password)

	Database, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func SaveEmail(mail string) (*Startup, error) {
	var startup Startup
	err := Database.Model(&startup).Create(&startup).Error
	if err != nil {
		return nil, err
	}
	return &startup, err
}

func GetEmail(search_mail string) (*Startup, error) {

}
