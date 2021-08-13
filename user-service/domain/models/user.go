package models

import (
	"database/sql"
	"time"
)

type User struct {
	id        string
	email     string
	firstName string
	lastName  string
	password  string
	roleId    int64
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime

	Role *Role
}

func (model *User) Id() string {
	return model.id
}

func (model *User) SetId(id string) *User {
	model.id = id

	return model
}

func (model *User) FirstName() string {
	return model.firstName
}

func (model *User) SetFirstName(firstName string) *User {
	model.firstName = firstName

	return model
}

func (model *User) LastName() string {
	return model.lastName
}

func (model *User) SetLastName(lastName string) *User {
	model.lastName = lastName

	return model
}

func (model *User) Email() string {
	return model.email
}

func (model *User) SetEmail(email string) *User {
	model.email = email

	return model
}

func (model *User) Password() string {
	return model.password
}

func (model *User) SetPassword(password string) *User {
	model.password = password

	return model
}

func (model *User) RoleId() int64 {
	return model.roleId
}

func (model *User) SetRoleId(roleId int64) *User {
	model.roleId = roleId

	return model
}

func (model *User) CreatedAt() time.Time {
	return model.createdAt
}

func (model *User) SetCreatedAt(createdAt time.Time) *User {
	model.createdAt = createdAt

	return model
}

func (model *User) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *User) SetUpdatedAt(updatedAt time.Time) *User {
	model.updatedAt = updatedAt

	return model
}

func (model *User) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *User) SetDeletedAt(deletedAt time.Time) *User {
	model.deletedAt.Time = deletedAt

	return model
}

func NewUserModel() *User {
	return &User{}
}

const (
	UserSelectStatement       = `SELECT u.id,u.email,u.first_name,u.last_name,u.password,u.created_at,u.updated_at,r.id,r.name FROM users u`
	UserSelectCountStatement  = `SELECT COUNT(u.id) FROM users u `
	UserJoinStatement         = `INNER JOIN roles r ON r.id = u.role_id`
	UserDefaultWhereStatement = `WHERE u.deleted_at IS NULL`
)

func (model *User) ScanRows(rows *sql.Rows) (interface{}, error) {
	model.Role = NewRoleModel()
	err := rows.Scan(&model.id, &model.email, &model.firstName, &model.lastName, &model.password, &model.createdAt,
		&model.updatedAt, &model.Role.id, &model.Role.name)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *User) ScanRow(row *sql.Row) (interface{}, error) {
	model.Role = NewRoleModel()
	err := row.Scan(&model.id, &model.email, &model.firstName, &model.lastName, &model.password, &model.createdAt,
		&model.updatedAt, &model.Role.id, &model.Role.name)
	if err != nil {
		return model, err
	}

	return model, nil
}
