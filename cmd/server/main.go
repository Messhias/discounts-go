package main

import (
	"dgoo/internal/handlers"
	"dgoo/internal/models"
	"dgoo/internal/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Product{})

	repo := repository.NewProductRepository(db)

	productHandler := handlers.NewProductHandler(repo)

	router := gin.Default()
	router.GET("/products", productHandler.GetProducts)

	var count int64
	db.Model(&models.Product{}).Count(&count)
	if count == 0 {
		seedProducts(db)
	}

	router.Run(":8080")
}

func seedProducts(db *gorm.DB) {
	products := []models.Product{
		{SKU: "000001", Name: "BV Lean leather ankle boots", Category: "boots", Price: 89000},
		{SKU: "000002", Name: "BV Lean leather ankle boots", Category: "boots", Price: 99000},
		{SKU: "000003", Name: "Ashlington leather ankle boots", Category: "boots", Price: 71000},
		{SKU: "000004", Name: "Naima embellished suede sandals", Category: "sandals", Price: 79500},
		{SKU: "000005", Name: "Nathane leather sneakers", Category: "sneakers", Price: 59000},
	}
	db.Create(&products)
}
