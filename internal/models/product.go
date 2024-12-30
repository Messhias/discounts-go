package models

type Product struct {
	SKU      string `json:"sku" gorm:"primaryKey"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

type PriceDetail struct {
	Original           int     `json:"original"`
	Final              int     `json:"final"`
	DiscountPercentage *string `json:"discount_percentage,omitempty"`
	Currency           string  `json:"currency"`
}

type ProductResponse struct {
	SKU      string      `json:"sku"`
	Name     string      `json:"name"`
	Category string      `json:"category"`
	Price    PriceDetail `json:"price"`
}
