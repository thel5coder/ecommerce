package commands

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ecommerce-service/transaction-service/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/thel5coder/pkg/postgresql"
	"testing"
	"time"
)

func TestAddCart(t *testing.T) {
	_, mock := NewMock()
	conn, _ := postgresql.NewConnectionMock().Connect()
	db := conn.GetDbInstance()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewCartModel().
		SetUserId(gofakeit.UUID()).
		SetProductId(gofakeit.UUID()).
		SetName(gofakeit.Name()).
		SetSku(gofakeit.UUID()).
		SetCategory("{\n\t\tType:     \"Ini jso\",\n\t\tRowCount: 0,\n\t\tFields:   nil,\n\t\tIndent:   false,\n\t}").
		SetPrice(gofakeit.Price(100000000, 150000000)).
		SetQuantity(gofakeit.Int64()).
		SetSubTotal(gofakeit.Price(100000000, 150000000)).
		SetCreatedAt(now).
		SetUpdatedAt(now)
	id := gofakeit.UUID()

	cmd := NewCartCommand(conn, model)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `INSERT INTO carts(user_id,product_id,name,sku,category,price,quantity,sub_total,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UserId(), model.ProductId(), model.Name(), model.Sku(), model.Category(), model.Price(), model.Quantity(), model.SubTotal(),
		model.CreatedAt(), model.UpdatedAt()).WillReturnRows(rows)
	res, err := cmd.Add()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestAddCartError(t *testing.T) {
	_, mock := NewMock()
	conn, _ := postgresql.NewConnectionMock().Connect()
	db := conn.GetDbInstance()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewCartModel().
		SetUserId(gofakeit.UUID()).
		SetProductId(gofakeit.UUID()).
		SetName(gofakeit.Name()).
		SetSku(gofakeit.UUID()).
		SetCategory("{\n\t\tType:     \"Ini jso\",\n\t\tRowCount: 0,\n\t\tFields:   nil,\n\t\tIndent:   false,\n\t}").
		SetPrice(gofakeit.Price(100000000, 150000000)).
		SetQuantity(gofakeit.Int64()).
		SetSubTotal(gofakeit.Price(100000000, 150000000)).
		SetCreatedAt(now).
		SetUpdatedAt(now)

	cmd := NewCartCommand(conn, model)
	statement := `INSERT INTO carts(user_id,product_id,name,sku,category,price,quantity,sub_total,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UserId(), model.ProductId(), model.Name(), model.Sku(), model.Category(), model.Price(), model.Quantity(), model.SubTotal(),
		model.CreatedAt(), model.UpdatedAt()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Add()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestEditCart(t *testing.T) {
	_, mock := NewMock()
	conn, _ := postgresql.NewConnectionMock().Connect()
	db := conn.GetDbInstance()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewCartModel().
		SetId(id).
		SetQuantity(gofakeit.Int64()).
		SetSubTotal(gofakeit.Price(100000000, 150000000)).
		SetUpdatedAt(now)

	cmd := NewCartCommand(conn, model)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE carts SET quantity=$1,sub_total=$2,updated_at=$3 WHERE id=$4 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.Quantity(), model.SubTotal(), model.UpdatedAt(), model.Id()).WillReturnRows(rows)
	res, err := cmd.Edit()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestEditCartError(t *testing.T) {
	_, mock := NewMock()
	conn, _ := postgresql.NewConnectionMock().Connect()
	db := conn.GetDbInstance()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewCartModel().
		SetQuantity(gofakeit.Int64()).
		SetSubTotal(gofakeit.Price(100000000, 150000000)).
		SetUpdatedAt(now)

	cmd := NewCartCommand(conn, model)
	statement := `UPDATE carts SET quantity=$1,sub_total=$2,updated_at=$3 WHERE id=$4 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.Quantity(), model.SubTotal(), model.UpdatedAt(), model.Id()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Edit()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestDeleteCart(t *testing.T) {
	_, mock := NewMock()
	conn, _ := postgresql.NewConnectionMock().Connect()
	db := conn.GetDbInstance()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewCartModel().SetUpdatedAt(now).SetDeletedAt(sql.NullTime{Time: now, Valid: true}).SetId(id)

	cmd := NewCartCommand(conn, model)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE carts SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(), model.DeletedAt().Time, model.Id()).WillReturnRows(rows)
	res, err := cmd.Delete()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteCartError(t *testing.T) {
	_, mock := NewMock()
	conn, _ := postgresql.NewConnectionMock().Connect()
	db := conn.GetDbInstance()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewCartModel().SetUpdatedAt(now).SetDeletedAt(sql.NullTime{Time: now, Valid: true})

	cmd := NewCartCommand(conn, model)
	statement := `UPDATE carts SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(), model.DeletedAt().Time, model.Id()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Delete()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestDeleteAllByUserID(t *testing.T) {
	_, mock := NewMock()
	conn, _ := postgresql.NewConnectionMock().Connect()
	db := conn.GetDbInstance()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewCartModel().SetUpdatedAt(now).SetDeletedAt(sql.NullTime{Time: now, Valid: true}).SetUserId(id)

	cmd := NewCartCommand(conn, model)
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE carts SET updated_at=$1,deleted_at=$2 WHERE user_id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(), model.DeletedAt().Time, model.UserId()).WillReturnRows(rows)
	res, err := cmd.Delete()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteAllByUserIDError(t *testing.T) {
	_, mock := NewMock()
	conn, _ := postgresql.NewConnectionMock().Connect()
	db := conn.GetDbInstance()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewCartModel().SetUpdatedAt(now).SetDeletedAt(sql.NullTime{Time: now, Valid: true})

	cmd := NewCartCommand(conn, model)
	statement := `UPDATE carts SET updated_at=$1,deleted_at=$2 WHERE user_id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(), model.DeletedAt().Time, model.UserId()).WillReturnError(sql.ErrNoRows)
	res, err := cmd.Delete()

	assert.Error(t, err)
	assert.Empty(t, res)
}
