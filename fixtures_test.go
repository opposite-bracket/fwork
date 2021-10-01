package fwork

import (
	"fmt"
	"reflect"
	"testing"
)

type Item struct {
	Field string
}

func GenerateItemsAsInterface(indexStartingPoint int, count int) []interface{} {
	var items []interface{}

	for i := indexStartingPoint; i < (indexStartingPoint + count); i++ {
		items = append(items, Item{Field: fmt.Sprintf("f-%d", i)})
	}

	return items
}

func TestExceedsPageCap(t *testing.T) {
	type args struct {
		pageNum  int
		pageSize int
		cap      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test happy path",
			args: args{
				pageNum:  3,
				pageSize: 100,
				cap:      1000,
			},
			want: false,
		},
		{
			name: "exceeds page cap",
			args: args{
				pageNum:  2,
				pageSize: 100,
				cap:      100,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExceedsPageCap(tt.args.pageNum, tt.args.pageSize, tt.args.cap); got != tt.want {
				t.Errorf("ExceedsPageCap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateFixtureList(t *testing.T) {
	type args struct {
		pageNum   int
		pageSize  int
		cap       int
		generator FixtureGenerator
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "happy path",
			args: args{
				pageNum:  1,
				pageSize: 10,
				cap:      200,
				generator: func(i int) interface{} {
					return Item{Field: fmt.Sprintf("f-%d", i)}
				},
			},
			want: GenerateItemsAsInterface(10, 10),
		},
		{
			name: "empty if exceeds cap",
			args: args{
				pageNum:  1,
				pageSize: 10,
				cap:      5,
				generator: func(i int) interface{} {
					return Item{Field: fmt.Sprintf("f-%d", i)}
				},
			},
			want: []interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateFixtureList(tt.args.pageNum, tt.args.pageSize, tt.args.cap, tt.args.generator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateFixtureList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexEndingPoint(t *testing.T) {
	type args struct {
		pageNum  int
		pageSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path",
			args: args{
				pageNum:  10,
				pageSize: 10,
			},
			want: 110,
		},
		{
			name: "zero test",
			args: args{
				pageNum:  0,
				pageSize: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexEndingPoint(tt.args.pageNum, tt.args.pageSize); got != tt.want {
				t.Errorf("IndexEndingPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexStartingPoint(t *testing.T) {
	type args struct {
		pageNum  int
		pageSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy path",
			args: args{
				pageNum:  10,
				pageSize: 10,
			},
			want: 100,
		},
		{
			name: "zero values",
			args: args{
				pageNum:  0,
				pageSize: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexStartingPoint(tt.args.pageNum, tt.args.pageSize); got != tt.want {
				t.Errorf("IndexStartingPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
