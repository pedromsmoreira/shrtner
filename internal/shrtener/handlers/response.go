package handlers

type UrlMetadata struct {
	Original       string `json:"original"`
	Short          string `json:"short,omitempty"`
	ExpirationDate string `json:"expiration_date,omitempty"`
	DateCreated    string `json:"date_created,omitempty"`
}
