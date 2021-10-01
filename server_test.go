package fwork

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"testing"
	"time"
)

const defaultTestingPort = ":50000"

type Response struct {
	Message string
}

var sampleBaseReqContext = ReqContext{
	Req: httptest.NewRequest(http.MethodGet, "/hello-world", nil),
}
var sampleReqContextWithParams = ReqContext{
	Req: httptest.NewRequest(http.MethodGet, "/hello-world/1234", nil),
}
var sampleRoute = Route{
	Url:    "/hello-world",
	Method: http.MethodGet,
	Handler: func(c *ReqContext) {
		c.JsonReply(http.StatusOK, Response{Message: "Hello World"})
	},
	ComputedIdPattern: regexp.MustCompile("^GET /hello-world$"),
}
var sampleRouteWithUrlParam = Route{
	Url:               "/hello-world/:id",
	Method:            http.MethodGet,
	Handler:           func(context *ReqContext) {},
	ComputedIdPattern: regexp.MustCompile("^GET /hello-world/(?P<id>.*)$"),
}

func TestNewApi(t *testing.T) {
	got := NewApi()
	want := &Engine{
		Server: http.Server{Addr: defaultPort},
		Routes: []Route{},
	}
	want.Handler = want

	if !reflect.DeepEqual(got, want) {
		t.Errorf("NewApi() = %v, want %v", got, want)
	}
}

func TestEngine_findRoute_withSimplePath(t *testing.T) {
	want := sampleRoute
	e := &Engine{
		Routes: []Route{want},
	}

	got, err := e.findRoute(&sampleBaseReqContext)
	if err != nil {
		t.Errorf("findRoute() error = %v", err)
	}
	if !reflect.DeepEqual(got.ComputedIdPattern, want.ComputedIdPattern) {
		t.Errorf("findRoute() got = %v, want %v", got, want)
	}
}

func TestEngine_findRoute_withUrlParams(t *testing.T) {
	want := sampleRouteWithUrlParam
	e := &Engine{
		Routes: []Route{want},
	}

	got, err := e.findRoute(&sampleReqContextWithParams)
	if err != nil {
		t.Errorf("findRoute() error = %v", err)
	}
	if !reflect.DeepEqual(got.ComputedIdPattern, want.ComputedIdPattern) {
		t.Errorf("findRoute() got = %v, want %v", got, want)
	}
}

func TestEngine_findRoute_withNoRouteFound(t *testing.T) {
	e := &Engine{
		Routes: []Route{},
	}

	_, err := e.findRoute(&sampleBaseReqContext)
	if err != RouteNotFoundError {
		t.Errorf("findRoute() got error = %v, want %v", err, RouteNotFoundError)
	}
}

func TestEngine_Get(t *testing.T) {
	handler := func(c *ReqContext) {}
	e := &Engine{
		Routes: []Route{},
	}

	e.Get("/hello-world", handler)
	if len(e.Routes) == 0 {
		t.Errorf("Get() route did not register")
	}
}

func TestEngine_RunServer(t *testing.T) {
	e := &Engine{
		Server: http.Server{Addr: defaultTestingPort},
		Routes: []Route{},
	}

	go func(e *Engine) {
		time.Sleep(10 * time.Millisecond)
		if err := e.Server.Shutdown(context.Background()); err != nil {
			log.Panicf("unable to shutdown: [err: %v]", err)
		}
	}(e)

	if err := e.RunServer(); err != nil && err != http.ErrServerClosed {
		t.Errorf("RunServer() got error = %v", err)
	}
}

func TestEngine_ServeHTTP(t *testing.T) {
	e := &Engine{
		Server: http.Server{Addr: defaultTestingPort},
		Routes: []Route{sampleRoute},
	}
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(
		sampleRoute.Method,
		sampleRoute.Url,
		nil,
	)
	want := Response{Message: "Hello World"}

	e.ServeHTTP(res, req)
	var got Response
	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Errorf("ServeHTTP() got unexpected error %v", err)
	}

	if got != want {
		t.Errorf("ServeHTTP() got %v, want %v", got, want)
	}
	if res.Code != http.StatusOK {
		t.Errorf("ServeHTTP() got %v, want %v", res.Code, http.StatusOK)
	}
}

func TestEngine_ServeHTTP_NotFound(t *testing.T) {
	e := &Engine{
		Server: http.Server{Addr: defaultTestingPort},
		Routes: []Route{sampleRoute},
	}
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodGet,
		"/invalid-url",
		nil,
	)
	want := Void{}

	e.ServeHTTP(res, req)
	var got Void
	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Errorf("ServeHTTP() got unexpected error %v", err)
	}

	if got != want {
		t.Errorf("ServeHTTP() got %v, want %v", got, want)
	}
	if res.Code != http.StatusNotFound {
		t.Errorf("ServeHTTP() got %v, want %v", res.Code, http.StatusNotFound)
	}
}

func TestEngine_ServeHTTP_FindRouteFailsWithOtherError(t *testing.T) {
	e := &Engine{
		Server: http.Server{Addr: defaultTestingPort},
		Routes: []Route{sampleRoute},
	}
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodGet,
		"/invalid-url",
		nil,
	)
	want := Void{}

	e.ServeHTTP(res, req)
	var got Void
	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Errorf("ServeHTTP() got unexpected error %v", err)
	}

	if got != want {
		t.Errorf("ServeHTTP() got %v, want %v", got, want)
	}
	if res.Code != http.StatusNotFound {
		t.Errorf("ServeHTTP() got %v, want %v", res.Code, http.StatusNotFound)
	}
}
