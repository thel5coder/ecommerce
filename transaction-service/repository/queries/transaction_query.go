package queries

import (
	"fmt"
	"github.com/ecommerce-service/transaction-service/domain/models"
	"github.com/ecommerce-service/transaction-service/domain/queries"
	"github.com/thel5coder/pkg/postgresql"
	"strings"
)

type TransactionQuery struct {
	db postgresql.IConnection
}

func NewTransactionQuery(db postgresql.IConnection) queries.ITransactionQuery {
	return &TransactionQuery{db: db}
}

func (q TransactionQuery) Browse(search, orderBy, sort, status string, limit, offset int) (interface{}, error) {
	var res []*models.Transactions
	whereStatement := models.TransactionDefaultWhereStatement + ` AND t.transaction_number LIKE $1`
	params := []interface{}{"%" + strings.ToLower(search) + "%", limit, offset}
	if status != "" {
		whereStatement += ` AND t.status=$4`
		params = append(params, status)
	}
	statement := models.TransactionSelectListStatement + ` FROM transactions t ` + whereStatement + ` ORDER BY ` + orderBy + ` ` + sort +
		` LIMIT $2 OFFSET $3`

	rows, err := q.db.GetDbInstance().Query(statement, params...)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewTransactionModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.Transactions))
	}

	return res, nil
}

func (q TransactionQuery) ReadBy(column, operator string, value interface{}) (interface{}, error) {
	statement := models.TransactionSelectListStatement + ` ` + models.TransactionSelectDetailStatement + ` FROM transactions t ` +
		models.TransactionDetailJoinStatement + ` ` + models.TransactionDefaultWhereStatement + ` AND ` + column + `` + operator + `$1 ` + models.TransactionGroupByStatement

	row := q.db.GetDbInstance().QueryRow(statement, value)
	res, err := models.NewTransactionModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q TransactionQuery) Count(search, userId, status string) (res int, err error) {
	whereStatement := models.TransactionDefaultWhereStatement + ` AND t.transaction_number LIKE $1`
	if status != "" {
		whereStatement += ` AND t.status='` + status + `'`
	}
	if userId != "" {
		whereStatement += ` AND t.user_id='` + userId + `'`
	}
	statement := models.TransactionSelectCountStatement + ` ` + whereStatement

	err = q.db.GetDbInstance().QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q TransactionQuery) CountAll() (res int, err error) {
	statement := models.TransactionSelectCountStatement

	err = q.db.GetDbInstance().QueryRow(statement).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q TransactionQuery) BrowseByUserId(search, orderBy, sort, userId, status string, limit, offset int) (interface{}, error) {
	var res []*models.Transactions
	whereStatement := models.TransactionDefaultWhereStatement + ` AND t.transaction_number LIKE $1 AND t.user_id=$2`
	params := []interface{}{"%" + strings.ToLower(search) + "%", userId, limit, offset}
	if status != "" {
		whereStatement += ` AND t.status=$5`
		params = append(params, status)
	}
	statement := models.TransactionSelectListStatement + ` FROM transactions t ` + whereStatement + ` ORDER BY ` + orderBy + ` ` + sort +
		` LIMIT $3 OFFSET $4`
fmt.Println(statement)
	rows, err := q.db.GetDbInstance().Query(statement, params...)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewTransactionModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.Transactions))
	}

	return res, nil
}
