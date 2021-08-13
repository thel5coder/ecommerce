package commands

type ICartCommand interface {
	IBaseCommand

	DeleteAllByUserID() (err error)
}
