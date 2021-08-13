package commands

import (
	"github.com/ecommerce-service/transaction-service/domain/commands"
	"github.com/ecommerce-service/transaction-service/domain/models"
	"github.com/thel5coder/pkg/postgresql"
)

type TransactionCommand struct {
	db    postgresql.IConnection
	model *models.Transactions
}

func NewTransactionCommand(db postgresql.IConnection, model *models.Transactions) commands.ITransactionCommand {
	return &TransactionCommand{
		db:    db,
		model: model,
	}
}

func (c TransactionCommand) Add() (res string, err error) {
	statement := `INSERT INTO transactions ` +
		`(user_id,transaction_number,status,total,discount,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`

	err = c.db.GetTx().QueryRow(statement, c.model.UserId(), c.model.TransactionNumber(), c.model.Status(), c.model.Total(), c.model.Discount(),
		c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c TransactionCommand) EditPaymentReceived() (err error) {
	statement := `UPDATE transactions set updated_at=$1,paid_at=$2 WHERE id=$3`

	_, err = c.db.GetTx().Exec(statement, c.model.UpdatedAt(), c.model.PaidAt().Time, c.model.Id())
	if err != nil {
		return err
	}

	return nil
}

func (c TransactionCommand) EditCancelPayment() (res string, err error) {
	statement := `UPDATE transactions set status=$1,updated_at=$2,canceled_at=$3 WHERE id=$4 RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.Status(),c.model.UpdatedAt(), c.model.CanceledAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
