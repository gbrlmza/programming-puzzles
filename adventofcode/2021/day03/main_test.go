package main

import (
	"testing"
)

func Test_partOne(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "partOne",
			args: args{
				input: []string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
			},
			wantResult: 198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := partOne(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("partOne() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_partTwo(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "partTwo",
			args: args{
				input: []string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
			},
			wantResult: 230,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := partTwo(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("partTwo() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
