package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/15

type grid struct {
	locations map[loc]int
	width     int
	height    int
}
type loc struct {
	x int
	y int
}

func main() {
	resultOne := partOne(getInput("input.txt"))
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(getInput("sample.txt"))
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

// part one solution
func partOne(g grid) int {
	return 0
}

// part two solution
func partTwo(g grid) int {
	return 0
}

func getInput(path string) grid {
	output := grid{
		locations: make(map[loc]int),
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 1
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "")
		output.width = len(values)
		for x, v := range values {
			n, _ := strconv.Atoi(v)
			output.locations[loc{x: x + 1, y: y}] = n
		}
		y++
	}
	output.height = y - 1

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}
