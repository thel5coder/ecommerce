package queries

type ICategoryQuery interface {
	IBaseQuery

	BrowseAll(search string) (interface{},error)
}
