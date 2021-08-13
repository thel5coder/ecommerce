package main

import (
	"github.com/ecommerce-service/transaction-service/domain/configs"
	bootApp "github.com/ecommerce-service/transaction-service/server/http/boot"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var (
	logFormat = `{"host":"${host}","pid":"${pid}","time":"${time}","request-id":"${locals:requestid}","status":"${status}","method":"${method}","latency":"${latency}","path":"${path}",` +
		`"user-agent":"${ua}","in":"${bytesReceived}","out":"${bytesSent}"}`
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	//load config
	config := configs.NewConfig().SetRedisConnection().SetValidator().SetJwe().SetJwt().SetDBConnection()
	db, err := config.DB.Connect()
	if err != nil {
		log.Fatal(err)
	}
	db.Pool()
	//db.Migration(os.Getenv("DB_MIGRATIONS_DIRECTORY"))
	defer db.GetDbInstance().Close()

	//initialization of fiber instance
	app := fiber.New()
	boot := bootApp.NewBoot(app, config)
	boot.App.Use(recover.New())
	boot.App.Use(requestid.New())
	boot.App.Use(cors.New())
	boot.App.Use(logger.New(logger.Config{
		Format:     logFormat + "\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Jakarta",
	}, logger.Config{}))

	//register all routers
	boot.RegisterAllRouters()

	//start http server
	log.Fatal(boot.App.Listen(os.Getenv("APP_HOST")))
}
