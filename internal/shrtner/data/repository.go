package data

import (
	"context"

	"github.com/pedromsmoreira/shrtener/internal/shrtner/domain/url"
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
	List(ctx context.Context, page, pageSize int) ([]*url.Url, error)
}

type Create interface {
	Create(ctx context.Context, url *url.Url) (*url.Url, error)
}

type GetById interface {
	GetById(ctx context.Context, id string) (*url.Url, error)
}

type Delete interface {
	Delete(ctx context.Context, id string) error
}

type GetRedirect interface {
	GetRedirect(ctx context.Context, id string) (*url.Redirect, error)
}
