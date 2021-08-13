package view_models

import (
	"github.com/ecommerce/product-service/domain/models"
	"time"
)

type ListProductVm struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Sku       string  `json:"sku"`
	Price     float64 `json:"price"`
	Discount  float64 `json:"discount"`
	Stock     int     `json:"stock"`
	MainImage FileVm  `json:"main_image"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func NewListProductVm(model *models.Product, file FileVm) ListProductVm {
	return ListProductVm{
		ID:        model.Id(),
		Name:      model.Name(),
		Sku:       model.Sku(),
		Price:     model.Price(),
		Discount:  model.Discount().Float64,
		Stock:     model.Stock(),
		MainImage: file,
		CreatedAt: model.CreatedAt().Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt().Format(time.RFC3339),
	}
}
