package models

import (
	"database/sql"
	"time"
)

type Carts struct {
	id        string
	userId    string
	productId string
	price     float64
	quantity  int64
	subTotal  float64
	name      string
	sku       string
	category  string
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func (model *Carts) Id() string {
	return model.id
}

func (model *Carts) SetId(id string) *Carts {
	model.id = id
	return model
}

func (model *Carts) UserId() string {
	return model.userId
}

func (model *Carts) SetUserId(userId string) *Carts {
	model.userId = userId
	return model
}

func (model *Carts) ProductId() string {
	return model.productId
}

func (model *Carts) SetProductId(productId string) *Carts {
	model.productId = productId
	return model
}

func (model *Carts) Name() string {
	return model.name
}

func (model *Carts) SetName(name string) *Carts {
	model.name = name
	return model
}

func (model *Carts) Sku() string {
	return model.sku
}

func (model *Carts) SetSku(sku string) *Carts {
	model.sku = sku
	return model
}

func (model *Carts) Category() string {
	return model.category
}

func (model *Carts) SetCategory(category string) *Carts {
	model.category = category
	return model
}

func (model *Carts) Price() float64 {
	return model.price
}

func (model *Carts) SetPrice(price float64) *Carts {
	model.price = price
	return model
}

func (model *Carts) Quantity() int64 {
	return model.quantity
}

func (model *Carts) SetQuantity(quantity int64) *Carts {
	model.quantity = quantity
	return model
}

func (model *Carts) SubTotal() float64 {
	return model.subTotal
}

func (model *Carts) SetSubTotal(subTotal float64) *Carts {
	model.subTotal = subTotal
	return model
}

func (model *Carts) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Carts) SetCreatedAt(createdAt time.Time) *Carts {
	model.createdAt = createdAt
	return model
}

func (model *Carts) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Carts) SetUpdatedAt(updatedAt time.Time) *Carts {
	model.updatedAt = updatedAt
	return model
}

func (model *Carts) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Carts) SetDeletedAt(deletedAt sql.NullTime) *Carts {
	model.deletedAt = deletedAt
	return model
}

func NewCartModel() *Carts {
	return &Carts{}
}

const (
	CartSelectStatement       = `SELECT c.id, c.user_id, c.product_id, c.name, c.sku, c.category, c.price, c.quantity, c.sub_total, c.created_at, c.updated_at FROM carts c `
	CartCountSelectStatement  = `SELECT count(c.id) FROM carts c `
	CartDefaultWhereStatement = `WHERE c.deleted_at IS NULL `
)

func (model *Carts) ScanRows(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&model.id, &model.userId, &model.productId, &model.name, &model.sku, &model.category, &model.price, &model.quantity, &model.subTotal,
		&model.createdAt, &model.updatedAt)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Carts) ScanRow(row *sql.Row) (interface{}, error) {
	err := row.Scan(&model.id, &model.userId, &model.name, &model.sku, &model.category, &model.productId, &model.price, &model.quantity, &model.subTotal,
		&model.createdAt, &model.updatedAt)
	if err != nil {
		return model, err
	}

	return model, nil
}
