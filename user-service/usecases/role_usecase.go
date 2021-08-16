package usecases

import (
	"github.com/thel5coder/ecommerce/user-service/domain/usecase"
	"github.com/thel5coder/ecommerce/user-service/domain/view_models"
	"github.com/thel5coder/ecommerce/user-service/repository/queries"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/logruslogger"
)

type RoleUseCase struct {
	*UseCaseContract
}

func NewRoleUseCase(useCaseContract *UseCaseContract) usecase.IRoleUseCase {
	return &RoleUseCase{UseCaseContract: useCaseContract}
}

func (uc RoleUseCase) BrowseAll(search string) (res []view_models.RoleVm, err error) {
	q := queries.NewRoleQuery(uc.Config.DB.GetDbInstance())

	roles, err := q.BrowseAll(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-role-browse")
		return res, err
	}
	for _, role := range roles {
		res = append(res, view_models.NewRoleVm(role))
	}

	return res, nil
}
