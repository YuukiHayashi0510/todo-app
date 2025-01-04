package staff

import "errors"

var (
	ErrStaffNotFound          = errors.New("staff not found")
	ErrStaffHasAlreadyDeleted = errors.New("staff has already deleted")
	ErrStaffIsNotDeleted      = errors.New("staff isn't deleted")
)
