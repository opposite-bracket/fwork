package fwork

// Route contains information
// about routes that the api engine
// is capable of handling
type Route struct {
	Url    string
	Method string
	Handler Handler
}

// Handler supports http requests
// represented by routes
type Handler func(*ReqContext)
