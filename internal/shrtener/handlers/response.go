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

type HttpError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func validateInput(condition bool, code string, msg string, details map[string]interface{}) *HttpError {
	if condition {
		return &HttpError{
			Code:    code,
			Message: msg,
			Details: details,
		}
	}

	return nil
}
