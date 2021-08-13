package queries

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ecommerce-service/transaction-service/domain/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCartQuery_BrowseByUser(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	userId := gofakeit.UUID()
	model := models.NewCartModel().SetId(gofakeit.UUID()).SetUserId(gofakeit.UUID()).SetProductId(gofakeit.UUID()).
		SetName(gofakeit.Name()).SetSku(gofakeit.UUID()).
		SetCategory("{\n\t\tType:     \"Ini jso\",\n\t\tRowCount: 0,\n\t\tFields:   nil,\n\t\tIndent:   false,\n\t}").
		SetPrice(gofakeit.Price(100000000, 130000000)).SetQuantity(gofakeit.Int64()).
		SetSubTotal(gofakeit.Price(100000000, 130000000)).SetCreatedAt(now).SetUpdatedAt(now)
	repository := CartQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"id", "c.user_id", "c.product_id", "c.name", "c.sku", "c.category", "c.price", "c.quantity", "c.subtotal","c.created_at", "c.updated_at"}).
		AddRow(model.Id(), model.UserId(), model.ProductId(), model.Name(), model.Sku(), model.Category(), model.Price(), model.Quantity(), model.SubTotal(), model.CreatedAt(), model.UpdatedAt())
	statement := `SELECT c.id, c.user_id, c.product_id, c.name, c.sku, c.category, c.price, c.quantity, c.sub_total, c.created_at, c.updated_at FROM carts c ` +
		`WHERE c.deleted_at IS NULL AND c.user_id=$1 ` +
		`ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(userId, 10, 0).WillReturnRows(rows)
	res, err := repository.BrowseByUser("", "created_at", "desc", userId, 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCartQuery_BrowseAllByUser(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	userId := gofakeit.UUID()
	model := models.NewCartModel().SetId(gofakeit.UUID()).SetUserId(gofakeit.UUID()).SetProductId(gofakeit.UUID()).
		SetName(gofakeit.Name()).SetSku(gofakeit.UUID()).
		SetCategory("{\n\t\tType:     \"Ini jso\",\n\t\tRowCount: 0,\n\t\tFields:   nil,\n\t\tIndent:   false,\n\t}").
		SetPrice(gofakeit.Price(100000000, 130000000)).SetQuantity(gofakeit.Int64()).
		SetSubTotal(gofakeit.Price(100000000, 130000000)).SetCreatedAt(now).SetUpdatedAt(now)
	repository := CartQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"id", "c.user_id", "c.product_id", "c.name", "c.sku", "c.category", "c.price", "c.quantity", "c.subtotal","c.created_at", "c.updated_at"}).
		AddRow(model.Id(), model.UserId(), model.ProductId(), model.Name(), model.Sku(), model.Category(), model.Price(), model.Quantity(), model.SubTotal(), model.CreatedAt(), model.UpdatedAt())
	statement := `SELECT c.id, c.user_id, c.product_id, c.name, c.sku, c.category, c.price, c.quantity, c.sub_total, c.created_at, c.updated_at FROM carts c ` +
		`WHERE c.deleted_at IS NULL AND c.user_id=$1 `
	mock.ExpectQuery(statement).WithArgs(userId).WillReturnRows(rows)
	res, err := repository.BrowseAllByUser(userId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCartQuery_ReadBy(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	userId := gofakeit.UUID()
	id := gofakeit.UUID()
	model := models.NewCartModel().SetId(gofakeit.UUID()).SetUserId(gofakeit.UUID()).SetProductId(gofakeit.UUID()).
		SetName(gofakeit.Name()).SetSku(gofakeit.UUID()).
		SetCategory("{\n\t\tType:     \"Ini jso\",\n\t\tRowCount: 0,\n\t\tFields:   nil,\n\t\tIndent:   false,\n\t}").
		SetPrice(gofakeit.Price(100000000, 130000000)).SetQuantity(gofakeit.Int64()).
		SetSubTotal(gofakeit.Price(100000000, 130000000)).SetCreatedAt(now).SetUpdatedAt(now)
	repository := CartQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"id", "c.user_id", "c.product_id", "c.name", "c.sku", "c.category", "c.price", "c.quantity", "c.subtotal","c.created_at", "c.updated_at"}).
		AddRow(model.Id(), model.UserId(), model.ProductId(), model.Name(), model.Sku(), model.Category(), model.Price(), model.Quantity(), model.SubTotal(), model.CreatedAt(), model.UpdatedAt())

	statement := `SELECT c.id, c.user_id, c.product_id, c.name, c.sku, c.category, c.price, c.quantity, c.sub_total, c.created_at, c.updated_at FROM carts c ` +
		`WHERE c.deleted_at IS NULL AND c.user_id=$1 AND c.id=$2`
	mock.ExpectQuery(statement).WithArgs(userId, id).WillReturnRows(rows)
	res, err := repository.ReadBy("c.id", "=", userId, id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCartQuery_Count(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	count := 1
	userId := gofakeit.UUID()
	repository := CartQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT count(c.id) FROM carts c WHERE c.deleted_at IS NULL AND c.user_id=$1`
	mock.ExpectQuery(statement).WithArgs(userId).WillReturnRows(rows)
	res, err := repository.Count("", userId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestCartQuery_CountBy(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	count := 1
	productId := gofakeit.UUID()
	userId := gofakeit.UUID()
	repository := CartQueryMock{db: db}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT count(c.id) FROM carts c WHERE c.deleted_at IS NULL AND c.user_id=$1 AND c.product_id=$2`
	mock.ExpectQuery(statement).WithArgs(userId,productId).WillReturnRows(rows)
	res, err := repository.CountBy("c.car_id", "=", userId, productId)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
