package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2021/day/14

type input struct {
	template   string
	insertions map[string]string
}
type pairs map[string]int

func main() {
	resultOne := partOne(getInput("input.txt"))
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(getInput("input.txt"))
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

// part one solution
func partOne(in input) int {
	return solve(in, 10)
}

// part two solution
func partTwo(in input) int {
	return solve(in, 40)
}

func solve(in input, steps int) int {
	ps := in.getPairs()
	for i := 0; i < steps; i++ {
		buffer := make(pairs)
		for pair, count := range ps {
			if ins, ok := in.insertions[pair]; ok {
				buffer[string(pair[0])+ins] += count
				buffer[ins+string(pair[1])] += count
				continue
			}
			buffer[pair] += count
		}
		ps = buffer
	}
	most, least := in.getMostAndLeastCommon(ps)
	return most - least
}

func (in input) getPairs() pairs {
	p := make(pairs)
	for i := 0; i < len(in.template)-1; i++ {
		pair := in.template[i : i+2]
		p[pair]++
	}
	return p
}

func (in input) getMostAndLeastCommon(ps pairs) (int, int) {
	var most, least string
	count := map[string]int{
		in.template[len(in.template)-1:]: 1, // +1 for last element
	}
	for pair, n := range ps { // count first element of each pair
		count[string(pair[0])] += n
	}
	for elem := range count {
		if most == "" || count[most] < count[elem] {
			most = elem
		}
		if least == "" || count[least] > count[elem] {
			least = elem
		}
	}
	return count[most], count[least]
}

func getInput(path string) input {
	in := input{
		insertions: make(map[string]string),
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		if in.template == "" {
			in.template = text
			continue
		}
		parts := strings.Split(text, " -> ")
		in.insertions[parts[0]] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return in
}
