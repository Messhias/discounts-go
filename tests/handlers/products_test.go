package handlers_test

import (
	"dgoo/internal/handlers"
	"dgoo/internal/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) GetProducts(category string, priceLessThan *int, limit int) ([]models.Product, error) {
	args := m.Called(category, priceLessThan, limit)

	if args.Get(0) != nil {
		return args.Get(0).([]models.Product), args.Error(1)
	}
	return []models.Product{}, args.Error(1)
}

func TestGetProducts_Success(t *testing.T) {

	mockRepo := new(MockProductRepo)
	handler := handlers.NewProductHandler(mockRepo)

	products := []models.Product{
		{SKU: "000001", Name: "BV Lean leather ankle boots", Category: "boots", Price: 89000},
	}
	mockRepo.On("GetProducts", "boots", (*int)(nil), 5).Return(products, nil)

	router := gin.Default()
	router.GET("/products", handler.GetProducts)

	req, _ := http.NewRequest("GET", "/products?category=boots", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string][]models.ProductResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response["products"], 1)
	assert.Equal(t, "000001", response["products"][0].SKU)
	mockRepo.AssertExpectations(t)
}

func TestGetProducts_InvalidPrice(t *testing.T) {

	mockRepo := new(MockProductRepo)
	handler := handlers.NewProductHandler(mockRepo)

	router := gin.Default()
	router.GET("/products", handler.GetProducts)

	req, _ := http.NewRequest("GET", "/products?category=boots&priceLessThan=invalid", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid priceLessThan parameter", response["error"])

	mockRepo.AssertNotCalled(t, "GetProducts", mock.Anything, mock.Anything, mock.Anything)
}

func TestGetProducts_NoFilters(t *testing.T) {

	mockRepo := new(MockProductRepo)
	handler := handlers.NewProductHandler(mockRepo)

	router := gin.Default()
	router.GET("/products", handler.GetProducts)

	req, _ := http.NewRequest("GET", "/products?category=boots", nil)
	w := httptest.NewRecorder()

	products := []models.Product{
		{SKU: "000001", Name: "BV Lean leather ankle boots", Category: "boots", Price: 89000},
		{SKU: "000002", Name: "BV Lean leather ankle boots", Category: "boots", Price: 99000},
	}
	mockRepo.On("GetProducts", "boots", (*int)(nil), 5).Return(products, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string][]models.ProductResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response["products"], 2)
	mockRepo.AssertExpectations(t)
}

func TestGetProducts_ServerError(t *testing.T) {

	mockRepo := new(MockProductRepo)
	handler := handlers.NewProductHandler(mockRepo)

	mockRepo.On("GetProducts", "boots", (*int)(nil), 5).Return(nil, errors.New("database error"))

	router := gin.Default()
	router.GET("/products", handler.GetProducts)

	req, _ := http.NewRequest("GET", "/products?category=boots", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Error fetching products", response["error"])
	mockRepo.AssertExpectations(t)
}
