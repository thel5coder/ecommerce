package commands

type ITransactionCommand interface {
	Add() (res string, err error)

	EditPaymentReceived() (err error)

	EditCancelPayment() (res string,err error)
}
