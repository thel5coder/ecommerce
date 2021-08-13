package requests

type ConfirmPaymentRequest struct {
	PaymentAmount int64 `json:"payment_amount"`
}
