package queries

import "github.com/thel5coder/ecommerce/user-service/domain/models"

type IRoleQuery interface {
	BrowseAll(search string) (res []*models.Role, err error)
}
