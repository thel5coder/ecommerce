package commands

type ITransactionDetailCommand interface {
	Add() (err error)
}
