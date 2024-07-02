package rest

// Rest represents an implementation of a REST service
type Rest struct{}

// New creates a new rest service instance
func New() (*Rest, error) {
	return &Rest{}, nil
}
