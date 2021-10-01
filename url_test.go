package fwork

import (
	"net/url"
	"testing"
)

func TestExtractIntQuery(t *testing.T) {
	type args struct {
		url    *url.URL
		key    string
		maxVal int
		defVal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test happy path",
			args: args{
				url: &url.URL{
					RawQuery: "some-key=100",
				},
				key:    "some-key",
				maxVal: 1000,
				defVal: 100,
			},
			want: 100,
		},
		{name: "test value getting default value",
			args: args{
				url:    &url.URL{},
				key:    "some-key",
				maxVal: 1000,
				defVal: 100,
			},
			want: 100,
		},
		{name: "test max value",
			args: args{
				url: &url.URL{
					RawQuery: "some-key=6000",
				},
				key:    "some-key",
				maxVal: 1000,
				defVal: 100,
			},
			want: 1000,
		},
		{name: "test invalid value",
			args: args{
				url: &url.URL{
					RawQuery: "some-key=abc",
				},
				key:    "some-key",
				maxVal: 1000,
				defVal: 100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractIntQuery(tt.args.url, tt.args.key, tt.args.maxVal, tt.args.defVal); got != tt.want {
				t.Errorf("ExtractIntQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractStringQuery(t *testing.T) {
	type args struct {
		url    *url.URL
		key    string
		defVal string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test happy path",
			args: args{
				url: &url.URL{
					RawQuery: "some-key=abc",
				},
				key:    "some-key",
				defVal: "hello-world",
			},
			want: "abc",
		},
		{name: "test default value",
			args: args{
				url:    &url.URL{},
				key:    "some-key",
				defVal: "hello-world",
			},
			want: "hello-world",
		},
		{name: "test default empty value",
			args: args{
				url:    &url.URL{},
				key:    "some-key",
				defVal: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractStringQuery(tt.args.url, tt.args.key, tt.args.defVal); got != tt.want {
				t.Errorf("ExtractStringQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
