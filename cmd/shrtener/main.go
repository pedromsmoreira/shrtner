package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/http"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/logger"
)

func main() {
	cfgFolder := "./internal/shrtener/configuration/"
	settings := configuration.NewSettings(cfgFolder)
	log := logrus.StandardLogger()
	log.SetLevel(logger.WithLogLevel(settings.Logging))
	log.SetFormatter(logger.WithFormatter(settings.Logging.Format))
	log.SetOutput(logger.WithOutput(settings.Logging))

	crDB, err := data.NewCockroachDbRepository(settings.Database)
	if err != nil {
		logrus.WithField("error", err.Error()).Error("error initializing db connection")
		os.Exit(1)
	}
	defer crDB.Close(context.Background())

	router := http.NewRouter(settings.DNS, crDB)

	s := http.NewServer(settings, router)

	go func() {
		err = s.Start()
		if err != nil {
			logrus.WithField("error", err.Error()).Error("error starting the server")
		}
	}()

	logrus.Info("Server is running! Ctrl-C to exit!")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-quit
	logrus.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = s.Shutdown(ctx); err != nil {
		logrus.WithField("error", err.Error()).Error("server forced to shutdown")
	}
	logrus.Info("Server exiting")
}
