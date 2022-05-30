package domain

import (
	"errors"
	"time"
)

const defaultUrlActiveTimeInHours time.Duration = time.Hour * 24

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
	createdDate := time.Now().UTC()
	dc, err := time.Parse(time.RFC3339, createdDate.String())
	if err != nil {
		return nil, ErrConvertingCreatedDateToRFC3339
	}

	if expirationDate == "" {
		expirationDate = createdDate.Add(defaultUrlActiveTimeInHours).String()
	}

	ts, err := time.Parse(time.RFC3339, expirationDate)
	if err != nil {
		return nil, ErrConvertingExpirationDateToRFC3339
	}

	return &Url{
		Original:       original,
		DateCreated:    dc,
		ExpirationDate: ts,
	}, nil
}

func Shorten(u *Url) *Url {
	// shorten
	shortened := u.Original

	// start converting

	// assign
	u.Short = shortened
	return u
}
