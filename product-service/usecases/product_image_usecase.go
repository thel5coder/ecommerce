package usecases

import (
	"github.com/ecommerce-service/product-service/domain/models"
	"github.com/ecommerce-service/product-service/domain/usecase"
	productImageCmd "github.com/ecommerce-service/product-service/repository/commands"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/logruslogger"
)

type ProductImageUseCase struct {
	*UseCaseContract
}

func NewProductImageUseCase(useCaseContract *UseCaseContract) usecase.IProductImageUseCase {
	return ProductImageUseCase{
		UseCaseContract: useCaseContract,
	}
}

func (uc ProductImageUseCase) Add(productID, imageKey string) (err error) {
	model := models.NewProductImageModel().SetProductId(productID).SetImageKey(imageKey)
	cmd := productImageCmd.NewProductImageCommand(uc.Config.DB, model)

	err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-productImage-add")
		return err
	}

	return nil
}

func (uc ProductImageUseCase) Delete(productID string) (err error) {
	model := models.NewProductImageModel().SetProductId(productID)
	cmd := productImageCmd.NewProductImageCommand(uc.Config.DB, model)

	res, err := cmd.Delete()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-productImage-delete")
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-productImage-rowResult")
		return err
	}

	return nil
}

func (uc ProductImageUseCase) Store(productID string, imageKeys []string) (err error) {
	err = uc.Delete(productID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productImage-delete")
		return err
	}

	for _, imageKey := range imageKeys {
		err = uc.Add(productID, imageKey)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productImage-add")
			return err
		}
	}

	return nil
}
