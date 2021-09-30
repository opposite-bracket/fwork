package fwork

// engine holds information and
// behaviour about the API server
type engine struct {
	Routes []*Route
}

// NewApi instantiates an api engine
func NewApi() *engine {
	return &engine{Routes: []*Route{}}
}
