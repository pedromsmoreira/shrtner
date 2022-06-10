package data

import (
	"github.com/pedromsmoreira/shrtener/internal/shrtener/domain"
)

type urlData struct {
	Original       string
	Short          string
	ExpirationDate string
	DateCreated    string
}

func ToDataModel(m *domain.Url) *urlData {
	return &urlData{
		Original:       m.Original,
		Short:          m.Short,
		ExpirationDate: m.ExpirationDate,
		DateCreated:    m.DateCreated,
	}
}
