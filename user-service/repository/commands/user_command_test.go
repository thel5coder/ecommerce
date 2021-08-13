package commands

import (
	"booking-car/domain/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserCommand_Add(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	password := gofakeit.Password(true,false,true,false,false,6)
	model := models.NewUserModel().SetId(gofakeit.UUID()).SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).SetPassword(password).
		SetAddress(gofakeit.Address().Address).SetPhoneNumber(gofakeit.Phone()).SetRoleId(2).SetCreatedAt(now).SetUpdatedAt(now)

	cmd := UserCommandMock{
		db:    db,
		model: model,
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(model.Id())
	statement := `INSERT INTO users (first_name,last_name,email,username,password,address,phone_number,role_id,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.FirstName(),model.LastName(),model.Email(),model.UserName(),model.Password(),model.Address().String,
		model.PhoneNumber(),model.RoleId(),model.CreatedAt(),model.UpdatedAt()).WillReturnRows(rows)
	res, err := cmd.Add()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUserCommand_EditWithOutPassword(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewUserModel().SetId(gofakeit.UUID()).SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).
		SetAddress(gofakeit.Address().Address).SetPhoneNumber(gofakeit.Phone()).SetRoleId(2).SetUpdatedAt(now)

	cmd := UserCommandMock{
		db:    db,
		model: model,
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(model.Id())
	statement := `UPDATE users SET first_name=$1,last_name=$2,email=$3,username=$4,address=$5,phone_number=$6,role_id=$7,updated_at=$8 WHERE id=$9 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.FirstName(),model.LastName(),model.Email(),model.UserName(),model.Address().String,
		model.PhoneNumber(),model.RoleId(),model.UpdatedAt(),model.Id()).WillReturnRows(rows)
	res, err := cmd.Edit()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUserCommand_EditWithPassword(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	password := gofakeit.Password(true,false,true,false,false,6)
	model := models.NewUserModel().SetId(gofakeit.UUID()).SetFirstName(gofakeit.FirstName()).SetLastName(gofakeit.LastName()).SetEmail(gofakeit.Email()).SetPassword(password).
		SetAddress(gofakeit.Address().Address).SetPhoneNumber(gofakeit.Phone()).SetRoleId(2).SetUpdatedAt(now)

	cmd := UserCommandMock{
		db:    db,
		model: model,
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(model.Id())
	statement := `UPDATE users SET first_name=$1,last_name=$2,email=$3,username=$4,address=$5,phone_number=$6,role_id=$7,updated_at=$8,password=$10 WHERE id=$9 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.FirstName(),model.LastName(),model.Email(),model.UserName(),model.Address().String,
		model.PhoneNumber(),model.RoleId(),model.UpdatedAt(),model.Id(),model.Password()).WillReturnRows(rows)
	res, err := cmd.Edit()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUserCommand_Delete(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewUserModel().SetId(gofakeit.UUID()).SetUpdatedAt(now).SetDeletedAt(now)

	cmd := UserCommandMock{
		db:    db,
		model: model,
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(model.Id())
	statement := `UPDATE users SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	mock.ExpectQuery(statement).WithArgs(model.UpdatedAt(),model.DeletedAt().Time,model.Id()).WillReturnRows(rows)
	res, err := cmd.Delete()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUserCommand_EditDeposit(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewUserModel().SetId(gofakeit.UUID()).SetDepositAmount(gofakeit.Price(30000000,50000000)).SetUpdatedAt(now)

	cmd := UserCommandMock{
		db:    db,
		model: model,
	}
	statement := `UPDATE users SET deposit_amount=$1,updated_at=$2 WHERE id=$3`
	mock.ExpectBegin()
	mock.ExpectExec(statement).WithArgs(model.DepositAmount().Float64,model.UpdatedAt(),model.Id()).WillReturnResult(sqlmock.NewResult(0,1))
	mock.ExpectCommit()
	err := cmd.EditDeposit()

	assert.NoError(t, err)
}
