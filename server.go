package fwork

import (
	"log"
	"net/http"
)

const defaultPort = ":5000"

var (
	// RouteNotFoundError is thrown when incoming request
	// did not match a request
	RouteNotFoundError = ApiError{
		Status:  http.StatusNotFound,
		Code:    "01",
		Message: "route not found",
	}
)

type Void struct{}
type ErrorResponse struct {
	Code    int
	Message string
}

// Engine holds information and
// behaviour about the API server
type Engine struct {
	http.Server
	Routes []Route
}

// ServeHTTP complies with Handler interface to
// be able to determine which route needs to be
// used to handle request
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewReqContext(w, r)
	if route, err := e.findRoute(c); err != nil && err == RouteNotFoundError {
		c.JsonReply(http.StatusNotFound, Void{})
	} else {
		route.Handler(c)
	}
}

// Get registers an http request with GET method
func (e *Engine) Get(url string, handler RouteHandler) {

	pattern, _ := GenerateUrlPatternMatcher(http.MethodGet, url)

	e.Routes = append(e.Routes, Route{
		Url:               url,
		Method:            http.MethodGet,
		Handler:           handler,
		ComputedIdPattern: pattern,
	})
}

// findRoute figures out if the incoming request is supported.
// throws the following errors when evaluating if a route
// matches the requested:
// 		RouteNotFoundError if route is not found when
// 		comparing the ComputedIdPattern of the route
//		InvalidRouteUrlError if the pattern is invalid
func (e *Engine) findRoute(c *ReqContext) (*Route, error) {
	expectedComputedId := ComputeRouteId(c.Req.Method, c.Req.URL.Path)
	for _, route := range e.Routes {

		var myExp = route.ComputedIdPattern
		match := myExp.FindStringSubmatch(expectedComputedId)
		matchCount := len(match)
		switch {
		case matchCount == 1:
			return &route, nil
		case matchCount > 1:
			result := make(map[string]string, matchCount)
			for i, name := range myExp.SubexpNames() {
				if i != 0 && name != "" {
					result[name] = match[i]
				}
			}
			c.Params.Url = result
			return &route, nil
		}

	}

	return nil, RouteNotFoundError
}

func (e *Engine) RunServer() error {
	log.Printf("Running on addr %s", defaultPort)
	return e.ListenAndServe()
}

// NewApi instantiates an api Engine
func NewApi() *Engine {
	e := &Engine{
		Server: http.Server{Addr: defaultPort},
		Routes: []Route{},
	}
	e.Handler = e
	return e
}
