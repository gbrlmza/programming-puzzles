package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://adventofcode.com/2021/day/10

var (
	points = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
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
					score += points[char]
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
	return 0
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
