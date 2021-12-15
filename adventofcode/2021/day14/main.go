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

func main() {
	resultOne := partOne(getInput("input.txt"))
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(getInput("input.txt"))
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

// part one solution
func partOne(in input) int {
	var buffer string
	for i := 0; i < 10; i++ {
		buffer = in.template[:1]
		pairs := in.getPairs()
		for _, pair := range pairs {
			if insertion, ok := in.insertions[pair]; ok {
				buffer += fmt.Sprintf("%s%s", insertion, pair[1:])
			} else {
				buffer += pair[1:]
			}
		}
		in.template = buffer
	}
	most, least := in.getMostAndLeastCommon()
	return most - least
}

// part two solution
func partTwo(in input) int {
	var buffer string
	for i := 0; i < 40; i++ {
		buffer = in.template[:1]
		pairs := in.getPairs()
		for _, pair := range pairs {
			if insertion, ok := in.insertions[pair]; ok {
				buffer += fmt.Sprintf("%s%s", insertion, pair[1:])
			} else {
				buffer += pair[1:]
			}
		}
		in.template = buffer
	}
	most, least := in.getMostAndLeastCommon()
	return most - least
}

func (in input) getPairs() []string {
	pairs := []string{}
	for i := 0; i < len(in.template)-1; i++ {
		pairs = append(pairs, in.template[i:i+2])
	}
	return pairs
}

func (in input) getMostAndLeastCommon() (int, int) {
	var most, least string
	count := make(map[string]int)
	for _, r := range in.template {
		count[string(r)]++
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
