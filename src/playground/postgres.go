package main

import (
	"fmt"
	"playground/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	dsn := "host=localhost user=fabioteichmann password= dbname=postgres port=5433"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	db, err := initDB()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DB = ", &db)

	prod := models.Product{}
	result := db.First(&prod, 2)
	fmt.Println(prod, result)

	var tables []string
	if err := db.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error; err != nil {
		panic(err)
	}
	fmt.Println(tables)
}
