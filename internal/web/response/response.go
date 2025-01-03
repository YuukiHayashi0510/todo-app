package response

import (
	"errors"
	"net/http"
)

var (
	ErrMissingRequiredParams = errors.New("missing required params")
)

type Response struct {
	HttpStatus   int `json:"-"`
	TemplatePath *string
	RedirectPath *string
	Data         interface{}
}

func NewMissingRequiredParamsError() *Response {
	return &Response{
		HttpStatus: http.StatusBadRequest,
		Data: ServerError{
			Parent:  ErrMissingRequiredParams,
			Message: ErrMissingRequiredParams.Error(),
		},
	}
}

func NewBadRequestError(err error) *Response {
	return &Response{
		HttpStatus: http.StatusBadRequest,
		Data: ServerError{
			Parent:  err,
			Message: "invalid request",
		},
	}
}

func NewNotFoundError(err error) *Response {
	return &Response{
		HttpStatus: http.StatusNotFound,
		Data: ServerError{
			Parent:  err,
			Message: "resource not found",
		},
	}
}

func NewInternalServerError(err error) *Response {
	return &Response{
		HttpStatus: http.StatusInternalServerError,
		Data: ServerError{
			Parent:  err,
			Message: "server error",
		},
	}
}
