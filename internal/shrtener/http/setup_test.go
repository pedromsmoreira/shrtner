package http

import (
	"github.com/pedromsmoreira/shrtener/internal/schema/db"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	cfg := configuration.NewSettings("../configuration/")

	err := db.CreateSchema(cfg.Database.SkipSchema, cfg.Database.Host, cfg.Database.DbName)

	if err != nil {
		log.Fatalf("error creating or updating the schema: %v", err)
	}
	h := &handlers.Handler{}
	router := NewRouter(h)

	server := NewServer(cfg, router)
	if err := server.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

	code := m.Run()

	if err := server.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	os.Exit(code)
}
