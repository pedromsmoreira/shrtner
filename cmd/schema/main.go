package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
	"strconv"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
)

func main() {
	skipSchema, err := strconv.ParseBool(os.Getenv("SKIP_SCHEMA"))
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	err = createSchema(skipSchema, host, dbName)
	if err != nil {
		panic(err)
	}
}

func createSchema(skipSchema bool, host string, dbName string) error {
	if skipSchema {
		log.Println("schema >>> schema creation skipped...")
		return nil
	}
	connStr := fmt.Sprintf("postgresql://root@%s/defaultdb?sslmode=disable", host)
	config, err := pgx.ParseConfig(connStr)
	config.Database = "defaultdb"
	if err != nil {
		return errors.New(fmt.Sprintf("error schema configuration: %v", err))
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return errors.New(fmt.Sprintf("error connecting to the database: %v", err))
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(conn, context.Background())

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		log.Println("schema >>> creating Database...")
		if _, err := tx.Exec(context.Background(),
			fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName)); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return errors.New(fmt.Sprintf("error creating database: %v", err))
	}

	log.Printf("schema >>> changing to %s database...", dbName)
	config.Database = dbName
	conn, err = pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return errors.New(fmt.Sprintf("error connecting to turbotodo database: %v", err))
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(conn, context.Background())

	log.Printf("schema >>> creating %s tables...", dbName)
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		log.Println("schema >>> creating urls table...")
		_, err := tx.Exec(context.Background(), `
						CREATE TABLE urls (
							tenant STRING(36),
							id STRING(36),
							created_date TIMESTAMP,
							modified_date TIMESTAMP,
							original_url STRING,
							short_url STRING,
							PRIMARY KEY (tenant, id)
						);`)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return errors.New(fmt.Sprintf("error creating todos table: %v", err))
	}

	return nil
}
