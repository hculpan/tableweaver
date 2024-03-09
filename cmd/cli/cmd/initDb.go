package cmd

import (
	"errors"
	"os"

	"github.com/hculpan/tableweaver/pkg/db"
	"github.com/joho/godotenv"
)

func initDb() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("error loading .env file")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		return errors.New("environment variable DB_PATH not set")
	}

	err = db.InitDB(dbPath)
	if err != nil {
		return err
	}

	return nil
}

func closeDb() error {
	return db.DB.Close()
}
