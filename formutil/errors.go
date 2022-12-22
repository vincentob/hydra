package formutil

import "net/http"

type FormError struct {
	Status int
	Err    string
}

func (err FormError) Error() string {
	return err.Err
}

type StatusOKError FormError

func (err StatusOKError) Error() string {
	return err.Err
}

func NewFormError(status int, msg string) FormError {
	return FormError{Status: status, Err: msg}
}

// 400
func NewBadRequestError() FormError {
	return FormError{Status: http.StatusBadRequest, Err: "BadRequest"}
}

// 401
func NewUnauthorizedError() FormError {
	return FormError{Status: http.StatusUnauthorized, Err: "Unauthorized"}
}

// 403
func NewForbiddenError() FormError {
	return FormError{Status: http.StatusForbidden, Err: "Forbidden"}
}

// 404
func NewNotFoundError() FormError {
	return FormError{Status: http.StatusNotFound, Err: "Not Found"}
}

// 500
func NewInternalError() FormError {
	return FormError{Status: http.StatusInternalServerError, Err: "InternalServerError"}
}

// Specify error, return 200 in http header, but non zero status in body.
func NewStatusOKError() StatusOKError {
	return StatusOKError{Status: http.StatusOK, Err: ""}
}
