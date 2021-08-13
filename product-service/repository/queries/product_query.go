package queries

import (
	"github.com/ecommerce-service/product-service/domain/models"
	"github.com/ecommerce-service/product-service/domain/queries"
	"github.com/thel5coder/pkg/postgresql"
	"strings"
)

type ProductQuery struct {
	db postgresql.IConnection
}

func NewProductQuery(db postgresql.IConnection) queries.IProductQuery {
	return ProductQuery{db: db}
}

func (q ProductQuery) Browse(search, orderBy, sort, category string, limit, offset int) (interface{}, error) {
	var res []*models.Product
	queryParams := []interface{}{"%" + strings.ToLower(search) + "%", limit, offset}
	optionalWhereStatement := ``
	if category != "" {
		optionalWhereStatement += `AND c.slug=$4`
		queryParams = append(queryParams, category)
	}

	statement := models.ProductSelectStatement + ` ` + models.ProductJoinSelectStatement + ` ` + models.ProductDefaultWhereStatement + ` AND (LOWER(p.name) like $1 OR LOWER(p.sku) ` +
		`LIKE $1 OR cast(p.price as varchar) LIKE $1) ` + optionalWhereStatement + ` ` + models.ProductGroupByStatement + ` ORDER BY ` + orderBy + ` ` + sort + ` LIMIT $2 OFFSET $3`
	rows, err := q.db.GetDbInstance().Query(statement, queryParams...)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := models.NewProductModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.Product))
	}

	return res, nil
}

func (q ProductQuery) ReadBy(column, operator string, value interface{}) (interface{}, error) {
	statement := models.ProductDetailSelectStatement + ` ` + models.ProductJoinSelectStatement + ` ` + models.ProductJoinDetailSelectStatement + ` ` +
		models.ProductDefaultWhereStatement + ` AND ` + column + `` + operator + `$1 `+models.ProductGroupByStatement

	row := q.db.GetDbInstance().QueryRow(statement, value)
	res, err := models.NewProductModel().ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q ProductQuery) Count(search, category string) (res int, err error) {
	queryParams := []interface{}{"%" + strings.ToLower(search) + "%"}
	optionalWhereStatement := ``
	if category != "" {
		optionalWhereStatement += `AND c.slug=$2`
		queryParams = append(queryParams, category)
	}
	statement := models.ProductSelectCountStatement + ` ` + models.ProductJoinSelectStatement + ` ` + models.ProductDefaultWhereStatement + ` AND (LOWER(p.name) like $1 OR LOWER(p.sku) ` +
		`LIKE $1 OR cast(p.price as varchar) LIKE $1) ` + optionalWhereStatement

	err = q.db.GetDbInstance().QueryRow(statement, queryParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (q ProductQuery) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	whereStatement := models.ProductDefaultWhereStatement + ` AND ` + column + `` + operator + `$1`
	whereParams := []interface{}{value}
	if id != "" {
		whereStatement += ` AND p.id<>$2`
		whereParams = append(whereParams, id)
	}
	statement := models.ProductSelectCountStatement + ` ` + models.ProductJoinSelectStatement + ` ` + whereStatement

	err = q.db.GetDbInstance().QueryRow(statement, whereParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
