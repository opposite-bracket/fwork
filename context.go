package fwork

import (
	"encoding/json"
	"net/http"
)

type ReqContext struct {
	Req *http.Request
	Res http.ResponseWriter
}

// JsonReply prepares a response to http client in JSON format
func (c *ReqContext) JsonReply(status int, body interface{}) {
	c.Res.WriteHeader(status)
	c.Res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Res).Encode(body)
}
