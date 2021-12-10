package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// https://adventofcode.com/2021/day/10

var (
	missingPoints = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	autocompletePoint = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	pairs = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
)

func main() {
	in := getInput("input.txt")

	resultOne := partOne(in)
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(in)
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

// part one solution
func partOne(in []string) int {
	var score int
	var stack, char string
	for _, line := range in {
		stack = ""
		for _, rune := range line {
			char = string(rune)
			match, isOpening := pairs[char]
			if isOpening {
				stack = match + stack
				continue
			} else {
				if len(stack) == 0 || string(stack[0]) != char {
					score += missingPoints[char]
					break
				} else {
					stack = stack[1:]
				}
			}
		}
	}
	return score
}

// part two solution
func partTwo(in []string) int {
	var scores []int
	var partial int
	var stack, char string
	var corrupted bool
	for _, line := range in {
		stack = ""
		corrupted = false
		for _, r := range line {
			char = string(r)
			match, isOpening := pairs[char]
			if isOpening {
				stack = match + stack
				continue
			} else {
				if len(stack) == 0 || string(stack[0]) != char {
					corrupted = true
					break
				} else {
					stack = stack[1:]
				}
			}
		}
		if corrupted {
			continue
		}
		partial = 0
		for _, r := range stack {
			char = string(r)
			partial = partial*5 + autocompletePoint[char]
		}
		scores = append(scores, partial)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func getInput(path string) []string {
	output := []string{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}
