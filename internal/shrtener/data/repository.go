package data

import (
	"context"
	"fmt"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v4"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
	"time"
)

type CockroachDbRepository struct {
	db *pgx.Conn
}

func NewCockroachDbRepository(dbCfg *configuration.Database) (*CockroachDbRepository, error) {
	connStr := fmt.Sprintf("postgresql://root@%s/%s?sslmode=disable", dbCfg.Host, dbCfg.DbName)
	config, err := pgx.ParseConfig(connStr)
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, NewErrConnectingToDb("error connecting to cockroach database", err)
	}

	return &CockroachDbRepository{
		db: conn,
	}, nil
}

func (r *CockroachDbRepository) Close(ctx context.Context) {
	err := r.db.Close(ctx)
	if err != nil {
		panic(err)
	}
}

func (r *CockroachDbRepository) Create(ctx context.Context, url *domain.Url) (*domain.Url, error) {
	err := crdbpgx.ExecuteTx(ctx, r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx,
			"INSERT INTO urls(short_url, original_url, created_date, expiration_date) VALUES ($1, $2, $3, $4)",
			url.Short, url.Original, url.DateCreated, url.ExpirationDate)

		return err
	})

	if err != nil {
		// TODO: add custom error
		return nil, NewErrPerformingOperationInDb("error creating data in db", err)
	}

	return url, nil
}

func (r *CockroachDbRepository) List(ctx context.Context) ([]*domain.Url, error) {
	rows, err := r.db.Query(ctx, "SELECT short_url, original_url, created_date, expiration_date FROM urls")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	urls := make([]*domain.Url, 0)
	if !rows.Next() {
		return urls, nil
	}

	for rows.Next() {
		var sUrl, origUrl string
		var expirationDate, createdDate time.Time

		err := rows.Scan(&sUrl, &origUrl, createdDate, expirationDate)
		if err != nil {
			// TODO: add custom error
			return nil, NewErrPerformingOperationInDb("error reading data from db", err)
		}

		urls = append(urls, &domain.Url{
			Original:       sUrl,
			Short:          origUrl,
			ExpirationDate: expirationDate,
			DateCreated:    createdDate,
		})
	}

	return urls, nil
}
