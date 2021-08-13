package usecases

import (
	"database/sql"
	"github.com/ecommerce-service/transaction-service/domain/models"
	"github.com/ecommerce-service/transaction-service/domain/requests"
	"github.com/ecommerce-service/transaction-service/domain/usecases"
	"github.com/ecommerce-service/transaction-service/domain/view_models"
	"github.com/ecommerce-service/transaction-service/repository/commands"
	"github.com/ecommerce-service/transaction-service/repository/queries"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/logruslogger"
	"time"
)

type CartUseCase struct {
	*UseCaseContract
}

func NewCartUseCase(useCaseContract *UseCaseContract) usecases.ICartUseCase {
	return &CartUseCase{UseCaseContract: useCaseContract}
}

func (uc CartUseCase) GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CartVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewCartQuery(uc.Config.DB)
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	cart, err := q.BrowseByUser(search, orderBy, sort, uc.UserID, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-browse")
		return res, pagination, err
	}
	for _, cart := range cart.([]*models.Carts) {
		res = append(res, view_models.NewCartVm(cart))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc CartUseCase) GetAllByUserId(userId string) (res []view_models.CartVm, err error) {
	q := queries.NewCartQuery(uc.Config.DB)

	cart, err := q.BrowseAllByUser(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-browseAllByUser")
		return res, err
	}
	for _, cart := range cart.([]*models.Carts) {
		res = append(res, view_models.NewCartVm(cart))
	}

	return res, nil
}

func (uc CartUseCase) GetByID(id string) (res view_models.CartVm, err error) {
	q := queries.NewCartQuery(uc.Config.DB)

	cart, err := q.ReadBy("c.id", "=", uc.UserID, id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-readByID")
		return res, err
	}
	res = view_models.NewCartVm(cart.(*models.Carts))

	return res, nil
}

func (uc CartUseCase) Edit(req *requests.CartRequest, id string) (res string, err error) {
	now := time.Now().UTC()
	model := models.NewCartModel()
	subTotal := float64(req.Quantity) * req.Price
	model.SetUserId(uc.UserID).SetProductId(req.ProductID).SetPrice(req.Price).SetQuantity(req.Quantity).SetSubTotal(subTotal).SetUpdatedAt(now).SetId(id)

	cmd := commands.NewCartCommand(uc.Config.DB, model)
	res, err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-cart-edit")
		return res, err
	}

	return res, nil
}

func (uc CartUseCase) Add(req *requests.CartRequest) (res string, err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("c.product_id", "=", req.ProductID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-countByProductId")
		return res, err
	}

	if count > 0 {
		cart, err := uc.GetBy("c.product_id", "=", req.ProductID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-getByProductId")
			return res, err
		}

		editQuantityRequest := requests.CartRequest{Quantity: req.Quantity + cart.Quantity}
		res, err = uc.Edit(&editQuantityRequest, cart.ID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-editQuantity")
			return res, err
		}
	} else {
		subTotal := float64(req.Quantity) * req.Price
		model := models.NewCartModel().SetUserId(uc.UserID).SetProductId(req.ProductID).SetName(req.Name).SetPrice(req.Price).SetQuantity(req.Quantity).SetSubTotal(subTotal).
			SetSku(req.Sku).SetCategory(req.Category).SetCreatedAt(now).SetUpdatedAt(now)
		cmd := commands.NewCartCommand(uc.Config.DB, model)
		res, err = cmd.Add()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-cart-add")
			return res, err
		}
	}

	return res, nil
}

func (uc CartUseCase) Delete(id string) (err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("c.id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-countBy")
		return err
	}
	if count > 0 {
		model := models.NewCartModel().SetUpdatedAt(now).SetDeletedAt(sql.NullTime{Time: now, Valid: true}).SetId(id)
		cmd := commands.NewCartCommand(uc.Config.DB, model)
		_, err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-cart-delete")
			return err
		}
	}

	return nil
}

func (uc CartUseCase) DeleteAllByUserId() (err error) {
	now := time.Now().UTC()

	count,err := uc.CountBy("c.user_id","=",uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-countBy")
		return err
	}

	if count > 0 {
		model := models.NewCartModel().SetUpdatedAt(now).SetDeletedAt(sql.NullTime{Time: now, Valid: true}).SetUserId(uc.UserID)
		cmd := commands.NewCartCommand(uc.Config.DB, model)
		err = cmd.DeleteAllByUserID()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-cart-deleteAllByUserId")
			return err
		}
	}

	return nil
}

func (uc CartUseCase) Count(search string) (res int, err error) {
	q := queries.NewCartQuery(uc.Config.DB)

	res, err = q.Count(search, uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-count")
		return res, err
	}

	return res, nil
}

func (uc CartUseCase) CountBy(column, operator string, value interface{}) (res int, err error) {
	q := queries.NewCartQuery(uc.Config.DB)

	res, err = q.CountBy(column, operator, uc.UserID, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-countBy")
		return res, err
	}

	return res, nil
}

func (uc CartUseCase) GetBy(column, operator string, value interface{}) (res view_models.CartVm, err error) {
	q := queries.NewCartQuery(uc.Config.DB)

	cart, err := q.ReadBy(column, operator, uc.UserID, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-getBy")
		return res, err
	}
	res = view_models.NewCartVm(cart.(*models.Carts))

	return res, nil
}
