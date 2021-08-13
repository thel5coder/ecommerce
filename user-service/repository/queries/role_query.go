package queries

import (
	"database/sql"
	"github.com/ecommerce/user-service/domain/models"
	"github.com/ecommerce/user-service/domain/queries"
	"strings"
)

type RoleQuery struct {
	db *sql.DB
}

func NewRoleQuery(DB *sql.DB) queries.IRoleQuery {
	return &RoleQuery{db: DB}
}

// BrowseAll this function have function to query all roles data with search to name field
func (q RoleQuery) BrowseAll(search string) (res []*models.Role, err error) {
	statement := models.RoleSelectStatement + ` ` + models.RoleWhereStatement + ` ` + models.RoleOrderStatement

	rows, err := q.db.Query(statement, "%"+strings.ToLower(search)+"%")
	if err != nil {
		return res, err
	}
	for rows.Next() {
		temp, err := models.NewRoleModel().ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp.(*models.Role))
	}

	return res, nil
}
