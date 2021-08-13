package view_models

import "github.com/thel5coder/pkg/str"

type TransactionDetailListVm struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Sku           string  `json:"sku"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	Discount      float64 `json:"discount"`
	Quantity      int     `json:"quantity"`
	SubTotal      float64 `json:"sub_total"`
}

func NewTransactionDetailListVm(transactionDetails []string) TransactionDetailListVm {
	return TransactionDetailListVm{
		ID:            transactionDetails[0],
		Name:          transactionDetails[1],
		Sku:           transactionDetails[2],
		Category:      transactionDetails[3],
		Price:         float64(str.StringToInt(transactionDetails[4])),
		Discount:      float64(str.StringToInt(transactionDetails[5])),
		Quantity:      str.StringToInt(transactionDetails[6]),
		SubTotal:      float64(str.StringToInt(transactionDetails[7])),
	}
}
