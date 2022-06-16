package main

import (
	"os"
	"strconv"

	"github.com/pedromsmoreira/shrtener/internal/schema/db"
)

func main() {
	skipSchema, err := strconv.ParseBool(os.Getenv("SKIP_SCHEMA"))
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	err = db.CreateSchema(skipSchema, host, dbName)
	if err != nil {
		panic(err)
	}
}
