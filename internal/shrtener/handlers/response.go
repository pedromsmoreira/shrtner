package handlers

type List struct {
	Data []*UrlMetadata `json:"data"`
}

type UrlMetadata struct {
	Original       string `json:"original"`
	Short          string `json:"short"`
	ExpirationDate string `json:"expiration_date"`
	DateCreated    string `json:"date_created"`
	DateModified   string `json:"date_modified"`
}
