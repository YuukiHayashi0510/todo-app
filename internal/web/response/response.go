package response

import "net/http"

type Response struct {
	HttpStatus   int `json:"-"`
	TemplatePath *string
	RedirectPath *string
	Data         interface{}
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
