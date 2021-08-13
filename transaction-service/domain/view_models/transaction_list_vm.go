package view_models

import (
	"github.com/ecommerce-service/transaction-service/domain/models"
	"time"
)

type TransactionListVm struct {
	ID                string  `json:"id"`
	UserID            string  `json:"user_id"`
	TransactionNumber string  `json:"transaction_number"`
	Status            string  `json:"status"`
	Total             float64 `json:"total"`
	Discount          float64 `json:"discount"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
	PaidAt            string  `json:"paid_at"`
	CanceledAt        string  `json:"canceled_at"`
}

func NewTransactionListVm(model *models.Transactions) TransactionListVm {
	return TransactionListVm{
		ID:                model.Id(),
		UserID:            model.UserId(),
		TransactionNumber: model.TransactionNumber(),
		Status:            model.Status(),
		Total:             model.Total(),
		Discount:          model.Discount().Float64,
		CreatedAt:         model.CreatedAt().Format(time.RFC3339),
		UpdatedAt:         model.UpdatedAt().Format(time.RFC3339),
		PaidAt:            model.PaidAt().Time.Format(time.RFC3339),
		CanceledAt:        model.CanceledAt().Time.Format(time.RFC3339),
	}
}
