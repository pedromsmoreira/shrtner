package handlers

type InternalServerError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewInternalServerError(message string) error {
	return &InternalServerError{
		Code:    "100001",
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
		Code:    "100002",
		Message: message,
		Details: details,
	}
}

func NewBadRequestErrorWithoutDetails(message string) error {
	return &BadRequestError{
		Code:    "100003",
		Message: message,
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
		Code:    "100004",
		Message: message,
	}
}

func (ce *ConflictError) Error() string {
	return "status_code: " + ce.Code + " message: " + ce.Message
}

type NotFoundError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewNotFoundError(id string) error {
	return &NotFoundError{
		Code:    "100005",
		Message: "path: " + id + " not found",
	}
}

func (nfe *NotFoundError) Error() string {
	return "status_code: " + nfe.Code + " message: " + nfe.Message
}

type ExpiredLinkError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewExpiredLinkError(id string, expirationDate string) error {
	return &ExpiredLinkError{
		Code:    "100006",
		Message: "id: " + id + " has expired at " + expirationDate,
	}
}

func (ele *ExpiredLinkError) Error() string {
	return "status_code: " + ele.Code + " message: " + ele.Message
}
