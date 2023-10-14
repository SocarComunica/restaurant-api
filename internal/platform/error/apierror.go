package platformerrors

import (
	"fmt"
	"net/http"
)

type CauseList []interface{}

type APIError struct {
	ErrorMessage string    `json:"message"`
	ErrorCode    string    `json:"error"`
	ErrorStatus  int       `json:"status"`
	ErrorCause   CauseList `json:"cause"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Message: %s;Error code: %s, Status: %d; Cause: %v", e.ErrorMessage, e.ErrorCode, e.ErrorStatus, e.ErrorCause)
}

const (
	internalServerError string = "internal_server_error"
	notFoundError       string = "resource_not_found"
	badRequestError     string = "bad_request"
)

func (e APIError) Status() int {
	return e.ErrorStatus
}

func NewBadRequestAPIError(message string, err error) APIError {
	cause := CauseList{}
	if err != nil {
		cause = append(cause, err.Error())
	}
	return APIError{message, badRequestError, http.StatusBadRequest, cause}
}

func NewNotFoundAPIError(message string, err error) APIError {
	cause := CauseList{}
	if err != nil {
		cause = append(cause, err.Error())
	}
	return APIError{message, notFoundError, http.StatusNotFound, cause}
}

func NewInternalServerAPIError(message string, err error) APIError {
	cause := CauseList{}
	if err != nil {
		cause = append(cause, err.Error())
	}
	return APIError{message, internalServerError, http.StatusInternalServerError, cause}
}
