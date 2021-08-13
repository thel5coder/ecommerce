package models

import (
	"database/sql"
	"time"
)

type Category struct {
	id        string
	name      string
	slug      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func NewCategoryModel() *Category {
	return &Category{}
}

func (model *Category) Id() string {
	return model.id
}

func (model *Category) SetId(id string) *Category {
	model.id = id

	return model
}

func (model *Category) Name() string {
	return model.name
}

func (model *Category) SetName(name string) *Category {
	model.name = name

	return model
}

func (model *Category) Slug() string {
	return model.slug
}

func (model *Category) SetSlug(slug string) *Category {
	model.slug = slug

	return model
}

func (model *Category) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Category) SetCreatedAt(createdAt time.Time) *Category {
	model.createdAt = createdAt

	return model
}

func (model *Category) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Category) SetUpdatedAt(updatedAt time.Time) *Category {
	model.updatedAt = updatedAt

	return model
}

func (model *Category) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Category) SetDeletedAt(deletedAt time.Time) *Category {
	model.deletedAt.Time = deletedAt

	return model
}

const (
	CategorySelectStatement       = `SELECT id,name,slug FROM categories`
	CategorySelectCountStatement  = `SELECT count(id) FROM categories`
	CategoryDefaultWhereStatement = `WHERE deleted_at IS NULL`
)

func (model *Category) ScanRows(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&model.id, &model.name, &model.slug)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Category) ScanRow(row *sql.Row) (interface{}, error) {
	err := row.Scan(&model.id, &model.name, &model.slug)
	if err != nil {
		return model, err
	}

	return model, nil
}
