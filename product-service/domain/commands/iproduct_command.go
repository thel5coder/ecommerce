package commands

type IProductCommand interface {
	Add() (res string, err error)

	Edit() (err error)

	Delete() (err error)
}
