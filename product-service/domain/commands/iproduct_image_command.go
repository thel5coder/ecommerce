package commands

import "database/sql"

type IProductImageCommand interface {
	Add() (err error)

	Delete() (res sql.Result, err error)
}
