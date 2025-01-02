package response

import "fmt"

type ServerError struct {
	Parent  error  `json:"-"`
	Message string `json:"message"`
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("message: %s, parent: %v", e.Message, e.Parent)
}

func (e *ServerError) Unwrap() error {
	return e.Parent
}
