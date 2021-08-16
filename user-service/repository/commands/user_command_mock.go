package commands

import (
	"database/sql"
	"github.com/thel5coder/ecommerce/user-service/domain/models"
)

type UserCommandMock struct {
	db    *sql.DB
	model *models.User
}

func (c UserCommandMock) Add() (res string, err error) {
	statement := `INSERT INTO users (email,first_name,last_name,password,role_id,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	err = c.db.QueryRow(statement, c.model.Email(), c.model.FirstName(), c.model.LastName(),c.model.Password(),c.model.RoleId(), c.model.CreatedAt(),
		c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c UserCommandMock) Edit() (res string, err error) {
	setStatement := `first_name=$1,last_name=$2,email=$3,role_id=$4,updated_at=$5`
	editParams := []interface{}{c.model.FirstName(), c.model.LastName(), c.model.Email(), c.model.RoleId(), c.model.UpdatedAt(),
		c.model.Id()}
	if c.model.Password() != "" {
		setStatement += `,password=$7`
		editParams = append(editParams, c.model.Password())
	}

	statement := `UPDATE users SET ` + setStatement + ` WHERE id=$6 RETURNING id`
	err = c.db.QueryRow(statement, editParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c UserCommandMock) Delete() (res string, err error) {
	statement := `UPDATE users SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`
	err = c.db.QueryRow(statement, c.model.UpdatedAt(), c.model.DeletedAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
