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

type ErrPerformingOperationInDb struct {
	Message      string
	WrappedError error
}

func NewErrPerformingOperationInDb(msg string, err error) *ErrPerformingOperationInDb {
	return &ErrPerformingOperationInDb{
		Message:      msg,
		WrappedError: err,
	}
}

func (ec *ErrPerformingOperationInDb) Error() string {
	err := fmt.Errorf("message: %s error: %w", ec.Message, ec.WrappedError)
	return err.Error()
}

type ErrEntryNotFoundInDb struct {
	Message      string
	WrappedError error
}

func NewEntryNotFoundInDbErr(id string, err error) *ErrEntryNotFoundInDb {
	return &ErrEntryNotFoundInDb{
		Message:      fmt.Sprintf("id %s not found in database", id),
		WrappedError: err,
	}
}

func (enf *ErrEntryNotFoundInDb) Error() string {
	err := fmt.Errorf("message: %s error: %w", enf.Message, enf.WrappedError)
	return err.Error()
}
