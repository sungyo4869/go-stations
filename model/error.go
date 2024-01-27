package model

type ErrNotFound struct{
	Resource string
}

func (e ErrNotFound) Error() string {
	if e.Resource == "" {
		return "Not found"
	}
	return "Not found: " + e.Resource
}