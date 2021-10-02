package fwork

import (
	"net/http"
	"testing"
)

func TestApiError_Error(t *testing.T) {
	type fields struct {
		Status  int
		Code    string
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "happy path",
			fields: fields{
				Status:  http.StatusInternalServerError,
				Code:    "fwork1",
				Message: "something wrong",
			},
			want: "[500:fwork1] something wrong",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ApiError{
				Status:  tt.fields.Status,
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ApiError() = %v, want %v", got, tt.want)
			}
		})
	}
}
