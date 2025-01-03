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

func NewMissingRequiredParamsErrorResponse() *Response {
	return &Response{
		HttpStatus: http.StatusBadRequest,
		Data: ServerError{
			Parent:  ErrMissingRequiredParams,
			Message: ErrMissingRequiredParams.Error(),
		},
	}
}

func NewInternalServerErrorResponse(err error) *Response {
	return &Response{
		HttpStatus: http.StatusInternalServerError,
		Data: ServerError{
			Parent:  err,
			Message: "api error",
		},
	}
}
