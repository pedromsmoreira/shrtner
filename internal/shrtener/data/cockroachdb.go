package data

import (
	"context"
	"fmt"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v4"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
)

type CockroachDbRepository struct {
	db *pgx.Conn
}

func NewCockroachDbRepository(dbCfg *configuration.Database) (*CockroachDbRepository, error) {
	connStr := fmt.Sprintf("postgresql://%s@%s/%s?sslmode=disable", dbCfg.Username, dbCfg.Host, dbCfg.DbName)
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

	// TODO: Apply check according to db error
	if err != nil {
		// TODO: add custom error
		return nil, NewErrPerformingOperationInDb("error creating data in db", err)
	}

	return url, nil
}

func (r *CockroachDbRepository) List(ctx context.Context, page, size int) ([]*domain.Url, error) {
	offset := page * size

	query := `SELECT short_url, original_url, created_date, expiration_date FROM urls ORDER BY created_date LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(ctx, query, size, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	urls := make([]*domain.Url, 0)
	for rows.Next() {
		var sUrl, origUrl, expirationDate, createdDate string

		err := rows.Scan(&sUrl, &origUrl, &createdDate, &expirationDate)
		if err != nil {
			return nil, NewErrPerformingOperationInDb("error reading data from db", err)
		}

		urls = append(urls, &domain.Url{
			Original:       origUrl,
			Short:          sUrl,
			ExpirationDate: expirationDate,
			DateCreated:    createdDate,
		})
	}

	return urls, nil
}

func (r *CockroachDbRepository) GetById(ctx context.Context, id string) (*domain.Url, error) {
	query := `SELECT short_url, original_url, created_date, expiration_date FROM urls WHERE short_url = $1`

	var sUrl, origUrl, expirationDate, createdDate string
	err := r.db.QueryRow(ctx, query, id).Scan(&sUrl, &origUrl, &createdDate, &expirationDate)
	switch err {
	case nil:
		return &domain.Url{
			Original:       origUrl,
			Short:          sUrl,
			ExpirationDate: expirationDate,
			DateCreated:    createdDate,
		}, nil
	case pgx.ErrNoRows:
		return nil, NewEntryNotFoundInDbErr(id, err)
	default:
		return nil, NewEntryNotFoundInDbErr(id, err)
	}
}

func (r *CockroachDbRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, "DELETE FROM urls WHERE short_url=$1", id)
	return err
}
