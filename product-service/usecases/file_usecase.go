package usecases

import (
	"github.com/ecommerce-service/product-service/domain/usecase"
	"github.com/ecommerce-service/product-service/domain/view_models"
	"github.com/globalsign/mgo/bson"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/logruslogger"
	"github.com/thel5coder/pkg/minio"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileUseCase struct {
	*UseCaseContract
}

func NewFileUseCase(useCaseContract *UseCaseContract) usecase.IFileUseCase {
	return &FileUseCase{UseCaseContract: useCaseContract}
}

func (uc FileUseCase) Upload(file *multipart.FileHeader) (res view_models.FileVm, err error) {
	//upload minio
	fileName := bson.NewObjectId().Hex() + filepath.Ext(file.Filename)
	minioModel := minio.NewMinioModel(uc.Config.Minio.GetClient())
	_, err = minioModel.Upload(os.Getenv("MINIO_BUCKET"), fileName, file)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-minioModel-uploadFile")
		return res, err
	}

	//get file
	path, err := minioModel.GetFile(os.Getenv("MINIO_BUCKET"), fileName)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-minioModel-getFile")
		return res, err
	}
	res = view_models.NewFileVm(fileName, filepath.Ext(file.Filename), path)

	return res, nil
}

func (uc FileUseCase) GetUrlByKey(key string) (res view_models.FileVm, err error) {
	minioModel := minio.NewMinioModel(uc.Config.Minio.GetClient())
	path, err := minioModel.GetFile(os.Getenv("MINIO_BUCKET"), key)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-minioModel-getFile")
		return res, err
	}
	res = view_models.NewFileVm(key, filepath.Ext(key), path)

	return res, nil
}
