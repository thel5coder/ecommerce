package middlewares

import (

	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ecommerce-service/transaction-service/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/thel5coder/pkg/functioncaller"
	jwtPkg "github.com/thel5coder/pkg/jwt"
	"github.com/thel5coder/pkg/logruslogger"
	"github.com/thel5coder/pkg/messages"
	"github.com/thel5coder/pkg/response"
	"github.com/thel5coder/pkg/str"
	"strings"
	"time"
)

type JwtMiddleware struct {
	*usecases.UseCaseContract
}

func NewJwtMiddleware(contract *usecases.UseCaseContract) JwtMiddleware {
	return JwtMiddleware{UseCaseContract: contract}
}

func (m JwtMiddleware) Use(ctx *fiber.Ctx) error {
	claims := &jwtPkg.CustomClaims{}

	//check header is present or not
	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		logruslogger.Log(logruslogger.WarnLevel, messages.AuthHeaderNotPresent, functioncaller.PrintFuncName(), "middleware-jwt-checkHeader")
		return response.NewResponse(response.NewResponseUnauthorized(errors.New(messages.Unauthorized))).Send(ctx)
	}

	//check claims and signing method
	token := strings.Replace(header, "Bearer ", "", -1)
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			logruslogger.Log(logruslogger.WarnLevel, messages.UnexpectedSigningMethod, functioncaller.PrintFuncName(), "middleware-jwt-checkSigningMethod")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := []byte(m.Config.Jwt.GetTokenSecret())
		return secret, nil
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "middleware-jwt-checkClaims")
		return response.NewResponse(response.NewResponseUnauthorized(errors.New(messages.Unauthorized))).Send(ctx)
	}

	//check token live time
	if claims.ExpiresAt < time.Now().Unix() {
		logruslogger.Log(logruslogger.WarnLevel, messages.ExpiredToken, functioncaller.PrintFuncName(), "middleware-jwt-checkTokenLiveTime")
		return response.NewResponse(response.NewResponseUnauthorized(errors.New(messages.Unauthorized))).Send(ctx)
	}

	//jwe roll back encrypted id
	jweRes, err := m.Config.Jwe.Rollback(claims.Payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-jwe-rollback")
		return response.NewResponse(response.NewResponseUnauthorized(errors.New(messages.Unauthorized))).Send(ctx)
	}
	if jweRes == nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "pkg-jwe-resultNil")
		return response.NewResponse(response.NewResponseUnauthorized(errors.New(messages.Unauthorized))).Send(ctx)
	}

	//set id to uce case contract
	roleID := fmt.Sprintf("%v", jweRes["roleID"])
	m.UseCaseContract.UserID = fmt.Sprintf("%v", jweRes["id"])
	m.UseCaseContract.RoleID = str.StringToInt(roleID)

	/*var userLoggedIn map[string]interface{}
	err = m.Config.Redis.GetData(token, &userLoggedIn)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "pkg-redis-getFromRedis")
		return response.NewResponse(response.NewResponseUnauthorized(errors.New(messages.Unauthorized))).Send(ctx)
	}*/

	return ctx.Next()
}
