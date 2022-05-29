package domain

import (
	"errors"
	"time"
)

var (
	ErrConvertingExpirationDateToRFC3339 = errors.New("could not convert expiration date to RFC3339")
	ErrConvertingCreatedDateToRFC3339    = errors.New("could not convert created date to RFC3339")
)

type Url struct {
	Original       string
	Short          string
	ExpirationDate time.Time
	DateCreated    time.Time
}

func NewUrl(original string, expirationDate string) (*Url, error) {
	ts, err := time.Parse(time.RFC3339, expirationDate)
	if err != nil {
		return nil, ErrConvertingExpirationDateToRFC3339
	}

	dc, err := time.Parse(time.RFC3339, time.Now().UTC().String())
	if err != nil {
		return nil, ErrConvertingCreatedDateToRFC3339
	}

	return &Url{
		DateCreated:    dc,
		ExpirationDate: ts,
	}, nil
}
