package queries

import (
	"github.com/thel5coder/ecommerce/product-service/domain/models"
	"github.com/thel5coder/ecommerce/product-service/domain/queries"
	"github.com/thel5coder/pkg/postgresql"
	"strings"
)

type CategoryQuery struct {
	db postgresql.IConnection
}

// NewCategoryQuery initialization for new command product
func NewCategoryQuery(db postgresql.IConnection) queries.ICategoryQuery {
	return &CategoryQuery{db: db}
}

//Browse query to select data from categories table with order,search,and limit offset
func (q CategoryQuery) Browse(search, orderBy, sort string, limit, offset int) (interface{}, error) {
	var res []*models.Category
	statement := models.CategorySelectStatement + ` ` + models.CategoryDefaultWhereStatement + ` AND LOWER(name) LIKE $1 ORDER BY ` + orderBy + ` ` + sort + ` LIMIT $2 OFFSET $3`

	rows, err := q.db.GetDbInstance().Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewCategoryModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.Category))
	}

	return res, nil
}

//BrowseAll query to select all data from categories table
func (q CategoryQuery) BrowseAll(search string) (interface{}, error) {
	var res []*models.Category
	statement := models.CategorySelectStatement + ` ` + models.CategoryDefaultWhereStatement + ` AND LOWER(name) LIKE $1`

	rows, err := q.db.GetDbInstance().Query(statement, "%"+strings.ToLower(search)+"%")
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewCategoryModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.Category))
	}

	return res, nil
}

//ReadBy query to select from categories table and filter by specific column by parameters
func (q CategoryQuery) ReadBy(column, operator string, value interface{}) (interface{}, error) {
	statement := models.CategorySelectStatement + ` ` + models.CategoryDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`

	row := q.db.GetDbInstance().QueryRow(statement, value)
	res, err := models.NewCategoryModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

//Count query to select count all data from categories table with filter
func (q CategoryQuery) Count(search string) (res int, err error) {
	statement := models.CategorySelectCountStatement+` `+models.CategoryDefaultWhereStatement+` AND LOWER(name) LIKE $1`

	err = q.db.GetDbInstance().QueryRow(statement,"%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res,err
	}

	return res,nil
}

//CountBy query to select count from categories table by specific column by parameters
func (q CategoryQuery) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	whereStatement := models.CategoryDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`
	whereParams := []interface{}{value}
	if id != "" {
		whereStatement += ` AND id<>$2`
		whereParams = append(whereParams, id)
	}
	statement := models.CategorySelectCountStatement + ` ` + whereStatement

	err = q.db.GetDbInstance().QueryRow(statement, whereParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
