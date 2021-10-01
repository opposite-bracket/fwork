package fwork

import "regexp"

// Route contains information
// about routes that the api Engine
// is capable of handling
type Route struct {
	Url               string
	Method            string
	Handler           RouteHandler
	ComputedIdPattern *regexp.Regexp
}

// RouteHandler supports http requests
// represented by routes
type RouteHandler func(*ReqContext)
