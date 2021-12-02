package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	type args struct {
		input []Cmd
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "example",
			args: args{
				input: []Cmd{
					{Direction: "forward", Distance: 5},
					{Direction: "down", Distance: 5},
					{Direction: "forward", Distance: 8},
					{Direction: "up", Distance: 3},
					{Direction: "down", Distance: 8},
					{Direction: "forward", Distance: 2},
				},
			},
			wantResult: 150,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := PartOne(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("PartOne() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	type args struct {
		input []Cmd
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "example",
			args: args{
				input: []Cmd{
					{Direction: "forward", Distance: 5},
					{Direction: "down", Distance: 5},
					{Direction: "forward", Distance: 8},
					{Direction: "up", Distance: 3},
					{Direction: "down", Distance: 8},
					{Direction: "forward", Distance: 2},
				},
			},
			wantResult: 900,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := PartTwo(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("PartTwo() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
