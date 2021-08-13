package models

import "database/sql"

type IModel interface {
	ScanRows(rows *sql.Rows) (interface{}, error)

	ScanRow(row *sql.Row) (interface{}, error)
}
