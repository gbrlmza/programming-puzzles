package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/9

type HeightMap map[coord]int

type coord struct {
	x int
	y int
}

var adj = []coord{{x: 0, y: -1}, {x: -1, y: 0}, {x: +1, y: 0}, {x: 0, y: +1}}

func main() {
	in := getInput("input.txt")

	resultOne := partOne(in)
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(in)
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

func partOne(hm HeightMap) int {
	var count int
	for c := range hm {
		h := hm.isLowpoint(c.x, c.y)
		if h > 0 {
			count += h
		}
	}
	return count
}

func partTwo(in HeightMap) int {
	return 0
}

func getCoord(x, y int) coord {
	return coord{x: x, y: y}
}

func (c coord) adjacent() []coord {
	coords := make([]coord, 0, len(adj))
	for _, a := range adj {
		coords = append(coords, coord{x: c.x + a.x, y: c.y + a.y})
	}
	return coords
}

func (hm HeightMap) isLowpoint(x, y int) int {
	coord := getCoord(x, y)
	height := hm[coord]
	for _, ac := range coord.adjacent() {
		if adjHeight, ok := hm[ac]; ok && adjHeight <= height {
			return 0
		}
	}
	return height + 1
}

func getInput(path string) HeightMap {
	output := make(HeightMap)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var x int
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), "")
		for y, v := range points {
			depth, _ := strconv.Atoi(v)
			output[getCoord(x, y)] = depth
		}
		x++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}
