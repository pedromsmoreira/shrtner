package data

import (
	"context"
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
)

// TODO: move to functions instead of interfaces
type Repository interface {
	List
	Create
	GetById
	Delete
}

type ReadDelete interface {
	Delete
	GetById
}

type List interface {
	List(ctx context.Context, page, pageSize int) ([]*domain.Url, error)
}

type Create interface {
	Create(ctx context.Context, url *domain.Url) (*domain.Url, error)
}

type GetById interface {
	GetById(ctx context.Context, id string) (*domain.Url, error)
}

type Delete interface {
	Delete(ctx context.Context, id string) error
}
