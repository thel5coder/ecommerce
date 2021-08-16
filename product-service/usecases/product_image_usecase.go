package usecases

import (
	"github.com/thel5coder/ecommerce/product-service/domain/models"
	"github.com/thel5coder/ecommerce/product-service/domain/usecase"
	productImageCmd "github.com/thel5coder/ecommerce/product-service/repository/commands"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/logruslogger"
)

type ProductImageUseCase struct {
	*UseCaseContract
}

//NewProductImageUseCase function to initialize new product image use case
func NewProductImageUseCase(useCaseContract *UseCaseContract) usecase.IProductImageUseCase {
	return ProductImageUseCase{
		UseCaseContract: useCaseContract,
	}
}

//Add use case function to add data into product_images table
func (uc ProductImageUseCase) Add(productID, imageKey string) (err error) {
	//Set product images model
	model := models.NewProductImageModel().SetProductId(productID).SetImageKey(imageKey)

	//initialize command product image to call command add function
	cmd := productImageCmd.NewProductImageCommand(uc.Config.DB, model)
	err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-productImage-add")
		return err
	}

	return nil
}

//Delete use case function to delete from product_images table
func (uc ProductImageUseCase) Delete(productID string) (err error) {
	//Set product images model
	model := models.NewProductImageModel().SetProductId(productID)

	//initialize command product image to call command delete function
	cmd := productImageCmd.NewProductImageCommand(uc.Config.DB, model)
	res, err := cmd.Delete()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-productImage-delete")
		return err
	}

	//check rows affected from command delete
	_, err = res.RowsAffected()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-productImage-rowResult")
		return err
	}

	return nil
}

//Store logical function to add and delete product_image tables
//delete all data inside product_images table by product id
//and then insert new product_images by imageKeys array
func (uc ProductImageUseCase) Store(productID string, imageKeys []string) (err error) {
	//call delete use case function
	err = uc.Delete(productID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productImage-delete")
		return err
	}

	//looping and call add usecase function
	for _, imageKey := range imageKeys {
		err = uc.Add(productID, imageKey)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-productImage-add")
			return err
		}
	}

	return nil
}
