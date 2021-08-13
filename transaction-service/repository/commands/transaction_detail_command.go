package commands

import (
	"github.com/ecommerce-service/transaction-service/domain/commands"
	"github.com/ecommerce-service/transaction-service/domain/models"
	"github.com/thel5coder/pkg/postgresql"
)

type TransactionDetailCommand struct {
	db    postgresql.IConnection
	model *models.TransactionDetails
}

func NewTransactionDetailCommand(db postgresql.IConnection, model *models.TransactionDetails) commands.ITransactionDetailCommand {
	return &TransactionDetailCommand{
		db:    db,
		model: model,
	}
}

func (c TransactionDetailCommand) Add() (err error) {
	statement := `INSERT INTO transaction_details(transaction_id,name,sku,category,price,discount,quantity,sub_total,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	_, err = c.db.GetTx().Exec(statement, c.model.TransactionId(), c.model.Name(), c.model.Sku(), c.model.Category(), c.model.Price(), c.model.Discount(),
		c.model.Quantity(), c.model.SubTotal(), c.model.CreatedAt(), c.model.UpdatedAt())
	if err != nil {
		return err
	}

	return nil
}
