package handlers

// TODO: handle short url full address with DNS entry
// This will help and solve problems with localhost urls
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

func (he *HttpError) Error() string {
	return "code: " + he.Code + " message: " + he.Message
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
