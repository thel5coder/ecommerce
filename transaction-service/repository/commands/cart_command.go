package commands

import (
	"github.com/ecommerce-service/transaction-service/domain/commands"
	"github.com/ecommerce-service/transaction-service/domain/models"
	"github.com/thel5coder/pkg/postgresql"
)

type CartCommand struct {
	db    postgresql.IConnection
	model *models.Carts
}

func NewCartCommand(db postgresql.IConnection, model *models.Carts) commands.ICartCommand {
	return &CartCommand{
		db:    db,
		model: model,
	}
}

func (c CartCommand) Add() (res string, err error) {
	statement := `INSERT INTO carts(user_id,product_id,name,sku,category,price,quantity,sub_total,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.UserId(), c.model.ProductId(), c.model.Name(), c.model.Sku(), c.model.Category(), c.model.Price(),
		c.model.Quantity(), c.model.SubTotal(), c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CartCommand) Edit() (res string, err error) {
	statement := `UPDATE carts SET quantity=$1,sub_total=$2,updated_at=$3 WHERE id=$4 RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.Quantity(), c.model.SubTotal(), c.model.UpdatedAt(), c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}


func (c CartCommand) Delete() (res string, err error) {
	statement := `UPDATE carts SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.UpdatedAt(), c.model.DeletedAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CartCommand) DeleteAllByUserID() (err error) {
	statement := `UPDATE carts SET updated_at=$1,deleted_at=$2 WHERE user_id=$3 RETURNING id`

	_, err = c.db.GetTx().Exec(statement, c.model.UpdatedAt(), c.model.DeletedAt().Time, c.model.UserId())
	if err != nil {
		return err
	}

	return nil
}
