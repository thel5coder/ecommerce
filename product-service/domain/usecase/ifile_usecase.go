package usecase

import (
	"github.com/ecommerce-service/product-service/domain/view_models"
	"mime/multipart"
)

type IFileUseCase interface {
	Upload(file *multipart.FileHeader) (res view_models.FileVm,err error)

	GetUrlByKey(key string) (res view_models.FileVm,err error)
}
