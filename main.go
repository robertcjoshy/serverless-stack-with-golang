package main

import (
	"fmt"

	"github.com/robert/serverless-stack-with-golang/src/models"
	"github.com/robert/serverless-stack-with-golang/src/routes"
)

func main() {
	//utils.LoadEnv()
	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	router := routes.SetupRoutes()
	router.Run("localhost:8080")
}

func robert() {
	fmt.Println("connected")
}
