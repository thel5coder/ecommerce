package queries

type IProductQuery interface {
	Browse(search, orderBy, sort,category string, limit, offset int) (interface{}, error)

	ReadBy(column, operator string, value interface{}) (interface{}, error)

	Count(search,category string) (res int, err error)

	CountBy(column, operator, id string, value interface{}) (res int, err error)
}
