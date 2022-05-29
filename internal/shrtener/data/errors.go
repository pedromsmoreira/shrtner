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
