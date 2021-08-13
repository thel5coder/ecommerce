package requests

type ProductRequest struct {
	CategoryId    string   `json:"category_id" validate:"required"`
	Name          string   `json:"name" validate:"required"`
	Sku           string   `json:"sku" validate:"required"`
	Price         float64  `json:"price" validate:"required"`
	Discount      float64  `json:"discount"`
	MainImage     string   `json:"main_image"`
	ProductImages []string `json:"product_images"`
}
