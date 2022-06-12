package http

import (
	"context"
	"fmt"
	"github.com/pedromsmoreira/shrtener/internal/schema/db"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	settings := configuration.NewSettings("../configuration/")

	err := db.CreateSchema(settings.Database.SkipSchema, settings.Database.Host, settings.Database.DbName)

	if err != nil {
		log.Fatalf("error creating or updating the schema: %v", err)
	}

	crDb, err := data.NewCockroachDbRepository(settings.Database)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	defer crDb.Close(context.Background())

	router := NewRouter(settings.DNS, crDb)

	server := NewServer(settings, router)
	go func() {
		err := server.Start()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	code := m.Run()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	os.Exit(code)
}
