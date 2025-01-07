package organization

import "errors"

var (
	ErrOrganizationNotFound          = errors.New("organization not found")
	ErrOrganizationHasAlreadyDeleted = errors.New("organization has already deleted")
	ErrOrganizationIsNotDeleted      = errors.New("organization isn't deleted")
)
