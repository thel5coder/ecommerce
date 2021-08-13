package usecases

import (
	"errors"
	"github.com/ecommerce/user-service/domain/models"
	"github.com/ecommerce/user-service/domain/requests"
	"github.com/ecommerce/user-service/domain/usecase"
	"github.com/ecommerce/user-service/domain/view_models"
	"github.com/ecommerce/user-service/repository/commands"
	"github.com/ecommerce/user-service/repository/queries"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/hashing"
	"github.com/thel5coder/pkg/logruslogger"
	"github.com/thel5coder/pkg/messages"
	"time"
)

type UserUseCase struct {
	*UseCaseContract
}

func NewUserUseCase(useCaseContract *UseCaseContract) usecase.IUserUseCase {
	return &UserUseCase{UseCaseContract: useCaseContract}
}

func (uc UserUseCase) GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.UserVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewUserQuery(uc.Config.DB.GetDbInstance())
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	users, err := q.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-browse")
		return res, pagination, err
	}
	for _, user := range users.([]*models.User) {
		res = append(res, view_models.NewUserVm(user))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc UserUseCase) GetByID(id string) (res view_models.UserVm, err error) {
	q := queries.NewUserQuery(uc.Config.DB.GetDbInstance())

	user, err := q.ReadBy("u.id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-readByID")
		return res, err
	}
	res = view_models.NewUserVm(user.(*models.User))

	return res, nil
}

func (uc UserUseCase) Edit(req *requests.UserEditRequest, id string) (res string, err error) {
	now := time.Now().UTC()
	var password string

	count, err := uc.CountBy("u.email", "=", id, req.Email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-countByEmail")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.EmailAlreadyExist, functioncaller.PrintFuncName(), "uc-user-countByEmail")
		return res, err
	}

	model := models.NewUserModel().SetFirstName(req.FirstName).SetLastName(req.LastName).SetEmail(req.Email).SetUpdatedAt(now).SetId(id).SetRoleId(req.RoleId)
	if req.Password != "" {
		password, _ = hashing.HashAndSalt(req.Password)
		model.SetPassword(password)
	}
	cmd := commands.NewUserCommand(uc.Config.DB, model)
	res, err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-user-edit")
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) Add(req *requests.UserAddRequest) (res string, err error) {
	now := time.Now().UTC()
	var password string

	count, err := uc.CountBy("u.email", "=", "", req.Email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-countByEmail")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.EmailAlreadyExist, functioncaller.PrintFuncName(), "uc-user-countByEmail")
		return res, errors.New(messages.EmailAlreadyExist)
	}

	model := models.NewUserModel().SetFirstName(req.FirstName).SetLastName(req.LastName).SetEmail(req.Email).SetCreatedAt(now).SetUpdatedAt(now).SetRoleId(req.RoleId)
	if req.Password != "" {
		password, _ = hashing.HashAndSalt(req.Password)
		model.SetPassword(password)
	}
	cmd := commands.NewUserCommand(uc.Config.DB, model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-user-add")
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) Delete(id string) (err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("u.id", "=", "", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-CountByID")
		return err
	}
	if count > 0 {
		model := models.NewUserModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)
		cmd := commands.NewUserCommand(uc.Config.DB, model)
		_, err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-user-delete")
			return err
		}
	}

	return nil
}

func (uc UserUseCase) Count(search string) (res int, err error) {
	q := queries.NewUserQuery(uc.Config.DB.GetDbInstance())

	res, err = q.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-count")
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	q := queries.NewUserQuery(uc.Config.DB.GetDbInstance())

	res, err = q.CountBy(column, operator, id, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-countBy")
		return res, err
	}

	return res, nil
}
