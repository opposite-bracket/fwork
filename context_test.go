package fwork

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var sampleGetReq1, _ = http.NewRequest(http.MethodGet, "/hello?pageSize=10", nil)
var sampleGetReq2, _ = http.NewRequest(http.MethodGet, "/hello?pageSize=6000", nil)
var sampleGetReq3, _ = http.NewRequest(http.MethodGet, "/hello?pageSize=invalid", nil)
var sampleGetReq4, _ = http.NewRequest(http.MethodGet, "/hello?name=abc", nil)
var sampleGetReqNoUrlVars, _ = http.NewRequest(http.MethodGet, "/hello", nil)
var sampleRequest = ReqContext{
	Req:       sampleGetReq1,
	Res:       httptest.NewRecorder(),
	UrlParams: make(map[string]string, 0),
}

func TestReqContext_GetUrlIntParam(t *testing.T) {
	type fields struct {
		Req    *http.Request
		Res    http.ResponseWriter
		Params map[string]string
	}
	type args struct {
		key    string
		maxVal int
		defVal int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "happy path",
			fields: fields{
				Req: sampleGetReq1,
			},
			args: args{
				key:    "pageSize",
				maxVal: 1000,
				defVal: 1000,
			},
			want: 10,
		},
		{
			name: "test value getting default value",
			fields: fields{
				Req: sampleGetReqNoUrlVars,
			},
			args: args{
				key:    "pageSize",
				maxVal: 1000,
				defVal: 500,
			},
			want: 500,
		},
		{
			name: "test max value",
			fields: fields{
				Req: sampleGetReq2,
			},
			args: args{
				key:    "pageSize",
				maxVal: 1000,
				defVal: 500,
			},
			want: 1000,
		},
		{
			name: "test invalid value",
			fields: fields{
				Req: sampleGetReq3,
			},
			args: args{
				key:    "pageSize",
				maxVal: 1000,
				defVal: 500,
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ReqContext{
				Req: tt.fields.Req,
				Res: tt.fields.Res,
			}
			if got := c.GetIntQuery(tt.args.key, tt.args.maxVal, tt.args.defVal); got != tt.want {
				t.Errorf("GetIntQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReqContext_ExtractStringQuery(t *testing.T) {
	type fields struct {
		Req    *http.Request
		Res    http.ResponseWriter
		Params map[string]string
	}
	type args struct {
		key    string
		defVal string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "happy path",
			fields: fields{
				Req: sampleGetReq4,
			},
			args: args{
				key:    "name",
				defVal: "",
			},
			want: "abc",
		},
		{
			name: "test default value",
			fields: fields{
				Req: sampleGetReqNoUrlVars,
			},
			args: args{
				key:    "name",
				defVal: "abc",
			},
			want: "abc",
		},
		{
			name: "test default empty value",
			fields: fields{
				Req: sampleGetReqNoUrlVars,
			},
			args: args{
				key:    "name",
				defVal: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ReqContext{
				Req: tt.fields.Req,
				Res: tt.fields.Res,
			}
			if got := c.ExtractStringQuery(tt.args.key, tt.args.defVal); got != tt.want {
				t.Errorf("ExtractStringQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewReqContext(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
		UrlParams map[string]string
	}
	tests := []struct {
		name string
		args args
		want *ReqContext
	}{
		{
			name: "happy path",
			args: args{
				w: sampleRequest.Res,
				r: sampleRequest.Req,
				UrlParams: sampleRequest.UrlParams,
			},
			want: &sampleRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReqContext(tt.args.w, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReqContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
