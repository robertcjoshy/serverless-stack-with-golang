package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func OpenDatabaseConnection() {
	var err error
	host := "database-robert.chsu6q8eqfk8.ap-south-1.rds.amazonaws.com"
	username := "postgres"
	password := "12345678"
	databaseName := ""
	port := "5432"

	dsn := fmt.Sprintf("host=%s,user=%s password=%s dbname=%s,port=%s", host, username, password, databaseName, port)

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("🚀🚀🚀---ASCENDE SUPERIUS---🚀🚀🚀")
	}
}

func AutoMigrateModels() {
	Database.AutoMigrate(&Startup{})
}
