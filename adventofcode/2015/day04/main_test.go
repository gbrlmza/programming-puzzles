package main

import (
	"testing"

	"gotest.tools/assert"
)

var testCasesOne = map[string]int{
	"abcdef":  609043,
	"pqrstuv": 1048970,
}

func TestCaseOne(t *testing.T) {
	for input, expectedResult := range testCasesOne {
		result := Solve(input)
		assert.Equal(t, expectedResult, result)
	}
}
