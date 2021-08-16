package usecases

import (
	"github.com/gosimple/slug"
	"github.com/thel5coder/ecommerce/product-service/domain/models"
	"github.com/thel5coder/ecommerce/product-service/domain/requests"
	"github.com/thel5coder/ecommerce/product-service/domain/usecase"
	"github.com/thel5coder/ecommerce/product-service/domain/view_models"
	"github.com/thel5coder/ecommerce/product-service/repository/commands"
	"github.com/thel5coder/ecommerce/product-service/repository/queries"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/logruslogger"
	"github.com/thel5coder/pkg/messages"
	"time"
)

type CategoryUseCase struct {
	*UseCaseContract
}

//NewCategoryUseCase function to initialize new category use case
func NewCategoryUseCase(useCaseContract *UseCaseContract) usecase.ICategoryUseCase {
	return &CategoryUseCase{UseCaseContract: useCaseContract}
}

//GetListWithPagination function to get list of categories data with pagination
func (uc CategoryUseCase) GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CategoryVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewCategoryQuery(uc.Config.DB)

	//set pagination parameter
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	//calling query browse of category
	categories, err := q.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-category-browse")
		return res, pagination, err
	}

	//looping data from models to parse into category view models
	for _, category := range categories.([]*models.Category) {
		res = append(res, view_models.NewCategoryVm(category))
	}

	//set pagination response
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-category-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

//GetAll function to get all categories data
func (uc CategoryUseCase) GetAll(search string) (res []view_models.CategoryVm, err error) {
	q := queries.NewCategoryQuery(uc.Config.DB)

	//calling browse all from query category
	categories, err := q.BrowseAll(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-category-browseAll")
		return res, err
	}

	//looping data from models to parse into category view models
	for _, category := range categories.([]*models.Category) {
		res = append(res, view_models.NewCategoryVm(category))
	}

	return res, nil
}

//GetByID function to get category data by id
func (uc CategoryUseCase) GetByID(id string) (res view_models.CategoryVm, err error) {
	q := queries.NewCategoryQuery(uc.Config.DB)

	//calling read by from query category
	category, err := q.ReadBy("id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-category-browseAll")
		return res, err
	}
	//parse data from models into category view models
	res = view_models.NewCategoryVm(category.(*models.Category))

	return res, nil
}

//Edit function to edit category data from product table
func (uc CategoryUseCase) Edit(req *requests.CategoryRequest, id string) (res string, err error) {
	//making category slug from category name
	categorySlug := slug.Make(req.Name)

	//calling CountBy Use Case function to count data inside category table by slug
	//if count more than zero it means the category name already exist and the function return error
	//with message data already exist
	count, err := uc.CountBy("slug", "=", id, categorySlug)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-category-countByName")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-category-countByName")
	}

	//set category model data from category request struct
	now := time.Now().UTC()
	model := models.NewCategoryModel().SetName(req.Name).SetSlug(categorySlug).SetUpdatedAt(now).SetId(id)

	//initialize command category to call command edit function
	cmd := commands.NewCategoryCommand(uc.Config.DB, model)
	res, err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-category-edit")
		return res, err
	}

	return res, nil
}

//Add use case function to add new category into categories table
func (uc CategoryUseCase) Add(req *requests.CategoryRequest) (res string, err error) {
	//making category slug from category name
	categorySlug := slug.Make(req.Name)

	//calling CountBy Use Case function to count data inside category table by slug
	//if count more than zero it means the category name already exist and the function return error
	//with message data already exist
	count, err := uc.CountBy("slug", "=", "", categorySlug)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-category-countByName")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-category-countByName")
	}

	//set category model data from category request struct
	now := time.Now().UTC()
	model := models.NewCategoryModel().SetName(req.Name).SetSlug(categorySlug).SetCreatedAt(now).SetUpdatedAt(now)

	//initialize command category to call command add function
	cmd := commands.NewCategoryCommand(uc.Config.DB, model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-category-add")
		return res, err
	}

	return res, nil
}

//Delete use case function to delete category data from categories table
func (uc CategoryUseCase) Delete(id string) (err error) {
	//calling CountBy Use Case functions to count data inside categories table by id
	//if count more than zero then call category command delete
	count, err := uc.CountBy("id", "=", "", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-category-countById")
		return err
	}

	if count > 0 {
		//set category model data
		now := time.Now().UTC()
		model := models.NewCategoryModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)

		//initialize command category to call command delete function
		cmd := commands.NewCategoryCommand(uc.Config.DB, model)
		_, err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-category-delete")
			return err
		}
	}

	return nil
}

//Count use case function to count category data from categories table
func (uc CategoryUseCase) Count(search string) (res int, err error) {
	q := queries.NewCategoryQuery(uc.Config.DB)

	//call query count function
	res, err = q.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-category-count")
		return res, err
	}

	return res, nil
}

//CountBy use case function to count category data from categories table by some specific column by parameters that pass from handler
func (uc CategoryUseCase) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	q := queries.NewCategoryQuery(uc.Config.DB)

	//call query count by function
	res, err = q.CountBy(column, operator, id, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-category-count")
		return res, err
	}

	return res, nil
}
