package main

import (
	"fmt"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
)

func main() {
	settings := configuration.NewSettings()
	host := fmt.Sprintf("DB_HOST: %s", settings.Database.Host)
	fmt.Println(host)
}
