package services_test

import (
	"dgoo/internal/models"
	"dgoo/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyDiscounts(t *testing.T) {
	products := []models.Product{
		{SKU: "000001", Name: "BV Lean leather ankle boots", Category: "boots", Price: 89000},
		{SKU: "000002", Name: "BV Lean leather ankle boots", Category: "boots", Price: 99000},
		{SKU: "000003", Name: "Ashlington leather ankle boots", Category: "boots", Price: 71000},
		{SKU: "000004", Name: "Naima embellished suede sandals", Category: "sandals", Price: 79500},
		{SKU: "000005", Name: "Nathane leather sneakers", Category: "sneakers", Price: 59000},
	}

	expected := []models.ProductResponse{
		{
			SKU:      "000001",
			Name:     "BV Lean leather ankle boots",
			Category: "boots",
			Price: models.PriceDetail{
				Original:           89000,
				Final:              62300,
				DiscountPercentage: stringPtr("30%"),
				Currency:           "EUR",
			},
		},
		{
			SKU:      "000002",
			Name:     "BV Lean leather ankle boots",
			Category: "boots",
			Price: models.PriceDetail{
				Original:           99000,
				Final:              69300,
				DiscountPercentage: stringPtr("30%"),
				Currency:           "EUR",
			},
		},
		{
			SKU:      "000003",
			Name:     "Ashlington leather ankle boots",
			Category: "boots",
			Price: models.PriceDetail{
				Original:           71000,
				Final:              49700,
				DiscountPercentage: stringPtr("30%"),
				Currency:           "EUR",
			},
		},
		{
			SKU:      "000004",
			Name:     "Naima embellished suede sandals",
			Category: "sandals",
			Price: models.PriceDetail{
				Original:           79500,
				Final:              79500,
				DiscountPercentage: nil,
				Currency:           "EUR",
			},
		},
		{
			SKU:      "000005",
			Name:     "Nathane leather sneakers",
			Category: "sneakers",
			Price: models.PriceDetail{
				Original:           59000,
				Final:              59000,
				DiscountPercentage: nil,
				Currency:           "EUR",
			},
		},
	}

	result := services.ApplyDiscounts(products)

	assert.Equal(t, expected, result)
}

func stringPtr(s string) *string {
	return &s
}
