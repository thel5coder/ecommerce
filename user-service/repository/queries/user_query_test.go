package queries

import (
	"booking-car/domain/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserQuery_Browse(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	search := "%%"
	role := models.NewRoleModel().SetId(int64(gofakeit.Int8())).SetName(gofakeit.Name())
	model := models.NewUserModel().SetId(gofakeit.UUID()).SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).
		SetPassword(gofakeit.Password(true, false, true, false, false, 6)).
		SetUserName(gofakeit.Username()).SetAddress(gofakeit.Address().Address).SetPhoneNumber(gofakeit.Phone()).SetDepositAmount(0).SetCreatedAt(now).
		SetUpdatedAt(now)
	model.Role = role
	repository := NewUserQuery(db)

	rows := sqlmock.NewRows([]string{"u.id", "u.first_name", "u.last_name", "u.email", "u.username", "u.password", "u.address", "u.phone_number", "u.deposit_amount",
		"u.created_at", "u.updated_at", "r.id", "r.name"}).AddRow(model.Id(), model.FirstName(), model.LastName(), model.Email(), model.UserName(), model.Password(), model.Address(),
		model.PhoneNumber(), model.DepositAmount(), model.CreatedAt(), model.UpdatedAt(), model.Role.Id(), model.Role.Name())
	statement := `SELECT u.id,u.first_name,u.last_name,u.email,u.username,u.password,u.address,u.phone_number,u.deposit_amount,u.created_at,u.updated_at,r.id,r.name FROM users u ` +
		`INNER JOIN roles r ON r.id = u.role_id WHERE u.deleted_at IS NULL AND (LOWER(u.first_name) LIKE $1 OR LOWER(last_name) LIKE $1 OR LOWER(email) LIKE $1 OR phone_number LIKE $1) ` +
		`ORDER BY created_at desc LIMIT $2 OFFSET $3`
	mock.ExpectQuery(statement).WithArgs(search, 10, 0).WillReturnRows(rows)
	res, err := repository.Browse("", "created_at", "desc", 10, 0)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUserQuery_ReadBy(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	role := models.NewRoleModel().SetId(int64(gofakeit.Int8())).SetName(gofakeit.Name())
	model := models.NewUserModel().SetId(gofakeit.UUID()).SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).
		SetPassword(gofakeit.Password(true, false, true, false, false, 6)).
		SetUserName(gofakeit.Username()).SetAddress(gofakeit.Address().Address).SetPhoneNumber(gofakeit.Phone()).SetDepositAmount(0).SetCreatedAt(now).
		SetUpdatedAt(now)
	model.Role = role
	repository := NewUserQuery(db)

	rows := sqlmock.NewRows([]string{"u.id", "u.first_name", "u.last_name", "u.email", "u.username", "u.password", "u.address", "u.phone_number", "u.deposit_amount",
		"u.created_at", "u.updated_at", "r.id", "r.name"}).AddRow(model.Id(), model.FirstName(), model.LastName(), model.Email(), model.UserName(), model.Password(), model.Address(),
		model.PhoneNumber(), model.DepositAmount(), model.CreatedAt(), model.UpdatedAt(), model.Role.Id(), model.Role.Name())
	statement := `SELECT u.id,u.first_name,u.last_name,u.email,u.username,u.password,u.address,u.phone_number,u.deposit_amount,u.created_at,u.updated_at,r.id,r.name FROM users u ` +
		`INNER JOIN roles r ON r.id = u.role_id WHERE u.deleted_at IS NULL AND u.id=$1`
	mock.ExpectQuery(statement).WithArgs(id).WillReturnRows(rows)
	res, err := repository.ReadBy("u.id", "=", id)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUserQuery_CountByWithOutID(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	count := 1
	userName := gofakeit.Username()
	repository := NewUserQuery(db)

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(u.id) FROM users u INNER JOIN roles r ON r.id = u.role_id WHERE u.deleted_at IS NULL AND u.username=$1`
	mock.ExpectQuery(statement).WithArgs(userName).WillReturnRows(rows)
	res, err := repository.CountBy("u.username", "=", "", userName)

	assert.NoError(t, err)
	assert.NotZero(t, res)
}

func TestUserQuery_CountByWithID(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	count := 1
	id := gofakeit.UUID()
	userName := gofakeit.Username()
	repository := NewUserQuery(db)

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(u.id) FROM users u INNER JOIN roles r ON r.id = u.role_id WHERE u.deleted_at IS NULL AND u.username=$1 AND u.id<>$2`
	mock.ExpectQuery(statement).WithArgs(userName, id).WillReturnRows(rows)
	res, err := repository.CountBy("u.username", "=", id, userName)

	assert.NoError(t, err)
	assert.NotZero(t, res)
}

func TestUserQuery_Count(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	count := 1
	repository := NewUserQuery(db)

	rows := sqlmock.NewRows([]string{"count"}).AddRow(count)
	statement := `SELECT COUNT(u.id) FROM users u INNER JOIN roles r ON r.id = u.role_id ` +
		`WHERE u.deleted_at IS NULL AND (LOWER(u.first_name) LIKE $1 OR LOWER(last_name) LIKE $1 OR LOWER(email) LIKE $1 OR phone_number LIKE $1)`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.Count("")

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
