package usecases

import (
	"github.com/ecommerce-service/product-service/domain/models"
	"github.com/ecommerce-service/product-service/domain/requests"
	"github.com/ecommerce-service/product-service/domain/usecase"
	"github.com/ecommerce-service/product-service/domain/view_models"
	"github.com/ecommerce-service/product-service/repository/commands"
	"github.com/ecommerce-service/product-service/repository/queries"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/logruslogger"
	"github.com/thel5coder/pkg/messages"
	"strings"
	"time"
)

type ProductUseCase struct {
	*UseCaseContract
}

func NewProductUseCase(useCaseContract *UseCaseContract) usecase.IProductUseCase {
	return &ProductUseCase{UseCaseContract: useCaseContract}
}

func (uc ProductUseCase) GetListWithPagination(search, orderBy, sort, category string, page, limit int) (res []view_models.ListProductVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewProductQuery(uc.Config.DB)
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)
	fileUc := NewFileUseCase(uc.UseCaseContract)

	products, err := q.Browse(search, orderBy, sort, category, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-browse")
		return res, pagination, err
	}
	for _, product := range products.([]*models.Product) {
		var mainImage view_models.FileVm
		if product.MainImageKey().String != "" {
			mainImage, err = fileUc.GetUrlByKey(product.MainImageKey().String)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-file-getUrlByKey")
			}
		}
		res = append(res, view_models.NewListProductVm(product, mainImage))
	}

	//set pagination
	totalCount, err := uc.Count(search, category)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-product-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc ProductUseCase) GetByID(id string) (res view_models.DetailProductVm, err error) {
	q := queries.NewProductQuery(uc.Config.DB)

	product, err := q.ReadBy("p.id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-readById")
		return res, err
	}

	model := product.(*models.Product)
	var mainImage view_models.FileVm
	fileUc := NewFileUseCase(uc.UseCaseContract)
	if model.MainImageKey().String != "" {
		mainImage, err = fileUc.GetUrlByKey(model.MainImageKey().String)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-file-getUrlByKeyMainImage")
		}
	}

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
	res = view_models.NewDetailProductVm(model, mainImage, productImagesVm)

	return res, nil
}

func (uc ProductUseCase) Edit(req *requests.ProductRequest, id string) (res string, err error) {
	count, err := uc.CountBy("p.sku", "=", id, req.Sku)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-product-countBySku")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-product-countBySku")
		return res, err
	}

	productImageUc := NewProductImageUseCase(uc.UseCaseContract)
	err = productImageUc.Store(id, req.ProductImages)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productImage-store")
		return res, err
	}

	now := time.Now().UTC()
	model := models.NewProductModel().SetCategoryId(req.CategoryId).SetName(req.Name).SetSku(req.Sku).SetPrice(req.Price).
		SetDiscount(req.Discount).SetMainImageKey(req.MainImage).SetUpdatedAt(now).SetId(id)
	cmd := commands.NewProductCommand(uc.Config.DB, model)
	err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-edit")
		return res, err
	}
	res = id

	return res, nil
}

func (uc ProductUseCase) Add(req *requests.ProductRequest) (res string, err error) {
	count, err := uc.CountBy("p.sku", "=", "", req.Sku)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-product-countBySku")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-product-countBySku")
		return res, err
	}

	now := time.Now().UTC()
	model := models.NewProductModel().SetCategoryId(req.CategoryId).SetName(req.Name).SetSku(req.Sku).SetPrice(req.Price).
		SetDiscount(req.Discount).SetMainImageKey(req.MainImage).SetCreatedAt(now).SetUpdatedAt(now)
	cmd := commands.NewProductCommand(uc.Config.DB, model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-add")
		return res, err
	}

	productImageUc := NewProductImageUseCase(uc.UseCaseContract)
	err = productImageUc.Store(res, req.ProductImages)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productImage-store")
		return res, err
	}

	return res, nil
}

func (uc ProductUseCase) Delete(id string) (err error) {
	count, err := uc.CountBy("p.id", "=", "", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-product-countById")
		return err
	}

	if count > 0 {
		now := time.Now().UTC()
		model := models.NewProductModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)
		cmd := commands.NewProductCommand(uc.Config.DB, model)
		err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-delete")
			return err
		}

		productImageUc := NewProductImageUseCase(uc.UseCaseContract)
		err = productImageUc.Delete(id)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productImage-delete")
			return err
		}
	}

	return nil
}

func (uc ProductUseCase) Count(search, category string) (res int, err error) {
	q := queries.NewProductQuery(uc.Config.DB)

	res, err = q.Count(search, category)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-count")
		return res, err
	}

	return res, nil
}

func (uc ProductUseCase) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	q := queries.NewProductQuery(uc.Config.DB)

	res, err = q.CountBy(column, operator, id, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-product-countBy")
		return res, err
	}

	return res, nil
}
