package main

import (
	"context"
	"fmt"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/data"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/http"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfgFolder := "./internal/shrtener/configuration/"
	settings := configuration.NewSettings(cfgFolder)

	db, err := data.NewCockroachDbRepository(settings.Database)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	router := http.NewRouter(handlers.NewRestHandler(db))

	s := http.NewServer(settings, router)

	go func() {
		err = s.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server is running! Ctrl-C to exit!")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = s.Shutdown(ctx)
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}
