package handlers

import (
	"dgoo/internal/repository"
	"dgoo/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	category := c.Query("category")

	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "category filter is mandatory",
		})
		return
	}

	priceStr := c.Query("priceLessThan")
	var priceLessThan *int

	if priceStr != "" {
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid priceLessThan parameter"})
			return
		}
		priceLessThan = &price
	}

	products, err := h.Repo.GetProducts(category, priceLessThan, 5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}

	response := services.ApplyDiscounts(products)
	c.JSON(http.StatusOK, gin.H{"products": response})
}
