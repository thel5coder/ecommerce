package commands

import (
	"database/sql"
	"github.com/ecommerce-service/transaction-service/domain/models"
)

type TransactionCommandMock struct {
	db    *sql.DB
	model *models.Transactions

}

func (c TransactionCommandMock) Add() (res string, err error) {
	statement := `INSERT INTO transactions ` +
		`(user_id,transaction_number,status,total,discount,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	tx, _ := c.db.Begin()

	err = tx.QueryRow(statement, c.model.UserId(), c.model.TransactionNumber(), c.model.Status(), c.model.Total(), c.model.Discount(),
		c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		tx.Rollback()
		return res, err
	}
	tx.Commit()

	return res, nil
}

func (c TransactionCommandMock) EditPaymentReceived() (err error) {
	statement := `UPDATE transactions set updated_at=$1,paid_at=$2 WHERE id=$3`
	tx, _ := c.db.Begin()
	_, err = tx.Exec(statement, c.model.UpdatedAt(), c.model.PaidAt().Time, c.model.Id())
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (c TransactionCommandMock) EditCancelPayment() (res string, err error) {
	statement := `UPDATE transactions set updated_at=$1,canceled_at=$2 WHERE id=$3 RETURNING id`
	tx, _ := c.db.Begin()

	err = tx.QueryRow(statement, c.model.UpdatedAt(), c.model.CanceledAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		tx.Rollback()
		return res, err
	}
	tx.Commit()

	return res, nil
}
