package handlers

type InternalServerError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewInternalServerError(message string) error {
	return &InternalServerError{
		Code:    "500",
		Message: message,
	}
}

func (ise *InternalServerError) Error() string {
	return "status_code: " + ise.Code + " message: " + ise.Message
}

type BadRequestError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func NewBadRequestError(message string, details interface{}) error {
	return &BadRequestError{
		Code:    "400",
		Message: message,
		Details: details,
	}
}

func (bde *BadRequestError) Error() string {
	return "status_code: " + bde.Code + " message: " + bde.Message
}

type ConflictError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewConflictError(message string) error {
	return &ConflictError{
		Code:    "409",
		Message: message,
	}
}

func (ce *ConflictError) Error() string {
	return "status_code: " + ce.Code + " message: " + ce.Message
}
