package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// https://adventofcode.com/2021/day/12

const (
	start = "start"
	end   = "end"
)

type grid map[string]cave
type cave struct {
	name        string
	connections []string
	visits      int
}

func main() {
	resultOne := partOne(getInput("input.txt"))
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(getInput("sample.txt"))
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

// part one solution
func partOne(g grid) int {
	paths := navigate(g, start, nil)
	return len(paths)
}

// part two solution
func partTwo(g grid) int {
	return 0
}

func navigate(g grid, caveName string, path []string) []string {
	var paths []string

	cave := g[caveName]
	if !cave.canVisit() {
		return nil
	}
	if cave.name == end {
		path = append(path, caveName)
		paths = append(paths, strings.Join(path, ","))
		return paths
	}

	path = append(path, caveName)
	cave.visits++
	g[caveName] = cave
	for _, conn := range cave.connections {
		additionalPaths := navigate(g.copy(), conn, path)
		paths = append(paths, additionalPaths...)
	}

	return paths
}

func (c cave) canVisit() bool {
	return c.visits == 0 || c.name != strings.ToLower(c.name)
}

func (g grid) copy() grid {
	c := make(grid, len(g))
	for k, v := range g {
		c[k] = v
	}
	return c
}

func getInput(path string) grid {
	output := make(grid)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "-")
		caveA := output[values[0]]
		caveA.name = values[0]
		caveA.connections = append(caveA.connections, values[1])
		caveB := output[values[1]]
		caveB.name = values[1]
		caveB.connections = append(caveB.connections, values[0])
		output[caveA.name] = caveA
		output[caveB.name] = caveB
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}
