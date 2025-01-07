package common

type SearchStatus string

const (
	StatusAll      SearchStatus = "all"
	StatusActive   SearchStatus = "active"
	StatusInActive SearchStatus = "in_active"
)

func (s SearchStatus) Is() bool {
	return s.IsAll() || s.IsActive() || s.IsInActive()
}

func (s SearchStatus) IsAll() bool {
	return s == StatusAll
}

func (s SearchStatus) IsActive() bool {
	return s == StatusActive
}

func (s SearchStatus) IsInActive() bool {
	return s == StatusInActive
}
