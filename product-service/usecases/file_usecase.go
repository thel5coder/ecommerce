package usecases

import (
	"github.com/globalsign/mgo/bson"
	"github.com/thel5coder/ecommerce/product-service/domain/usecase"
	"github.com/thel5coder/ecommerce/product-service/domain/view_models"
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

//NewFileUseCase function to initialize new file use case
func NewFileUseCase(useCaseContract *UseCaseContract) usecase.IFileUseCase {
	return &FileUseCase{UseCaseContract: useCaseContract}
}

//Upload logical function to upload file into min.io
func (uc FileUseCase) Upload(file *multipart.FileHeader) (res view_models.FileVm, err error) {
	//Set min.io model
	fileName := bson.NewObjectId().Hex() + filepath.Ext(file.Filename)
	minioModel := minio.NewMinioModel(uc.Config.Minio.GetClient())
	//upload file into min.io
	_, err = minioModel.Upload(os.Getenv("MINIO_BUCKET"), fileName, file)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-minioModel-uploadFile")
		return res, err
	}

	//calling GetFile use case function to get file path
	path, err := minioModel.GetFile(os.Getenv("MINIO_BUCKET"), fileName)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-minioModel-getFile")
		return res, err
	}

	//parse minio model into file view models
	res = view_models.NewFileVm(fileName, filepath.Ext(file.Filename), path)

	return res, nil
}

//GetUrlByKey logical function to get file from min.io by ky
func (uc FileUseCase) GetUrlByKey(key string) (res view_models.FileVm, err error) {
	//Set min.io model
	minioModel := minio.NewMinioModel(uc.Config.Minio.GetClient())

	//get file url
	path, err := minioModel.GetFile(os.Getenv("MINIO_BUCKET"), key)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-minioModel-getFile")
		return res, err
	}

	//parse minio model into file view models
	res = view_models.NewFileVm(key, filepath.Ext(key), path)

	return res, nil
}
