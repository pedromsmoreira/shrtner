package handlers

type List struct {
	Data []*UrlMetadata `json:"data"`
}

type UrlMetadata struct {
	Id           string   `json:"id"`
	Type         string   `json:"type"`
	Data         *UrlData `json:"data"`
	DateCreated  string   `json:"date_created"`
	DateModified string   `json:"date_modified"`
	Version      int64    `json:"version"`
}

type UrlData struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}
