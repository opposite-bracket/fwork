package fwork

import (
	"fmt"
	"regexp"
)

var sampleRegexp = regexp.MustCompile(`:([a-zA-Z0-9_]+)`)

// ComputeRouteIdPattern generates a pattern for evaluating
// registered routes against incoming ones
func ComputeRouteIdPattern(method string, url string) string {
	return fmt.Sprintf("^%s %s$", method, url)
}

// ComputeRouteId generates a format for evaluating
// registered routes against incoming ones
func ComputeRouteId(method string, url string) string {
	return fmt.Sprintf("%s %s", method, url)
}

// GenerateUrlPatternMatcher creates a pattern that allows us to
// extract values from url.
func GenerateUrlPatternMatcher(method string, url string) (*regexp.Regexp, error) {
	pattern := sampleRegexp.ReplaceAllString(
		fmt.Sprintf("^%s %s$", method, url),
		"(?P<${1}>.*)",
	)
	return regexp.Compile(pattern)
}
