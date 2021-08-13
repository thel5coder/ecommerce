package usecase

import (
	"github.com/ecommerce/user-service/domain/requests"
	"github.com/ecommerce/user-service/domain/view_models"
)

type IAuthenticationUseCase interface {
	Login(req *requests.LoginRequest) (res view_models.LoginVm, err error)

	Registration(req *requests.RegisterRequest) (err error)

	GenerateJWT(issuer, payload string) (res view_models.LoginVm, err error)

	ValidateCredential(username, password string) (view_models.UserVm, bool)
}
