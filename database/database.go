package database

import (
	"fmt"
	"image-service/helpers"
	"image-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dbName string) error {
	var (
		databaseUser     = helpers.GetValueFromEnv("DB_USER")
		databasePassword = helpers.GetValueFromEnv("DB_PASSWORD")
		databaseHost     = helpers.GetValueFromEnv("DB_HOST")
		databasePort     = helpers.GetValueFromEnv("DB_PORT")
		databaseName     = dbName
	)
	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseHost, databasePort, databaseName)
	database, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		return err
	}
	database.AutoMigrate(&models.Image{}, &models.Item{}, &models.Shop{})
	DB = database
	fmt.Printf("Conntected to database\n")
	return nil
}


