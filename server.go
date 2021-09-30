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
	c := ReqContext{
		Req: r,
		Res: w,
	}
	if route, err := e.findRoute(r); err != nil && err == RouteNotFoundError {
		c.JsonReply(http.StatusNotFound, Void{})
	} else if err != nil {
		c.JsonReply(http.StatusInternalServerError, Void{})
	} else {
		route.Handler(&c)
	}

}

// Get registers an http request with GET method
func (e *engine) Get(url string, handler RouteHandler) {
	e.Routes = append(e.Routes, Route{
		Url:     url,
		Method:  http.MethodGet,
		Handler: handler,
		ComputedId: ComputeRouteId(http.MethodGet, url),
	})
}

func (e *engine) RunServer() {
	log.Printf("Running on addr %s", defaultPort)
	http.ListenAndServe(defaultPort, e)
}

func (e *engine) findRoute(req *http.Request) (*Route, error) {
	expectedComputedId := ComputeRouteId(req.Method, req.URL.Path)
	for _, route := range e.Routes {
		if route.ComputedId == expectedComputedId {
			return &route, nil
		}
	}

	return nil, RouteNotFoundError
}

// NewApi instantiates an api engine
func NewApi() *engine {
	return &engine{Routes: []Route{}}
}
