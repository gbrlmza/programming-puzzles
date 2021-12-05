package main

import (
	"reflect"
	"testing"
)

func Test_getPoints(t *testing.T) {
	type args struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "single ponint",
			args: args{x1: 10, y1: 7, x2: 10, y2: 7},
			want: [][]int{{10, 7}},
		},
		{
			name: "h",
			args: args{x1: 0, y1: 0, x2: 5, y2: 0},
			want: [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}},
		},
		{
			name: "h",
			args: args{x1: 0, y1: 0, x2: 0, y2: 5},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}},
		},
		{
			name: "o1",
			args: args{x1: 9, y1: 7, x2: 7, y2: 9},
			want: [][]int{{9, 7}, {8, 8}, {7, 9}},
		},
		{
			name: "o2",
			args: args{x1: 1, y1: 1, x2: 3, y2: 3},
			want: [][]int{{1, 1}, {2, 2}, {3, 3}},
		},
		{
			name: "o3",
			args: args{x1: 3, y1: 10, x2: 0, y2: 7},
			want: [][]int{{3, 10}, {2, 9}, {1, 8}, {0, 7}},
		},
		{
			name: "non-supported case 1",
			args: args{x1: 3, y1: 10, x2: 20, y2: 7},
			want: nil,
		},
		{
			name: "non-supported case 2",
			args: args{x1: 3, y1: 10, x2: 0, y2: 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPoints(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPoints() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
