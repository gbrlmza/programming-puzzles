package main

import (
	"testing"

	"gotest.tools/assert"
)

var testCasesOne = map[string]int{
	"^v^v^v^v^v": 2,
	"^>v<":       4,
}

var testCasesTwo = map[string]int{
	"^>v<":       3,
	"^v^v^v^v^v": 11,
}

func TestCaseOne(t *testing.T) {
	for input, expectedResult := range testCasesOne {
		result := Solve(input)
		assert.Equal(t, expectedResult, result)
	}
}

func TestCaseTwo(t *testing.T) {
	for input, expectedResult := range testCasesTwo {
		result := SolveTwo(input)
		assert.Equal(t, expectedResult, result)
	}
}
