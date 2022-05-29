package data

import (
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
	"time"
)

type urlData struct {
	Original       string
	Short          string
	ExpirationDate time.Time
	DateCreated    time.Time
}

func ToDataModel(m *domain.Url) *urlData {
	return &urlData{
		Original:       m.Original,
		Short:          m.Short,
		ExpirationDate: m.ExpirationDate,
		DateCreated:    m.DateCreated,
	}
}
