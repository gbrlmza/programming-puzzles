package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://adventofcode.com/2021/day/8

type input struct {
	pattern []string
	output  []string
}

func main() {
	in := getInput("input.txt")

	resultOne := partOne(in)
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(in)
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

func partOne(ins []input) int {
	var count int
	for _, in := range ins {
		for _, out := range in.output {
			d := digits(out)
			if d != 5 && d != 6 {
				count++
			}
		}
	}
	return count
}

func partTwo(ins []input) int {
	return 0
}

func digits(s string) int {
	digits := make(map[string]struct{}, 7)
	for _, c := range s {
		digits[string(c)] = struct{}{}
	}
	return len(digits)
}

func getInput(path string) []input {
	var in []input

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	v := make([]string, 14)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&v[0], &v[1], &v[2], &v[3], &v[4], &v[5], &v[6], &v[7], &v[8], &v[9], &v[10], &v[11], &v[12], &v[13])
		i := input{
			pattern: make([]string, 10),
			output:  make([]string, 4),
		}
		copy(i.pattern, v[:10])
		copy(i.output, v[10:])
		in = append(in, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return in
}
