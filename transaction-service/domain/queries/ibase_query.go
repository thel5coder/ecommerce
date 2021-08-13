package queries

type IBaseQuery interface {
	Browse(search, orderBy, sort string, limit, offset int) (interface{}, error)

	ReadBy(column, operator string, value interface{}) (interface{}, error)

	Count(search string) (res int, err error)

	CountBy(column, operator, id string, value interface{}) (res int, err error)
}
