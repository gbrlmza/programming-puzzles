package main

import "testing"

func TestSolveOne(t *testing.T) {
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
			args:       args{input: []int{1,9,10,3,2,3,11,0,99,30,40,50}},
			wantResult: 3500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := SolveOne(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("SolveOne() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}