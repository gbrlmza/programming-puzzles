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
	insertions map[string][]string
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
	ps, last := in.getPairs()
	for i := 0; i < steps; i++ {
		buffer := make(pairs)
		for pair, count := range ps {
			if newPairs, ok := in.insertions[pair]; ok {
				buffer[newPairs[0]] += count
				buffer[newPairs[1]] += count
				if pair == last {
					last = newPairs[1]
				}
				continue
			}
			buffer[pair] += count
		}
		ps = buffer
	}
	most, least := getMostAndLeastCommon(ps, last)
	return most - least
}

func (in input) getPairs() (pairs, string) {
	var lastPair string
	p := make(pairs)
	for i := 0; i < len(in.template)-1; i++ {
		pair := in.template[i : i+2]
		p[pair]++
		lastPair = pair
	}
	return p, lastPair
}

func getMostAndLeastCommon(ps pairs, last string) (int, int) {
	var most, least string
	count := make(map[string]int)
	for pair, n := range ps {
		if pair == last {
			// add 1 for the last letter of the last pair
			// since we are only counting the first letter
			// of each pair
			count[string(pair[1])] += 1
		}
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
		insertions: make(map[string][]string),
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
		in.insertions[parts[0]] = []string{
			string(parts[0][0]) + string(parts[1]),
			string(parts[1]) + string(parts[0][1]),
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return in
}
