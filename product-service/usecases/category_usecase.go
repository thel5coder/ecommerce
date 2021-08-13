package usecases

import (
	"github.com/ecommerce-service/product-service/domain/models"
	"github.com/ecommerce-service/product-service/domain/requests"
	"github.com/ecommerce-service/product-service/domain/usecase"
	"github.com/ecommerce-service/product-service/domain/view_models"
	"github.com/ecommerce-service/product-service/repository/commands"
	"github.com/ecommerce-service/product-service/repository/queries"
	"github.com/gosimple/slug"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/logruslogger"
	"github.com/thel5coder/pkg/messages"
	"time"
)

type CategoryUseCase struct {
	*UseCaseContract
}

func NewCategoryUseCase(useCaseContract *UseCaseContract) usecase.ICategoryUseCase {
	return &CategoryUseCase{UseCaseContract: useCaseContract}
}

func (uc CategoryUseCase) GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CategoryVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewCategoryQuery(uc.Config.DB)
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	categories, err := q.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-category-browse")
		return res, pagination, err
	}
	for _, category := range categories.([]*models.Category) {
		res = append(res, view_models.NewCategoryVm(category))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-category-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc CategoryUseCase) GetAll(search string) (res []view_models.CategoryVm, err error) {
	q := queries.NewCategoryQuery(uc.Config.DB)

	categories, err := q.BrowseAll(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-category-browseAll")
		return res, err
	}
	for _, category := range categories.([]*models.Category) {
		res = append(res, view_models.NewCategoryVm(category))
	}

	return res, nil
}

func (uc CategoryUseCase) GetByID(id string) (res view_models.CategoryVm, err error) {
	q := queries.NewCategoryQuery(uc.Config.DB)

	category, err := q.ReadBy("id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-category-browseAll")
		return res, err
	}
	res = view_models.NewCategoryVm(category.(*models.Category))

	return res, nil
}

func (uc CategoryUseCase) Edit(req *requests.CategoryRequest, id string) (res string, err error) {
	categorySlug := slug.Make(req.Name)
	count, err := uc.CountBy("slug", "=", id, categorySlug)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-category-countByName")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-category-countByName")
	}

	now := time.Now().UTC()
	model := models.NewCategoryModel().SetName(req.Name).SetSlug(categorySlug).SetUpdatedAt(now).SetId(id)
	cmd := commands.NewCategoryCommand(uc.Config.DB, model)
	res, err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-category-edit")
		return res, err
	}

	return res, nil
}

func (uc CategoryUseCase) Add(req *requests.CategoryRequest) (res string, err error) {
	categorySlug := slug.Make(req.Name)
	count, err := uc.CountBy("slug", "=", "", categorySlug)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-category-countByName")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-category-countByName")
	}

	now := time.Now().UTC()
	model := models.NewCategoryModel().SetName(req.Name).SetSlug(categorySlug).SetCreatedAt(now).SetUpdatedAt(now)
	cmd := commands.NewCategoryCommand(uc.Config.DB, model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-category-add")
		return res, err
	}

	return res, nil
}

func (uc CategoryUseCase) Delete(id string) (err error) {
	count, err := uc.CountBy("id", "=", "", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-category-countById")
		return err
	}

	if count > 0 {
		now := time.Now().UTC()
		model := models.NewCategoryModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)
		cmd := commands.NewCategoryCommand(uc.Config.DB, model)
		_, err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-category-delete")
			return err
		}
	}

	return nil
}

func (uc CategoryUseCase) Count(search string) (res int, err error) {
	q := queries.NewCategoryQuery(uc.Config.DB)

	res, err = q.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-category-count")
		return res, err
	}

	return res, nil
}

func (uc CategoryUseCase) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	q := queries.NewCategoryQuery(uc.Config.DB)

	res, err = q.CountBy(column, operator, id, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-category-count")
		return res, err
	}

	return res, nil
}
