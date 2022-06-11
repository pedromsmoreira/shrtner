package data

import (
	"context"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
)

type Repository interface {
	List
	Create
}

type List interface {
	List(ctx context.Context) ([]*domain.Url, error)
}

type Create interface {
	Create(ctx context.Context, url *domain.Url) (*domain.Url, error)
}
