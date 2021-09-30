package fwork

import (
	"errors"
	"log"
	"net/http"
)

const defaultPort = ":5000"

var (
	// RouteNotFoundError is thrown when incoming request
	// did not match a request
	RouteNotFoundError = errors.New("route not found")
	InvalidRouteUrlError = errors.New("url route is not valid")
)

type Void struct {}

// engine holds information and
// behaviour about the API server
type engine struct {
	Routes []Route
}

// ServeHTTP complies with Handler interface to
// be able to determine which route needs to be
// used to handle request
func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewReqContext(w, r)
	if route, err := e.findRoute(c); err != nil && err == RouteNotFoundError {
		c.JsonReply(http.StatusNotFound, Void{})
	} else if err != nil {
		c.JsonReply(http.StatusInternalServerError, Void{})
	} else {
		route.Handler(c)
	}

}

// Get registers an http request with GET method
func (e *engine) Get(url string, handler RouteHandler) {

	pattern, _ := GenerateUrlPatternMatcher(http.MethodGet, url)

	e.Routes = append(e.Routes, Route{
		Url:               url,
		Method:            http.MethodGet,
		Handler:           handler,
		ComputedIdPattern: pattern,
	})
}

func (e *engine) RunServer() {
	log.Printf("Running on addr %s", defaultPort)
	http.ListenAndServe(defaultPort, e)
}

// findRoute figures out if the incoming request is supported.
// throws the following errors when evaluating if a route
// matches the requested:
// 		RouteNotFoundError if route is not found when
// 		comparing the ComputedIdPattern of the route
//		InvalidRouteUrlError if the pattern is invalid
func (e *engine) findRoute(c *ReqContext) (*Route, error) {
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

// NewApi instantiates an api engine
func NewApi() *engine {
	return &engine{Routes: []Route{}}
}
