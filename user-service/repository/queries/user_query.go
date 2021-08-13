package queries

import (
	"database/sql"
	"github.com/ecommerce/user-service/domain/models"
	"github.com/ecommerce/user-service/domain/queries"
	"strings"
)

type UserQuery struct {
	db *sql.DB
}

func NewUserQuery(db *sql.DB) queries.IUserQuery {
	return &UserQuery{db: db}
}

func (q UserQuery) Browse(search, orderBy, sort string, limit, offset int) (interface{}, error) {
	var res []*models.User

	statement := models.UserSelectStatement + ` ` + models.UserJoinStatement + ` ` + models.UserDefaultWhereStatement +
		` AND (LOWER(u.first_name) LIKE $1 OR LOWER(last_name) LIKE $1 OR LOWER(email) LIKE $1) ORDER BY ` + orderBy + ` ` + sort + ` LIMIT $2 OFFSET $3`
	rows, err := q.db.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewUserModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.User))
	}

	return res, nil
}

func (q UserQuery) ReadBy(column, operator string, value interface{}) (interface{}, error) {
	statement := models.UserSelectStatement + ` ` + models.UserJoinStatement + ` ` + models.UserDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`
	row := q.db.QueryRow(statement, value)
	res, err := models.NewUserModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, err
}

func (q UserQuery) Count(search string) (res int, err error) {
	statement := models.UserSelectCountStatement + `` + models.UserJoinStatement + ` ` + models.UserDefaultWhereStatement +
		` AND (LOWER(u.first_name) LIKE $1 OR LOWER(last_name) LIKE $1 OR LOWER(email) LIKE $1)`
	err = q.db.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q UserQuery) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	whereStatement := models.UserDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`
	whereParams := []interface{}{value}
	if id != "" {
		whereStatement += ` AND u.id<>$2`
		whereParams = append(whereParams, id)
	}

	statement := models.UserSelectCountStatement + `` + models.UserJoinStatement + ` ` + whereStatement
	err = q.db.QueryRow(statement, whereParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
