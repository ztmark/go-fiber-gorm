package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"go-fiber-gorm/config"
	"go-fiber-gorm/handler"
	"log"
)

func main() {
	app := fiber.New()

	setupMiddleware(app)
	setupRoute(app)
	cfg := config.InitConfig()

	log.Fatalln(app.Listen(fmt.Sprintf(":%d", cfg.ServerCfg.Port)))

}

func setupMiddleware(app *fiber.App) {
	app.Use(logger.New())
	app.Use(recover2.New())
}

func setupRoute(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/hello", handler.Hello)
	api.Get("/post", handler.GetPost)
	api.Post("/post", handler.SavePost)
}
