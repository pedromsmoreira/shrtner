package data

import "fmt"

type ErrConnectingToDb struct {
	Message      string
	WrappedError error
}

func NewErrConnectingToDb(msg string, err error) *ErrConnectingToDb {
	return &ErrConnectingToDb{
		Message:      msg,
		WrappedError: err,
	}
}

func (ec *ErrConnectingToDb) Error() string {
	err := fmt.Errorf("message: %s error: %w", ec.Message, ec.WrappedError)
	return err.Error()
}

const duplicateKeyCode = "23505"

type ErrPerformingOperationInDb struct {
	Message string
}

func NewErrPerformingOperationInDb(code string, msg string) *ErrPerformingOperationInDb {
	if code == duplicateKeyCode {
		msg = "key already exists"
	}

	return &ErrPerformingOperationInDb{
		Message: msg,
	}
}

func (ec *ErrPerformingOperationInDb) Error() string {
	return ec.Message
}

type ErrEntryNotFoundInDb struct {
	Message string
	Detail  string
}

func NewEntryNotFoundInDbErr(id string, detail string) *ErrEntryNotFoundInDb {
	return &ErrEntryNotFoundInDb{
		Message: fmt.Sprintf("id %s not found in database", id),
		Detail:  detail,
	}
}

func (enf *ErrEntryNotFoundInDb) Error() string {
	return fmt.Sprintf("message: %s details: %s", enf.Message, enf.Detail)
}
