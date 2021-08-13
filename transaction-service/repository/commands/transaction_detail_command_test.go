package commands

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ecommerce-service/transaction-service/domain/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddTransactionDetail(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	quantity := int64(gofakeit.Int8())
	price := gofakeit.Price(100000000, 130000000)
	subTotal := float64(quantity) * price
	model := models.NewTransactionDetailModel().SetTransactionId(gofakeit.UUID()).SetName(gofakeit.Name()).SetSku(gofakeit.UUID()).
		SetCategory("{\n\t\tType:     \"Ini jso\",\n\t\tRowCount: 0,\n\t\tFields:   nil,\n\t\tIndent:   false,\n\t}").
		SetPrice(price).SetQuantity(quantity).
		SetSubTotal(subTotal).SetCreatedAt(now).SetUpdatedAt(now)

	cmd := TransactionDetailCommandMock{
		db:    db,
		model: model,
	}
	statement := `INSERT INTO transaction_details(transaction_id,name,sku,category,price,discount,quantity,sub_total,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	mock.ExpectBegin()
	mock.ExpectExec(statement).WithArgs(model.TransactionId(), model.Name(), model.Sku(), model.Category(), model.Price(), model.Discount(),
		model.Quantity(), model.SubTotal(), model.CreatedAt(), model.UpdatedAt()).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	err := cmd.Add()

	assert.NoError(t, err)
}

func TestAddTransactionDetailError(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	quantity := int64(gofakeit.Int8())
	price := gofakeit.Price(100000000, 130000000)
	subTotal := float64(quantity) * price
	model := models.NewTransactionDetailModel().SetName(gofakeit.Name()).SetSku(gofakeit.UUID()).
		SetCategory("{\n\t\tType:     \"Ini jso\",\n\t\tRowCount: 0,\n\t\tFields:   nil,\n\t\tIndent:   false,\n\t}").
		SetPrice(price).SetQuantity(quantity).
		SetSubTotal(subTotal).SetCreatedAt(now).SetUpdatedAt(now)

	cmd := TransactionDetailCommandMock{
		db:    db,
		model: model,
	}
	statement := `INSERT INTO transaction_details(transaction_id,name,sku,category,price,discount,quantity,sub_total,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	mock.ExpectBegin()
	mock.ExpectExec(statement).WithArgs(model.TransactionId(), model.Name(), model.Sku(), model.Category(), model.Price(), model.Discount(),
		model.Quantity(), model.SubTotal(), model.CreatedAt(), model.UpdatedAt()).WillReturnResult(sqlmock.NewResult(0, 0)).
		WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()
	err := cmd.Add()

	assert.Error(t, err)
}
