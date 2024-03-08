package database

import (
	"codecreeo/internal/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	db *gorm.DB
}

func NewDbConnection() *DbConnection {
	dbUrl := os.Getenv("DB_CONNECTION_STRING")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &DbConnection{
		db: db,
	}
}

func (conn *DbConnection) CloseDbConnection() {
	sqlDB, err := conn.db.DB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get SQL database: %v\n", err)
		return
	}
	sqlDB.Close()
}

func (conn *DbConnection) AutoMigrate() error {
	if err := conn.db.AutoMigrate(&model.User{}, &model.QRCode{}); err != nil {
		return err
	}
	return nil
}
