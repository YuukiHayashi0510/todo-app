package response

type Response struct {
	HttpStatus   int `json:"-"`
	TemplatePath *string
	RedirectPath *string
	Data         interface{}
}
