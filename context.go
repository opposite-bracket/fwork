package fwork

import (
	"encoding/json"
	"net/http"
)

type Params struct {
	Url map[string]string
}

type ReqContext struct {
	Req    *http.Request
	Res    http.ResponseWriter
	Params Params
}

// JsonReply prepares a response to http client in JSON format
func (c *ReqContext) JsonReply(status int, body interface{}) {
	c.Res.WriteHeader(status)
	c.Res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Res).Encode(body)
}

func NewReqContext(w http.ResponseWriter, r *http.Request) *ReqContext {
	return &ReqContext{
		Req: r,
		Res: w,
		Params: Params{Url: make(map[string]string, 0)},
	}
}
