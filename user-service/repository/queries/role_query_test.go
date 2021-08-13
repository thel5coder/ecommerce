package queries

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ecommerce-service/user-service/domain/models"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func NewSqlMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestRoleQuery_BrowseAll(t *testing.T) {
	db, mock := NewSqlMock()
	defer func() {
		db.Close()
	}()

	search := "%%"
	color := models.NewRoleModel().SetId(gofakeit.Int64()).SetName(gofakeit.Name())
	repository := NewRoleQuery(db)

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(color.Id(), color.Name())
	statement := `SELECT id,name FROM roles WHERE LOWER(name) LIKE $1 ORDER BY id ASC`
	mock.ExpectQuery(statement).WithArgs(search).WillReturnRows(rows)
	res, err := repository.BrowseAll("")

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
