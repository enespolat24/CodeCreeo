package main

import (
	"codecreeo/internal/handler"
	"codecreeo/internal/model"
	"codecreeo/internal/repository"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	app *fiber.App
	db  *gorm.DB
}

func (a *App) Register() {
	a.app.Get("/monitor", handler.Monitor())
	a.app.Post("/users", handler.NewUserHandler(repository.NewUserRepository(a.db)).CreateUser)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to automigrate tables: %v", err)
	}
	fmt.Println("Tables migrated successfully")

	app := fiber.New()

	app.Use(cors.New())
	app.Use(healthcheck.New())

	application := &App{app: app, db: db}
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
		log.Fatalf("app error: %s", err)
	}
}
