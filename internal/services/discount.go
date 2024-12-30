package services

import (
	"dgoo/internal/models"
	"fmt"
	"math"
)

func ApplyDiscounts(products []models.Product) []models.ProductResponse {
	var responses []models.ProductResponse
	for _, p := range products {
		discount := 0

		if p.Category == "boots" {
			discount = 30
		}

		if p.SKU == "000003" && 15 > discount {
			discount = 15
		}

		var finalPrice int
		var discountStr *string
		if discount > 0 {
			finalPrice = p.Price - int(math.Round(float64(p.Price)*float64(discount)/100))
			d := fmt.Sprintf("%d%%", discount)
			discountStr = &d
		} else {
			finalPrice = p.Price
			discountStr = nil
		}

		responses = append(responses, models.ProductResponse{
			SKU:      p.SKU,
			Name:     p.Name,
			Category: p.Category,
			Price: models.PriceDetail{
				Original:           p.Price,
				Final:              finalPrice,
				DiscountPercentage: discountStr,
				Currency:           "EUR",
			},
		})
	}
	return responses
}
