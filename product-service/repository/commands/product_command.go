package commands

import (
	"github.com/ecommerce-service/product-service/domain/commands"
	"github.com/ecommerce-service/product-service/domain/models"
	"github.com/thel5coder/pkg/postgresql"
)

type ProductCommand struct {
	db    postgresql.IConnection
	model *models.Product
}

func NewProductCommand(db postgresql.IConnection, model *models.Product) commands.IProductCommand {
	return &ProductCommand{
		db:    db,
		model: model,
	}
}

func (c ProductCommand) Add() (res string, err error) {
	statement := `INSERT INTO products(category_id,name,sku,price,discount,stock,main_image_key,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`

	err = c.db.GetTx().QueryRow(statement, c.model.CategoryId(), c.model.Name(), c.model.Sku(), c.model.Price(), c.model.Discount().Float64, c.model.Stock(), c.model.MainImageKey().String,
		c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c ProductCommand) Edit() (err error) {
	statement := `UPDATE products SET category_id=$1,name=$2,sku=$3,price=$4,discount=$5,stock=$6,main_image_key=$7,updated_at=$8 WHERE id=$9 RETURNING id`

	_, err = c.db.GetTx().Exec(statement, c.model.CategoryId(), c.model.Name(), c.model.Sku(), c.model.Price(), c.model.Discount().Float64, c.model.Stock(), c.model.MainImageKey().String,
		c.model.UpdatedAt(), c.model.Id())
	if err != nil {
		return err
	}

	return nil
}

func (c ProductCommand) Delete() (err error) {
	statement := `UPDATE products SET updated_at=$1,deleted_at=$2 WHERE id=$3`

	_, err = c.db.GetTx().Exec(statement, c.model.UpdatedAt(), c.model.DeletedAt().Time, c.model.Id())
	if err != nil {
		return err
	}

	return nil
}
