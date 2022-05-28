package http

import (
	"github.com/pedromsmoreira/shrtener/internal/schema/db"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	cfg := &configuration.Settings{
		Database: configuration.Database{
			Host:     "localhost:26257",
			Username: "admin",
			Password: "password",
		},
		Server: configuration.Server{
			Port: 5000,
			Host: "localhost",
		},
	}

	skipSchema, err := strconv.ParseBool(os.Getenv("SKIP_SCHEMA"))
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	err = db.CreateSchema(skipSchema, host, dbName)

	if err != nil {
		log.Fatalf("error creating or updating the schema: %v", err)
	}

	server := NewServer(cfg, nil)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	code := m.Run()

	if err := server.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	os.Exit(code)
}
