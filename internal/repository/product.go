package repository

import (
	"dgoo/internal/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts(category string, priceLessThan *int, limit int) ([]models.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db}
}

func (r *productRepo) GetProducts(category string, priceLessThan *int, limit int) ([]models.Product, error) {
	var products []models.Product
	query := r.db

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if priceLessThan != nil {
		query = query.Where("price <= ?", *priceLessThan)
	}

	query = query.Limit(limit)

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
