package queries

import (
	"database/sql"
	"fmt"
	"github.com/ecommerce-service/transaction-service/domain/models"
	"strings"
)

type TransactionQueryMock struct{
	db *sql.DB
}

func (q TransactionQueryMock) Browse(search, orderBy, sort, status string, limit, offset int) (interface{}, error) {
	var res []*models.Transactions
	whereStatement := models.TransactionDefaultWhereStatement + ` AND t.transaction_number LIKE $1`
	params := []interface{}{"%" + strings.ToLower(search) + "%", limit, offset}
	if status != "" {
		whereStatement += ` AND t.status=$4`
		params = append(params, status)
	}
	statement := models.TransactionSelectListStatement + ` FROM transactions t ` + whereStatement + ` ORDER BY ` + orderBy + ` ` + sort +
		` LIMIT $2 OFFSET $3`

	rows, err := q.db.Query(statement, params...)
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

func (q TransactionQueryMock) ReadBy(column, operator string, value interface{}) (interface{}, error) {
	statement := models.TransactionSelectListStatement + ` ` + models.TransactionSelectDetailStatement + ` FROM transactions t ` +
		models.TransactionDetailJoinStatement + ` ` + models.TransactionDefaultWhereStatement + ` AND ` + column + `` + operator + `$1 ` + models.TransactionGroupByStatement
	row := q.db.QueryRow(statement, value)
	res, err := models.NewTransactionModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q TransactionQueryMock) Count(search, userId, status string) (res int, err error) {
	whereStatement := models.TransactionDefaultWhereStatement + ` AND t.transaction_number LIKE $1`
	if status != "" {
		whereStatement += ` AND t.status='` + status + `'`
	}
	if userId != "" {
		whereStatement += ` AND t.user_id='` + userId + `'`
	}
	statement := models.TransactionSelectCountStatement + ` ` + whereStatement
	fmt.Println(statement)

	err = q.db.QueryRow(statement, "%" + strings.ToLower(search) + "%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q TransactionQueryMock) CountBy(column, operator string, value interface{}) (res int, err error) {
	statement := models.TransactionSelectCountStatement + ` WHERE ` + column + `` + operator + `$1`
	fmt.Println(statement)

	err = q.db.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q TransactionQueryMock) CountAll() (res int, err error) {
	statement := models.TransactionSelectCountStatement

	err = q.db.QueryRow(statement).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q TransactionQueryMock) BrowseByUserId(search, orderBy, sort, userId, status string, limit, offset int) (interface{}, error) {
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

	rows, err := q.db.Query(statement, params...)
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

