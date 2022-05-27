package main

import (
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/handlers"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/http"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	settings := configuration.NewSettings()
	handlers := &handlers.Handler{}
	router := http.NewRouter(handlers)
	s := http.NewServer(settings, router)
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server is running! Ctrl-C to exit!")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-quit
	log.Println("Server is terminating!")
	err = s.Shutdown()
	log.Fatalf("Failed to shutdown server: %v", err)
}
