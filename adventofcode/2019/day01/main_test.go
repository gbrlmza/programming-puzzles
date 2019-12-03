package main

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name:       "Case 1",
			args:       args{input: []int{1969}},
			wantResult: 654,
		},
		{
			name:       "Case 2",
			args:       args{input: []int{100756}},
			wantResult: 33583,
		},
		{
			name:       "Case 3",
			args:       args{input: []int{1969, 100756}},
			wantResult: 34237,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Solve(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("Solve() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSolveTwo(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name:       "Case 1",
			args:       args{input: []int{1969}},
			wantResult: 654,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Solve(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("Solve() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
