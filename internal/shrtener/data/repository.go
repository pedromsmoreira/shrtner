package data

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
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

func (r *CockroachDbRepository) List() ([]*domain.Url, error) {
	return make([]*domain.Url, 0), nil
}
