package http

import (
	"context"
	"fmt"
	"github.com/pedromsmoreira/shrtener/internal/schema/db"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	cfg := configuration.NewSettings("../configuration/")

	err := db.CreateSchema(cfg.Database.SkipSchema, cfg.Database.Host, cfg.Database.DbName)

	if err != nil {
		log.Fatalf("error creating or updating the schema: %v", err)
	}

	cr, err := data.NewCockroachDbRepository(cfg.Database)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	defer cr.Close(context.Background())

	router := NewRouter(handlers.NewRestHandler(cr))

	server := NewServer(cfg, router)
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
