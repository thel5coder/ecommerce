package commands

import (
	"database/sql"
	"fmt"
	"github.com/thel5coder/ecommerce/product-service/domain/commands"
	"github.com/thel5coder/ecommerce/product-service/domain/models"
	"github.com/thel5coder/pkg/postgresql"
)

type ProductImageCommand struct {
	db    postgresql.IConnection
	model *models.ProductImage
}

// NewProductImageCommand initialization for new command product image
func NewProductImageCommand(db postgresql.IConnection, model *models.ProductImage) commands.IProductImageCommand {
	return &ProductImageCommand{
		db:    db,
		model: model,
	}
}

//Add query to insert into product_images table
func (c ProductImageCommand) Add() (err error) {
	statement := `INSERT INTO product_images (product_id,image_key) values($1,$2)`
	fmt.Println(c.model.ImageKey())

	_, err = c.db.GetTx().Exec(statement, c.model.ProductId(), c.model.ImageKey())
	if err != nil {
		return err
	}

	return nil
}

//Delete query to delete data in product_images table
func (c ProductImageCommand) Delete() (res sql.Result, err error) {
	statement := `DELETE FROM product_images WHERE product_id=$1`

	res, err = c.db.GetTx().Exec(statement, c.model.ProductId())
	if err != nil {
		return res, err
	}

	return res, nil
}
