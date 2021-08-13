package commands

type IBaseCommand interface {
	Add() (res string, err error)

	Edit() (res string, err error)

	Delete() (res string, err error)
}
