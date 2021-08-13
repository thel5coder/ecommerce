package models

import (
	"database/sql"
	"time"
)

type TransactionDetails struct {
	id            string
	transactionId string
	name          string
	sku           string
	category      string
	price         float64
	discount      sql.NullFloat64
	quantity      int64
	subTotal      float64
	createdAt     time.Time
	updatedAt     time.Time
	deletedAt     sql.NullTime
}

func NewTransactionDetailModel() *TransactionDetails {
	return &TransactionDetails{}
}

func (model *TransactionDetails) Id() string {
	return model.id
}

func (model *TransactionDetails) SetId(id string) *TransactionDetails {
	model.id = id

	return model
}

func (model *TransactionDetails) TransactionId() string {
	return model.transactionId
}

func (model *TransactionDetails) SetTransactionId(transactionId string) *TransactionDetails {
	model.transactionId = transactionId

	return model
}

func (model *TransactionDetails) Name() string {
	return model.name
}

func (model *TransactionDetails) SetName(name string) *TransactionDetails {
	model.name = name

	return model
}

func (model *TransactionDetails) Sku() string {
	return model.sku
}

func (model *TransactionDetails) SetSku(sku string) *TransactionDetails {
	model.sku = sku

	return model
}

func (model *TransactionDetails) Category() string {
	return model.category
}

func (model *TransactionDetails) SetCategory(category string) *TransactionDetails {
	model.category = category

	return model
}

func (model *TransactionDetails) Price() float64 {
	return model.price
}

func (model *TransactionDetails) SetPrice(price float64) *TransactionDetails {
	model.price = price

	return model
}

func (model *TransactionDetails) Discount() sql.NullFloat64 {
	return model.discount
}

func (model *TransactionDetails) SetDiscount(discount sql.NullFloat64) *TransactionDetails {
	model.discount = discount

	return model
}

func (model *TransactionDetails) Quantity() int64 {
	return model.quantity
}

func (model *TransactionDetails) SetQuantity(quantity int64) *TransactionDetails {
	model.quantity = quantity

	return model
}

func (model *TransactionDetails) SubTotal() float64 {
	return model.subTotal
}

func (model *TransactionDetails) SetSubTotal(subTotal float64) *TransactionDetails {
	model.subTotal = subTotal

	return model
}

func (model *TransactionDetails) CreatedAt() time.Time {
	return model.createdAt
}

func (model *TransactionDetails) SetCreatedAt(createdAt time.Time) *TransactionDetails {
	model.createdAt = createdAt

	return model
}

func (model *TransactionDetails) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *TransactionDetails) SetUpdatedAt(updatedAt time.Time) *TransactionDetails {
	model.updatedAt = updatedAt

	return model
}

func (model *TransactionDetails) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *TransactionDetails) SetDeletedAt(deletedAt time.Time) *TransactionDetails {
	model.deletedAt.Time = deletedAt

	return model
}