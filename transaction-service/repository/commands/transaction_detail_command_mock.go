package commands

import (
	"database/sql"
	"github.com/ecommerce-service/transaction-service/domain/models"
)

type TransactionDetailCommandMock struct {
	db    *sql.DB
	model *models.TransactionDetails
}

func (c TransactionDetailCommandMock) Add() (err error) {
	statement := `INSERT INTO transaction_details(transaction_id,name,sku,category,price,discount,quantity,sub_total,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	tx, _ := c.db.Begin()

	_, err = tx.Exec(statement, c.model.TransactionId(), c.model.Name(), c.model.Sku(), c.model.Category(), c.model.Price(), c.model.Discount(),
		c.model.Quantity(), c.model.SubTotal(), c.model.CreatedAt(), c.model.UpdatedAt())
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
