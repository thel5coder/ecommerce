package models

import (
	"database/sql"
	"time"
)

type Product struct {
	id            string
	categoryId    string
	name          string
	sku           string
	price         float64
	discount      sql.NullFloat64
	stock         int
	mainImageKey  sql.NullString
	createdAt     time.Time
	updatedAt     time.Time
	deletedAt     sql.NullTime
	productImages sql.NullString

	Category *Category
}

func NewProductModel() *Product {
	return &Product{}
}

func (model *Product) Id() string {
	return model.id
}

func (model *Product) SetId(id string) *Product {
	model.id = id

	return model
}

func (model *Product) CategoryId() string {
	return model.categoryId
}

func (model *Product) SetCategoryId(categoryId string) *Product {
	model.categoryId = categoryId

	return model
}

func (model *Product) Name() string {
	return model.name
}

func (model *Product) SetName(name string) *Product {
	model.name = name

	return model
}

func (model *Product) Sku() string {
	return model.sku
}

func (model *Product) SetSku(sku string) *Product {
	model.sku = sku

	return model
}

func (model *Product) Price() float64 {
	return model.price
}

func (model *Product) SetPrice(price float64) *Product {
	model.price = price

	return model
}

func (model *Product) Discount() sql.NullFloat64 {
	return model.discount
}

func (model *Product) SetDiscount(discount float64) *Product {
	model.discount.Float64 = discount

	return model
}

func (model *Product) Stock() int {
	return model.stock
}

func (model *Product) SetStock(stock int) *Product {
	model.stock = stock

	return model
}

func (model *Product) MainImageKey() sql.NullString {
	return model.mainImageKey
}

func (model *Product) SetMainImageKey(mainImageKey string) *Product {
	model.mainImageKey.String = mainImageKey

	return model
}

func (model *Product) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Product) SetCreatedAt(createdAt time.Time) *Product {
	model.createdAt = createdAt

	return model
}

func (model *Product) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Product) SetUpdatedAt(updatedAt time.Time) *Product {
	model.updatedAt = updatedAt

	return model
}

func (model *Product) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Product) SetDeletedAt(deletedAt time.Time) *Product {
	model.deletedAt.Time = deletedAt

	return model
}

func (model *Product) ProductImages() sql.NullString {
	return model.productImages
}

func (model *Product) SetProductImages(productImages sql.NullString) {
	model.productImages = productImages
}

const (
	ProductSelectStatement = `SELECT p.id,p.category_id,p.name,p.sku,p.price,p.discount,p.stock,p.main_image_key,p.created_at,p.updated_at,` +
		`c.name FROM products p`
	ProductDetailSelectStatement = `SELECT p.id,p.category_id,p.name,p.sku,p.price,p.discount,p.stock,p.main_image_key,p.created_at,p.updated_at,` +
		`c.name,ARRAY_TO_STRING(ARRAY_AGG(pi.image_key),',') FROM products p`
	ProductSelectCountStatement      = `SELECT COUNT(DISTINCT p.id) FROM products p`
	ProductJoinSelectStatement       = `INNER JOIN categories c ON c.id=p.category_id AND c.deleted_at IS NULL `
	ProductJoinDetailSelectStatement = `LEFT JOIN product_images pi ON pi.product_id=p.id`
	ProductDefaultWhereStatement     = `WHERE p.deleted_at IS NULL`
	ProductGroupByStatement          = `GROUP BY p.id,c.id`
)

func (model *Product) ScanRows(rows *sql.Rows) (interface{}, error) {
	model.Category = NewCategoryModel()
	err := rows.Scan(&model.id, &model.categoryId, &model.name, &model.sku, &model.price, &model.discount, &model.stock, &model.mainImageKey, &model.createdAt, &model.updatedAt, &model.Category.name)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Product) ScanRow(row *sql.Row) (interface{}, error) {
	model.Category = NewCategoryModel()
	err := row.Scan(&model.id, &model.categoryId, &model.name, &model.sku, &model.price, &model.discount, &model.stock, &model.mainImageKey, &model.createdAt, &model.updatedAt, &model.Category.name,
		&model.productImages)
	if err != nil {
		return model, err
	}

	return model, nil
}
