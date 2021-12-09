package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/9

const maxHeight = 9

type (
	heightMap map[location]int
	location  struct {
		x int
		y int
	}
)

// adjacent offsets to get adjacent locations
var adjacent = []location{{x: 0, y: -1}, {x: -1, y: 0}, {x: +1, y: 0}, {x: 0, y: +1}}

func main() {
	in := getInput("input.txt")

	resultOne := partOne(in)
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(in)
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

// partOne solution to https://adventofcode.com/2021/day/9
func partOne(hm heightMap) int {
	var sum int
	for loc := range hm {
		riskLevel := hm.getRiskLevel(loc)
		if riskLevel > 0 {
			sum += riskLevel
		}
	}
	return sum
}

// partTwo solution to https://adventofcode.com/2021/day/9#part2
func partTwo(hm heightMap) int {
	// get basins sizes
	sizes := make([]int, 3)
	for loc := range hm {
		riskLevel := hm.getRiskLevel(loc)
		if riskLevel > 0 {
			sizes = append(sizes, hm.getBasinSize(loc))
		}
	}

	// multiply the 3 largest basins
	sort.Ints(sizes)
	result := 1
	for i := len(sizes) - 1; i < 0 || i > len(sizes)-4; i-- {
		if sizes[i] != 0 {
			result *= sizes[i]
		}
	}

	return result
}

// adjacent returns adjacents locations of the current location
func (l location) adjacent() []location {
	coords := make([]location, 0, len(adjacent))
	for _, a := range adjacent {
		coords = append(coords, location{x: l.x + a.x, y: l.y + a.y})
	}
	return coords
}

// copy returns a copy of the heightMap
func (hm heightMap) copy() heightMap {
	new := make(heightMap, len(hm))
	for k, v := range hm {
		new[k] = v
	}
	return new
}

// getRiskLevel returns the risk level of a low point.
// returns 0 if the provided location isn't a low point
func (hm heightMap) getRiskLevel(l location) int {
	height, ok := hm[l]
	if !ok || height == maxHeight {
		return 0
	}
	for _, offsets := range l.adjacent() {
		if adjHeight, ok := hm[offsets]; ok && adjHeight <= height {
			return 0
		}
	}
	return height + 1
}

// getBasinSize returns the size of the basin that the provided location belongs to
func (hm heightMap) getBasinSize(c location) int {
	return calculateBasinSize(hm.copy(), c.x, c.y, 0, true, true, true, true)
}

// calculateBasinSize calculates basin size
func calculateBasinSize(hm heightMap, x, y, size int, up, down, left, right bool) int {
	loc := location{x: x, y: y}
	height, ok := hm[loc]
	if height == maxHeight || !ok {
		return size
	}
	size++
	delete(hm, loc)

	if up {
		size = calculateBasinSize(hm, x, y-1, size, true, false, true, true)
	}
	if down {
		size = calculateBasinSize(hm, x, y+1, size, false, true, true, true)
	}
	if left {
		size = calculateBasinSize(hm, x-1, y, size, true, true, true, false)
	}
	if right {
		size = calculateBasinSize(hm, x+1, y, size, true, true, false, true)
	}

	return size
}

func getInput(path string) heightMap {
	output := make(heightMap)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var y int
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), "")
		for x, v := range points {
			height, _ := strconv.Atoi(v)
			output[location{x: x, y: y}] = height
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}
