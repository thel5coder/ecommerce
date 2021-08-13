package queries

type ICartQuery interface {
	BrowseByUser(search, orderBy, sort, userId string, limit, offset int) (interface{}, error)

	BrowseAllByUser(userId string) (interface{}, error)

	ReadBy(column, operator, userId string, value interface{}) (interface{}, error)

	Count(search, userId string) (res int, err error)

	CountBy(column, operator, userId string, value interface{}) (res int, err error)
}
