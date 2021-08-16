package usecases

import (
	"github.com/thel5coder/ecommerce/product-service/domain/models"
	"github.com/thel5coder/ecommerce/product-service/domain/requests"
	"github.com/thel5coder/ecommerce/product-service/domain/usecase"
	"github.com/thel5coder/ecommerce/product-service/domain/view_models"
	"github.com/thel5coder/ecommerce/product-service/repository/commands"
	"github.com/thel5coder/ecommerce/product-service/repository/queries"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/logruslogger"
	"github.com/thel5coder/pkg/messages"
	"strings"
	"time"
)

type ProductUseCase struct {
	*UseCaseContract
}

//NewProductUseCase function to initialize new product use case
func NewProductUseCase(useCaseContract *UseCaseContract) usecase.IProductUseCase {
	return &ProductUseCase{UseCaseContract: useCaseContract}
}

//GetListWithPagination function to get list of products data with pagination
func (uc ProductUseCase) GetListWithPagination(search, orderBy, sort, category string, page, limit int) (res []view_models.ListProductVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewProductQuery(uc.Config.DB)

	//set pagination parameter
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	//initialization file use case
	fileUc := NewFileUseCase(uc.UseCaseContract)

	//calling query browse of product
	products, err := q.Browse(search, orderBy, sort, category, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-browse")
		return res, pagination, err
	}

	//looping data from models to parse into list product view models
	for _, product := range products.([]*models.Product) {
		//get file url from file use case get url by key
		var mainImage view_models.FileVm
		if product.MainImageKey().String != "" {
			mainImage, err = fileUc.GetUrlByKey(product.MainImageKey().String)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-file-getUrlByKey")
			}
		}
		res = append(res, view_models.NewListProductVm(product, mainImage))
	}

	//set pagination response
	totalCount, err := uc.Count(search, category)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-product-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

//GetByID function to get category data by id
func (uc ProductUseCase) GetByID(id string) (res view_models.DetailProductVm, err error) {
	q := queries.NewProductQuery(uc.Config.DB)

	//calling read by from query product
	product, err := q.ReadBy("p.id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-readById")
		return res, err
	}

	//get main image url from min.io and parse it into file view models
	model := product.(*models.Product)
	var mainImage view_models.FileVm
	fileUc := NewFileUseCase(uc.UseCaseContract)
	if model.MainImageKey().String != "" {
		mainImage, err = fileUc.GetUrlByKey(model.MainImageKey().String)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-file-getUrlByKeyMainImage")
		}
	}

	//get product images url from min.io and parse it into file view models
	var productImagesVm []view_models.FileVm
	if model.ProductImages().String != "" {
		productImages := strings.Split(model.ProductImages().String, ",")
		for _, productImage := range productImages {
			productImageVm, err := fileUc.GetUrlByKey(productImage)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-file-getUrlByKeyProductImage")
			}
			productImagesVm = append(productImagesVm, productImageVm)
		}
	}

	//parse data from models into product detail view models
	res = view_models.NewDetailProductVm(model, mainImage, productImagesVm)

	return res, nil
}

//Edit function to edit product data from product table
func (uc ProductUseCase) Edit(req *requests.ProductRequest, id string) (res string, err error) {
	//calling CountBy Use Case functions to count data inside products table by sku
	//if count more than zero it means the product sku already exist and the function return error
	//with message data already exist
	count, err := uc.CountBy("p.sku", "=", id, req.Sku)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-product-countBySku")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-product-countBySku")
		return res, err
	}

	//initialize product image use case and calling store function to insert product images into product_images table
	productImageUc := NewProductImageUseCase(uc.UseCaseContract)
	err = productImageUc.Store(id, req.ProductImages)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productImage-store")
		return res, err
	}

	//set product model data from product request struct
	now := time.Now().UTC()
	model := models.NewProductModel().SetCategoryId(req.CategoryId).SetName(req.Name).SetSku(req.Sku).SetPrice(req.Price).
		SetDiscount(req.Discount).SetMainImageKey(req.MainImage).SetUpdatedAt(now).SetId(id)

	//initialize command product to call command edit function
	cmd := commands.NewProductCommand(uc.Config.DB, model)
	err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-edit")
		return res, err
	}
	res = id

	return res, nil
}

//Add use case function to add new category into categories table
func (uc ProductUseCase) Add(req *requests.ProductRequest) (res string, err error) {
	//calling CountBy Use Case functions to count data inside products table by sku
	//if count more than zero it means the product sku already exist and the function return error
	//with message data already exist
	count, err := uc.CountBy("p.sku", "=", "", req.Sku)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-product-countBySku")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-product-countBySku")
		return res, err
	}

	//set product model data from product request struct
	now := time.Now().UTC()
	model := models.NewProductModel().SetCategoryId(req.CategoryId).SetName(req.Name).SetSku(req.Sku).SetPrice(req.Price).
		SetDiscount(req.Discount).SetMainImageKey(req.MainImage).SetCreatedAt(now).SetUpdatedAt(now)

	//initialize command product to call command add function
	cmd := commands.NewProductCommand(uc.Config.DB, model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-add")
		return res, err
	}

	//initialize product image use case and calling store function to insert product images into product_images table
	productImageUc := NewProductImageUseCase(uc.UseCaseContract)
	err = productImageUc.Store(res, req.ProductImages)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productImage-store")
		return res, err
	}

	return res, nil
}

//Delete use case function to delete product data from products table
func (uc ProductUseCase) Delete(id string) (err error) {
	//calling CountBy Use Case functions to count data inside products table by id
	//if count more than zero then call category command delete
	count, err := uc.CountBy("p.id", "=", "", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-product-countById")
		return err
	}


	if count > 0 {
		//set product model data
		now := time.Now().UTC()
		model := models.NewProductModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)

		//initialize command product to call command delete function
		cmd := commands.NewProductCommand(uc.Config.DB, model)
		err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-delete")
			return err
		}

		//initialize product use case and call delete function
		productImageUc := NewProductImageUseCase(uc.UseCaseContract)
		err = productImageUc.Delete(id)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productImage-delete")
			return err
		}
	}

	return nil
}

//Count use case function to count product data from products table
func (uc ProductUseCase) Count(search, category string) (res int, err error) {
	q := queries.NewProductQuery(uc.Config.DB)

	//call query count function
	res, err = q.Count(search, category)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-count")
		return res, err
	}

	return res, nil
}

//CountBy use case function to count product data from products table by some specific column by parameters that pass from handler
func (uc ProductUseCase) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	q := queries.NewProductQuery(uc.Config.DB)

	//call query count by function
	res, err = q.CountBy(column, operator, id, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-countBy")
		return res, err
	}

	return res, nil
}
