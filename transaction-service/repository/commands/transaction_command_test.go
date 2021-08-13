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

func TestAddTransaction(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewTransactionModel().SetUserId(gofakeit.UUID()).SetTransactionNumber(gofakeit.UUID()).SetStatus("on_going").
		SetTotal(gofakeit.Price(1000000000, 120000000)).SetDiscount(sql.NullFloat64{Float64: gofakeit.Price(1000000000, 120000000), Valid: true}).
		SetCreatedAt(now).SetUpdatedAt(now)
	id := gofakeit.UUID()

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `INSERT INTO transactions (user_id,transaction_number,status,total,discount,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	mock.ExpectBegin()
	mock.ExpectQuery(statement).WithArgs(model.UserId(), model.TransactionNumber(), model.Status(), model.Total(), model.Discount(),
		model.CreatedAt(), model.UpdatedAt()).WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := cmd.Add()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestAddTransactionError(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewTransactionModel().SetStatus("on_going").SetTransactionNumber(gofakeit.UUID()).SetCreatedAt(now).SetUpdatedAt(now)

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	statement := `INSERT INTO transactions (user_id,transaction_number,status,total,discount,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	mock.ExpectBegin()
	mock.ExpectQuery(statement).WithArgs(model.UserId(), model.TransactionNumber(), model.Status(), model.Total(), model.Discount(),
		model.CreatedAt(), model.UpdatedAt()).WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()
	res, err := cmd.Add()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestEditPaymentReceived(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewTransactionModel().SetStatus("success").
		SetUpdatedAt(now).SetPaidAt(now).SetId(id)

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	statement := `UPDATE transactions set updated_at=$1,paid_at=$2 WHERE id=$3`
	mock.ExpectBegin()
	mock.ExpectExec(statement).WithArgs(model.UpdatedAt(), model.PaidAt().Time, model.Id()).
		WillReturnResult(sqlmock.NewResult(0,1))
	mock.ExpectCommit()
	err := cmd.EditPaymentReceived()

	assert.NoError(t, err)
}

func TestEditPaymentReceivedError(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewTransactionModel().SetStatus("success").
		SetUpdatedAt(now).SetPaidAt(now)

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	statement := `UPDATE transactions set updated_at=$1,paid_at=$2 WHERE id=$3`
	mock.ExpectBegin()
	mock.ExpectExec(statement).WithArgs(model.UpdatedAt(), model.PaidAt().Time, model.Id()).
		WillReturnResult(sqlmock.NewResult(0,0)).WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()
	err := cmd.EditPaymentReceived()

	assert.Error(t, err)
}

func TestPaymentCancel(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewTransactionModel().SetStatus("canceled").SetUpdatedAt(now).SetCanceledAt(now).SetId(id)

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE transactions set updated_at=$1,canceled_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectBegin()
	mock.ExpectQuery(statement).WithArgs(model.Status(),model.UpdatedAt(),model.CanceledAt().Time,model.Id()).WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := cmd.EditCancelPayment()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestPaymentCancelError(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewTransactionModel().SetStatus("canceled").SetUpdatedAt(now).SetCanceledAt(now)

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	statement := `UPDATE transactions set updated_at=$1,canceled_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectBegin()
	mock.ExpectQuery(statement).WithArgs(model.Status(),model.UpdatedAt(),model.CanceledAt().Time,model.Id()).WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()
	res, err := cmd.EditCancelPayment()

	assert.Error(t, err)
	assert.Empty(t, res)
}
