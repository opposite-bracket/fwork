package fwork

import (
	"fmt"
)

func ComputeRouteId(method string, url string) string {
	return fmt.Sprintf("%s %s", method, url)
}
