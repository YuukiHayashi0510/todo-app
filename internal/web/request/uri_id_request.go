package request

type PathParams struct {
	ID int64 `uri:"id" binding:"number,gt=0"`
}
