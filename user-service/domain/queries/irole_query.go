package queries

import "github.com/ecommerce/user-service/domain/models"

type IRoleQuery interface {
	BrowseAll(search string) (res []*models.Role, err error)
}
