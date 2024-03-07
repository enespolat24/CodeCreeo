package main

import (
	"codecreeo/handler"
	"codecreeo/repository"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/joho/godotenv"
)

type App struct {
	app  *fiber.App
	repo *repository.Repository
}

func (a *App) Register() {
	a.app.Get("/monitor", handler.Monitor())
}

func main() {
	// png, err := qrcode.Encode("https://www.google.com", qrcode.Medium, 256)
	repo := repository.New()
	defer repo.Close()

	app := fiber.New()

	app.Use(cors.New())
	app.Use(healthcheck.New())

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	application := &App{app: app, repo: repo}
	application.Register()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		_ = <-c
		log.Fatalf("application gracefully shutting down..")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":80"); err != nil {
		log.Fatalf(fmt.Sprintf("app error: %s", err.Error()))
	}
}
