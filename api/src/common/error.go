package common

import (
	"fmt"
	"runtime"

	"github.com/lib/pq"
)

// Error ..
type Error struct {
	Code          int
	Message       string
	StackTrace    string
	DatabaseError *pq.Error
}

func (err Error) Error() (res string) {
	res += fmt.Sprintf("Message: %s\n", err.Message)
	res += "---------------------------------------\n"
	res += fmt.Sprintf("Stacktrace: %s\n", err.StackTrace)
	res += "---------------------------------------\n"
	if err.DatabaseError != nil {
		res += "Database error."
		res += fmt.Sprintf("Code: %s\n", err.DatabaseError.Code)
		res += fmt.Sprintf("Details: %s\n", err.DatabaseError.Detail)
	}
	return
}

func (err *Error) writeStackTrace() {
	trace := make([]byte, 4096)
	runtime.Stack(trace, false)
	err.StackTrace = string(trace)
}

// GetMessageForUser ..
func (err *Error) GetMessageForUser() (message string) {
	message += fmt.Sprintf("Message: %s\n", err.Message)
	if err.DatabaseError != nil {
		message += fmt.Sprintf("Details: %s\n", err.DatabaseError.Detail)
	}
	return
}

// CreateErrorWithMessage ..
func CreateErrorWithMessage(message string) (err Error) {
	err.writeStackTrace()
	err.Message = message
	return
}

// CreateError ..
func CreateError(err error) (res Error) {
	fullErr, ok := err.(*pq.Error)

	if ok {
		res.writeStackTrace()
		res.Message = err.Error()
		res.DatabaseError = fullErr
	} else {
		copyErr, ok := err.(Error)
		if ok {
			res.Message = copyErr.Message
			res.StackTrace = copyErr.StackTrace
			res.DatabaseError = copyErr.DatabaseError
		} else {
			res.writeStackTrace()
			res.Message = err.Error()
		}

	}

	return
}
