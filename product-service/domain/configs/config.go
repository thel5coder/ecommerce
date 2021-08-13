package configs

import (
	"github.com/thel5coder/pkg/jwe"
	"github.com/thel5coder/pkg/jwt"
	"github.com/thel5coder/pkg/minio"
	"github.com/thel5coder/pkg/postgresql"
	"github.com/thel5coder/pkg/redis"
	"github.com/thel5coder/pkg/str"
	"github.com/thel5coder/pkg/validator"
	"log"
	"os"
)

type IConfig interface {
	SetDBConnection() *Config

	SetRedisConnection() *Config

	SetJwe() *Config

	SetJwt() *Config

	SetValidator() *Config

	SetMinioConnection() *Config
}

type Config struct {
	DB        postgresql.IConnection
	Redis     redis.IConnection
	Jwe       jwe.IJwe
	Jwt       jwt.IJwt
	Validator validator.IValidator
	Minio     minio.IConnection
}

func NewConfig() *Config {
	return &Config{}
}

//SetDBConnection connection to postgresql database
func (c *Config) SetDBConnection() *Config {
	config := postgresql.NewConfig(os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER_NAME"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_TIME_ZONE"), os.Getenv("DB_SSL_MODE"),
		str.StringToInt(os.Getenv("DB_MAX_CONNECTION")), str.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
		str.StringToInt(os.Getenv("DB_MAX_LIFETIME_CONNECTION")))

	c.DB = postgresql.NewConnection(config)

	return c
}

//SetRedisConnection set redis connection
func (c *Config) SetRedisConnection() *Config {
	c.Redis = redis.NewConnection(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD")).Connect()

	return c
}

//SetJwe set jwe configuration jwe key location and passphrase for key location
func (c *Config) SetJwe() *Config {
	c.Jwe = jwe.NewJwe(os.Getenv("JWE_KEY_LOCATION"), os.Getenv("JWE_PASSPHRASE"))

	return c
}

//SetJwt set jwt configuration for secret and expired time
func (c *Config) SetJwt() *Config {
	c.Jwt = jwt.NewJwt(os.Getenv("SECRET"), os.Getenv("SECRET_REFRESH_TOKEN"),
		str.StringToInt(os.Getenv("TOKEN_EXP_TIME")), str.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")))

	return c
}

//SetValidator configuration locale for validator
func (c *Config) SetValidator() *Config {
	c.Validator = validator.NewValidator(os.Getenv("APP_LOCALE")).SetValidator().SetTranslator()

	return c
}

//SetMinioConnection set configuration connection for min.io storage service
func (c *Config) SetMinioConnection() *Config {
	var err error
	useSsl := true
	if os.Getenv("MINIO_USE_SSL") == "false" {
		useSsl = false
	}

	minioConnection := minio.NewConnection()
	c.Minio, err = minioConnection.SetSecretKey(os.Getenv("MINIO_SECRET_KEY")).SetAccessKey(os.Getenv("MINIO_ACCESS_KEY")).SetEndPoint(os.Getenv("MINIO_ENDPOINT")).
		SetUseSsl(useSsl).Connect()
	if err != nil {
		log.Fatal(err)
	}

	return c
}
