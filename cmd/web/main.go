package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/hculpan/tableweaver/cmd/web/handler"
	"github.com/hculpan/tableweaver/pkg/db"
	"github.com/hculpan/tableweaver/pkg/session"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the port from the environment variables
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set.")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH environment variable not set.")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable not set.")
	}

	session.InitializeSessionManagement(secretKey)

	err = db.InitDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	mux := http.NewServeMux()

	handler.Routes(mux)

	slog.Info("server listening on port " + port)
	http.ListenAndServe(":"+port, mux)
}
