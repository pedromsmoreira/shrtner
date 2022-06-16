package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/http"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/logger"
)

func main() {
	cfgFolder := "./internal/shrtener/configuration/"
	settings := configuration.NewSettings(cfgFolder)
	logger.ConfigureLogrus(settings)

	crDB, err := data.NewCockroachDbRepository(settings.Database)
	if err != nil {
		logger.Error("error initializing db connection", map[string]interface{}{"error": err})
		os.Exit(1)
	}
	defer crDB.Close(context.Background())

	router := http.NewRouter(settings.DNS, crDB)

	s := http.NewServer(settings, router)

	go func() {
		err = s.Start()
		if err != nil {
			logger.Error("error starting the server", map[string]interface{}{"error": err})
		}
	}()

	logger.Info("Server is running! Ctrl-C to exit!")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = s.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown: ", map[string]interface{}{"error": err})
	}
	logger.Info("Server exiting")
}
