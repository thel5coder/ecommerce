package view_models

import (
	"github.com/ecommerce/product-service/domain/models"
	"time"
)

type DetailProductVm struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Sku           string   `json:"sku"`
	Price         float64  `json:"price"`
	Discount      float64  `json:"discount"`
	Stock         int      `json:"stock"`
	MainImage     FileVm   `json:"main_image"`
	ProductImages []FileVm `json:"product_images"`
	Category      string   `json:"category"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
}

func NewDetailProductVm(model *models.Product, mainImage FileVm, productImages []FileVm) DetailProductVm {
	return DetailProductVm{
		ID:            model.Id(),
		Name:          model.Name(),
		Sku:           model.Sku(),
		Price:         model.Price(),
		Discount:      model.Discount().Float64,
		Stock:         model.Stock(),
		MainImage:     mainImage,
		ProductImages: productImages,
		Category:      model.Category.Name(),
		CreatedAt:     model.CreatedAt().Format(time.RFC3339),
		UpdatedAt:     model.UpdatedAt().Format(time.RFC3339),
	}
}
