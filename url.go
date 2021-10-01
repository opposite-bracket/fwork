package fwork

import (
	"net/url"
	"strconv"
)

// ExtractIntQuery extracts an int value for Query parameter.
// Converts to int, Sets a cap and a default value
func ExtractIntQuery(url *url.URL, key string, maxVal int, defVal int) int {
	if strVal := url.Query().Get(key); strVal == "" {
		return defVal
	} else if intVal, err := strconv.Atoi(strVal); err == nil && intVal > maxVal {
		return maxVal
	} else if err != nil {
		return defVal
	}

	return defVal
}

// ExtractStringQuery extracts a string from Query parameter
// Sets default value if absent
func ExtractStringQuery(url *url.URL, key string, defVal string) string {
	if strVal := url.Query().Get(key); strVal != "" {
		return strVal
	}

	return defVal
}
