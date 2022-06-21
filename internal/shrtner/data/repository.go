package data

import (
	"context"

	"github.com/pedromsmoreira/shrtener/internal/shrtner/domain"
)

type ReadKeyRepository interface {
	GetRedirect
}

type ReadWriteRepository interface {
	Create
	ReadDelete
	ReadRepository
	ReadKeyRepository
}

type ReadRepository interface {
	List
	GetById
}

type ReadDelete interface {
	ReadRepository
	Delete
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

type GetRedirect interface {
	GetRedirect(ctx context.Context, id string) (*domain.Redirect, error)
}
