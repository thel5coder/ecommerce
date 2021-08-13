package queries

import "github.com/ecommerce-service/user-service/domain/models"

type IRoleQuery interface {
	BrowseAll(search string) (res []*models.Role, err error)
}
