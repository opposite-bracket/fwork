package fwork

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"
)

func TestGenerateUrlPattern(t *testing.T) {
	type args struct {
		method string
		url    string
	}
	tests := []struct {
		name string
		args args
		want *regexp.Regexp
	}{
		{
			"No params",
			args{method: http.MethodGet, url: "/hello-world"},
			regexp.MustCompile("^GET /hello-world$"),
		},
		{
			"UrlParams with 1 alphabetic var",
			args{method: http.MethodGet, url: "/hello-world/:id"},
			regexp.MustCompile("^GET /hello-world/(?P<id>.*)$"),
		},
		{
			"UrlParams with 1 alphanumeric var",
			args{method: http.MethodGet, url: "/hello-world/:id123"},
			regexp.MustCompile("^GET /hello-world/(?P<id123>.*)$"),
		},
		{
			"UrlParams with 1 var with all characters",
			args{method: http.MethodGet, url: "/hello-world/:id123_AB"},
			regexp.MustCompile("^GET /hello-world/(?P<id123_AB>.*)$"),
		},
		{
			"UrlParams with 1 alphabetic var and a prefix",
			args{method: http.MethodGet, url: "/hello-world/hello-world-:id"},
			regexp.MustCompile("^GET /hello-world/hello-world-(?P<id>.*)$"),
		},
		{
			"UrlParams with 1 alphanumeric var",
			args{method: http.MethodGet, url: "/hello-world/hello-world-:id123"},
			regexp.MustCompile("^GET /hello-world/hello-world-(?P<id123>.*)$"),
		},
		{
			"UrlParams with 1 var with all characters",
			args{method: http.MethodGet, url: "/hello-world/hello-world-:id123_AB"},
			regexp.MustCompile("^GET /hello-world/hello-world-(?P<id123_AB>.*)$"),
		},

		{
			"UrlParams with multiple alphabetic vars",
			args{method: http.MethodGet, url: "/hello-world/:first/:second"},
			regexp.MustCompile("^GET /hello-world/(?P<first>.*)/(?P<second>.*)$"),
		},
		{
			"UrlParams with multiple alphanumeric vars",
			args{method: http.MethodGet, url: "/hello-world/:id1/:id2"},
			regexp.MustCompile("^GET /hello-world/(?P<id1>.*)/(?P<id2>.*)$"),
		},
		{
			"UrlParams with multiple vars with all characters",
			args{method: http.MethodGet, url: "/hello-world/:id1_A/:id2_B"},
			regexp.MustCompile("^GET /hello-world/(?P<id1_A>.*)/(?P<id2_B>.*)$"),
		},
		{
			"UrlParams with multiple spread out vars with all characters",
			args{method: http.MethodGet, url: "/hello/:id1_A/world/:id2_B"},
			regexp.MustCompile("^GET /hello/(?P<id1_A>.*)/world/(?P<id2_B>.*)$"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := GenerateUrlPatternMatcher(tt.args.method, tt.args.url); err != nil {
				t.Errorf("failed with error %v", err)
			} else if got.String() != tt.want.String() {
				t.Errorf("GenerateUrlPatternMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeRouteId(t *testing.T) {
	method := http.MethodGet
	url := "/hello-world"
	got := ComputeRouteId(method, url)
	want := fmt.Sprintf("%s %s", method, url)
	if got != want {
		t.Errorf("ComputeRouteId() = %v, want %v", got, want)
	}
}
