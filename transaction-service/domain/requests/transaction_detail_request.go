package requests

type TransactionDetailRequest struct {
	Name     string  `json:"name"`
	Sku      string  `json:"sku"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"`
	Quantity int     `json:"quantity"`
	SubTotal float64 `json:"sub_total"`
}
