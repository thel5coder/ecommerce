package models

import (
	"database/sql"
	"time"
)

type Transactions struct {
	id                string
	userId            string
	transactionNumber string
	status            string
	total             float64
	discount          sql.NullFloat64
	createdAt         time.Time
	updatedAt         time.Time
	paidAt            sql.NullTime
	canceledAt        sql.NullTime
	deletedAt         sql.NullTime
	transactionDetail string
}

func NewTransactionModel() *Transactions {
	return &Transactions{}
}

func (model *Transactions) Id() string {
	return model.id
}

func (model *Transactions) SetId(id string) *Transactions {
	model.id = id

	return model
}

func (model *Transactions) UserId() string {
	return model.userId
}

func (model *Transactions) SetUserId(userId string) *Transactions {
	model.userId = userId

	return model
}

func (model *Transactions) TransactionNumber() string {
	return model.transactionNumber
}

func (model *Transactions) SetTransactionNumber(transactionNumber string) *Transactions {
	model.transactionNumber = transactionNumber

	return model
}

func (model *Transactions) Status() string {
	return model.status
}

func (model *Transactions) SetStatus(status string) *Transactions {
	model.status = status

	return model
}

func (model *Transactions) Total() float64 {
	return model.total
}

func (model *Transactions) SetTotal(total float64) *Transactions {
	model.total = total

	return model
}

func (model *Transactions) Discount() sql.NullFloat64 {
	return model.discount
}

func (model *Transactions) SetDiscount(discount sql.NullFloat64) *Transactions {
	model.discount = discount

	return model
}

func (model *Transactions) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Transactions) SetCreatedAt(createdAt time.Time) *Transactions {
	model.createdAt = createdAt

	return model
}

func (model *Transactions) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Transactions) SetUpdatedAt(updatedAt time.Time) *Transactions {
	model.updatedAt = updatedAt

	return model
}

func (model *Transactions) PaidAt() sql.NullTime {
	return model.paidAt
}

func (model *Transactions) SetPaidAt(paidAt time.Time) *Transactions {
	model.paidAt.Time = paidAt

	return model
}

func (model *Transactions) CanceledAt() sql.NullTime {
	return model.canceledAt
}

func (model *Transactions) SetCanceledAt(canceledAt time.Time) *Transactions {
	model.canceledAt.Time = canceledAt

	return model
}

func (model *Transactions) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Transactions) SetDeletedAt(deletedAt time.Time) *Transactions {
	model.deletedAt.Time = deletedAt

	return model
}

func (model *Transactions) TransactionDetail() string {
	return model.transactionDetail
}

const (
	TransactionSelectListStatement   = `SELECT t.id,t.user_id,t.transaction_number,t.status,t.total,t.discount,t.created_at,t.updated_at,t.paid_at,t.canceled_at `
	TransactionSelectDetailStatement = `,ARRAY_TO_STRING(ARRAY_AGG(td.id ||':'|| COALESCE(td.name,'') ||':'|| COALESCE(td.sku,'') ||':'|| COALESCE(td.category,'') ||':'|| `+
		` COALESCE(td.price,0) ||':'|| COALESCE(td.discount,0) ||':'|| COALESCE(td.quantity,0) ||':'|| COALESCE(td.sub_total,0)),',')`
	TransactionSelectCountStatement  = `SELECT COUNT(t.id) FROM transactions t`
	TransactionDetailJoinStatement   = `INNER JOIN transaction_details td ON td.transaction_id = t.id AND td.deleted_at IS NULL`
	TransactionDefaultWhereStatement = `WHERE t.deleted_at IS NULL`
	TransactionGroupByStatement      = `GROUP BY t.id`
)

func (model *Transactions) ScanRows(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&model.id, &model.userId, &model.transactionNumber, &model.status, &model.total, &model.discount, &model.createdAt, &model.updatedAt,
		&model.paidAt, &model.canceledAt)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Transactions) ScanRow(row *sql.Row) (interface{}, error) {
	err := row.Scan(&model.id, &model.userId, &model.transactionNumber, &model.status, &model.total, &model.discount, &model.createdAt, &model.updatedAt,
		&model.paidAt, &model.canceledAt, &model.transactionDetail)
	if err != nil {
		return model, err
	}

	return model, nil
}
