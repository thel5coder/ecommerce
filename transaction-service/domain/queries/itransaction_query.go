package queries

type ITransactionQuery interface {
	Browse(search, orderBy, sort, status string, limit, offset int) (interface{}, error)

	ReadBy(column, operator string, value interface{}) (interface{}, error)

	Count(search, userID, status string) (res int, err error)

	CountAll() (res int, err error)

	BrowseByUserId(search, orderBy, sort, userId, status string, limit, offset int) (interface{}, error)
}
