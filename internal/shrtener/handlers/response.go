package handlers

type List struct {
	Data []*UrlMetadata `json:"data"`
}

type UrlMetadata struct {
	Original       string `json:"original"`
	Short          string `json:"short,omitempty"`
	ExpirationDate string `json:"expiration_date,omitempty"`
	DateCreated    string `json:"date_created,omitempty"`
}

type Error struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}
