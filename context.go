package fwork

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqContext struct {
	Req       *http.Request
	Res       http.ResponseWriter
	UrlParams map[string]string
}

// JsonReply prepares a response to http client in JSON format
func (c *ReqContext) JsonReply(status int, body interface{}) error {
	c.Res.Header().Set("Content-Type", "application/json")
	c.Res.WriteHeader(status)
	return json.NewEncoder(c.Res).Encode(body)
}

// GetIntQuery extracts an int value from Query parameter.
// Converts to int, Sets a cap and a default value
func (c *ReqContext) GetIntQuery(key string, maxVal int, defVal int) int {
	if strVal := c.Req.URL.Query().Get(key); strVal == "" {
		return defVal
	} else if intVal, err := strconv.Atoi(strVal); err == nil && intVal > maxVal {
		return maxVal
	} else if err != nil {
		return defVal
	} else {
		return intVal
	}
}

// GetInt64Query extracts an int value from Query parameter.
// Converts to int, Sets a cap and a default value
func (c *ReqContext) GetInt64Query(key string, maxVal int64, defVal int64) int64 {
	if strVal := c.Req.URL.Query().Get(key); strVal == "" {
		return defVal
	} else if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil && intVal > maxVal {
		return maxVal
	} else if err != nil {
		return defVal
	} else {
		return intVal
	}
}

// GetStringQuery extracts a string from Query parameter
// Sets default value if absent
func (c *ReqContext) GetStringQuery(key string, defVal string) string {
	if strVal := c.Req.URL.Query().Get(key); strVal != "" {
		return strVal
	}

	return defVal
}

// GetIntUrlParam extracts an int value from URL parameter.
// Converts to int, Sets a cap and a default value
func (c *ReqContext) GetIntUrlParam(key string, maxVal int, defVal int) int {
	if strVal := c.UrlParams[key]; strVal == "" {
		return defVal
	} else if intVal, err := strconv.Atoi(strVal); err == nil && intVal > maxVal {
		return maxVal
	} else if err != nil {
		return defVal
	} else {
		return intVal
	}
}

// GetStringUrlParam extracts a string from Url parameter
// Sets default value if absent
func (c *ReqContext) GetStringUrlParam(key string, defVal string) string {
	if strVal := c.UrlParams[key]; strVal != "" {
		return strVal
	}

	return defVal
}

// NewReqContext instantiates a req context
func NewReqContext(w http.ResponseWriter, r *http.Request) *ReqContext {
	return &ReqContext{
		Req:       r,
		Res:       w,
		UrlParams: make(map[string]string, 0),
	}
}
