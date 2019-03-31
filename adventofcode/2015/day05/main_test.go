package main

import (
	"testing"

	"gotest.tools/assert"
)

var testCasesOne = map[string]int{
	"aeiouaeiouaeiou":  0,
	"xazegov":          0,
	"abcdde":           0,
	"aabbccdd":         0,
	"aeiouaeiouabeiou": 0,
	"xazegcdov":        0,
	"abcpqdde":         0,
	"aabbccddxy":       0,
	"jchzalrnumimnmhp": 0,
	"haegwjzuvuyypxyu": 0,
	"dvszwmarrgswjxmb": 0,
	"ugknbfddgicrmopn": 1,
	"aaa":              1,
}

func TestCaseOne(t *testing.T) {
	for input, expectedResult := range testCasesOne {
		result := Solve(input)
		assert.Equal(t, expectedResult, result, input)
	}
}
