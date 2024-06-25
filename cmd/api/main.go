package main

import (
	"image-service/database"
	"image-service/helpers"
	"image-service/routes"
)

func main() {
	routes.InitRoute()
	var dbName = helpers.GetValueFromEnv("DB_NAME")
	if err := database.InitDatabase(dbName); err != nil {
		panic(err.Error())
	}
}
