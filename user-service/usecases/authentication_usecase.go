package usecases

import (
	"errors"
	"github.com/ecommerce/user-service/domain/models"
	"github.com/ecommerce/user-service/domain/requests"
	"github.com/ecommerce/user-service/domain/usecase"
	"github.com/ecommerce/user-service/domain/view_models"
	"github.com/ecommerce/user-service/repository/queries"
	"github.com/thel5coder/pkg/functioncaller"
	"github.com/thel5coder/pkg/hashing"
	"github.com/thel5coder/pkg/logruslogger"
	"github.com/thel5coder/pkg/messages"
	"os"
)

type AuthenticationUseCase struct {
	*UseCaseContract
}

func NewAuthenticationUseCase(useCaseContract *UseCaseContract) usecase.IAuthenticationUseCase {
	return &AuthenticationUseCase{UseCaseContract: useCaseContract}
}

func (uc AuthenticationUseCase) Registration(req *requests.RegisterRequest) (err error) {
	userUc := NewUserUseCase(uc.UseCaseContract)
	userRequest := requests.UserAddRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		RoleId:    DefaultIDNormalUsers,
	}
	_, err = userUc.Add(&userRequest)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.CredentialDoNotMatch, functioncaller.PrintFuncName(), "uc-user-add")
		return err
	}

	return nil
}

func (uc AuthenticationUseCase) Login(req *requests.LoginRequest) (res view_models.LoginVm, err error) {
	user, valid := uc.ValidateCredential(req.Email, req.Password)
	if !valid {
		logruslogger.Log(logruslogger.WarnLevel, messages.CredentialDoNotMatch, functioncaller.PrintFuncName(), "uc-authentication-validateCredential")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	//generate jwt payload and encrypted with jwe
	payload := map[string]interface{}{
		"id":     user.ID,
		"roleID": user.Role.ID,
		"role":   user.Role.Name,
	}
	jwePayload, err := uc.Config.Jwe.GenerateJwePayload(payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-generate-jwe-payload")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	//generate jwt token
	res, err = uc.GenerateJWT(req.Email, jwePayload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-generate-jwt-token")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	userLoggedIn := map[string]interface{}{
		"email": user.Email,
		"role": map[string]interface{}{
			"id":   user.Role.ID,
			"name": user.Role.Name,
		},
	}
	err = uc.Config.Redis.StoreWithExpired(res.Token, userLoggedIn, os.Getenv("TOKEN_EXP_TIME")+`h`)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-generate-jwt-token")
		return res, err
	}

	return res, nil
}

func (uc AuthenticationUseCase) GenerateJWT(issuer, payload string) (res view_models.LoginVm, err error) {
	res.Token, res.TokenExpiredAt, err = uc.Config.Jwt.GenerateToken(issuer, payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-jwt-getToken")
		return res, err
	}

	res.RefreshToken, res.RefreshTokenExpiredAt, err = uc.Config.Jwt.GenerateRefreshToken(issuer, payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-jwt-getRefreshToken")
		return res, err
	}

	return res, nil
}

func (uc AuthenticationUseCase) ValidateCredential(username, password string) (view_models.UserVm, bool) {
	q := queries.NewUserQuery(uc.Config.DB.GetDbInstance())
	var res view_models.UserVm

	user, err := q.ReadBy("u.email", "=", username)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-readByEmail")
		return res, false
	}

	model := user.(*models.User)
	isValid := hashing.CheckHashString(password, model.Password())
	res = view_models.NewUserVm(model)

	return res, isValid
}
