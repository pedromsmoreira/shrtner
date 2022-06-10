package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v4"
	"log"
)

func CreateSchema(skipSchema bool, host string, dbName string) error {
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
						CREATE TABLE IF NOT EXISTS urls (
							created_date STRING,
							expiration_date STRING,
							original_url STRING,
							short_url STRING,
							PRIMARY KEY (short_url)
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
