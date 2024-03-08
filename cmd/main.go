package main

import (
	"codecreeo/database"
	"codecreeo/internal/handler"
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
	app    *fiber.App
	DbPool database.DbConnection
}

func (a *App) Register() {
	a.app.Get("/monitor", handler.Monitor())
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// png, err := qrcode.Encode("https://www.google.com", qrcode.Medium, 256)
	dbConnection := database.NewDbConnection()
	defer database.NewDbConnection().CloseDbConnection()

	if err := database.NewDbConnection().AutoMigrate(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to automigrate tables: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Tables migrated successfully")

	app := fiber.New()

	app.Use(cors.New())
	app.Use(healthcheck.New())

	application := &App{app: app, DbPool: *dbConnection}
	application.Register()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		_ = <-c
		log.Fatal("application gracefully shutting down..")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":80"); err != nil {
		log.Fatalf(fmt.Sprintf("app error: %s", err.Error()))
	}
}
