package main

import (
	"fmt"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	fmt.Printf("PORT: %v\n\n", port)
	dbHost := os.Getenv("DB_HOST")
	fmt.Printf("DB_HOST: %v\n\n", dbHost)
	dbName := os.Getenv("DB_NAME")
	fmt.Printf("DB_NAME: %v\n\n", dbName)

}
