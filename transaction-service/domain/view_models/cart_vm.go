package view_models

import (
	"github.com/ecommerce-service/transaction-service/domain/models"
	"time"
)

type CartVm struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Sku       string  `json:"sku"`
	Category  string  `json:"category"`
	Price     float64 `json:"price"`
	Quantity  int64   `json:"quantity"`
	SubTotal  float64 `json:"sub_total"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func NewCartVm(model *models.Carts) CartVm {
	return CartVm{
		ID:        model.Id(),
		UserID:    model.UserId(),
		ProductID: model.ProductId(),
		Name:      model.Name(),
		Sku:       model.Sku(),
		Category:  model.Category(),
		Price:     model.Price(),
		Quantity:  model.Quantity(),
		SubTotal:  model.SubTotal(),
		CreatedAt: model.CreatedAt().Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt().Format(time.RFC3339),
	}
}
