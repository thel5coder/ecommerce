package models

import (
	"database/sql"
)

type Role struct {
	id        int64
	name      string
}

func NewRoleModel() *Role {
	return &Role{}
}

func (model *Role) Id() int64 {
	return model.id
}

func (model *Role) SetId(id int64) *Role {
	model.id = id

	return model
}

func (model *Role) Name() string {
	return model.name
}

func (model *Role) SetName(name string) *Role {
	model.name = name

	return model
}

const (
	RoleSelectStatement = `SELECT id,name FROM roles`
	RoleWhereStatement  = `WHERE LOWER(name) LIKE $1`
	RoleOrderStatement  = `ORDER BY id ASC`
)

func (model *Role) ScanRows(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&model.id, &model.name)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Role) ScanRow(row *sql.Row) (interface{}, error) {
	err := row.Scan(&model.id, &model.name)
	if err != nil {
		return model, err
	}

	return model, nil
}
