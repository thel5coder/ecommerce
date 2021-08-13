package configs

import (
	"github.com/thel5coder/pkg/jwe"
	"github.com/thel5coder/pkg/jwt"
	"github.com/thel5coder/pkg/postgresql"
	"github.com/thel5coder/pkg/redis"
	"github.com/thel5coder/pkg/str"
	"github.com/thel5coder/pkg/validator"
	"os"
)

type IConfig interface {
	SetDBConnection() *Config

	SetRedisConnection() *Config

	SetJwe() *Config

	SetJwt() *Config

	SetValidator() *Config
}

type Config struct {
	DB        postgresql.IConnection
	Redis     redis.IConnection
	Jwe       jwe.IJwe
	Jwt       jwt.IJwt
	Validator validator.IValidator
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetDBConnection() *Config {
	config := postgresql.NewConfig(os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER_NAME"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_TIME_ZONE"), os.Getenv("DB_SSL_MODE"),
		str.StringToInt(os.Getenv("DB_MAX_CONNECTION")), str.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
		str.StringToInt(os.Getenv("DB_MAX_LIFETIME_CONNECTION")))

	c.DB = postgresql.NewConnection(config)

	return c
}

func (c *Config) SetRedisConnection() *Config {
	c.Redis = redis.NewConnection(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD")).Connect()

	return c
}

func (c *Config) SetJwe() *Config {
	c.Jwe = jwe.NewJwe(os.Getenv("JWE_KEY_LOCATION"),os.Getenv("JWE_PASSPHRASE"))

	return c
}

func (c *Config) SetJwt() *Config {
	c.Jwt = jwt.NewJwt(os.Getenv("SECRET"),os.Getenv("SECRET_REFRESH_TOKEN"),
		str.StringToInt(os.Getenv("TOKEN_EXP_TIME")),str.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")))

	return c
}

func (c *Config) SetValidator() *Config {
	c.Validator = validator.NewValidator(os.Getenv("APP_LOCALE")).SetValidator().SetTranslator()

	return c
}
