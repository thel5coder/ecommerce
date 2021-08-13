package requests

type CartRequest struct {
	ProductID string  `json:"product_id" validate:"required"`
	Name      string  `json:"name"`
	Sku       string  `json:"sku"`
	Category  string  `json:"category"`
	Price     float64 `json:"price" validate:"required"`
	Quantity  int64   `json:"quantity" validate:"required"`
}
