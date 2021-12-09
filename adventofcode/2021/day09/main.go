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
		h := hm.isLowpoint(c)
		if h > 0 {
			count += h
		}
	}
	return count
}

func partTwo(hm HeightMap) int {
	sizes := make([]int, 3)
	for c := range hm {
		h := hm.isLowpoint(c)
		if h > 0 {
			sizes = append(sizes, hm.getBasin(c))
		}
	}
	sort.Ints(sizes)
	count := len(sizes)
	return sizes[count-1] * sizes[count-2] * sizes[count-3]
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

func (hm HeightMap) copy() HeightMap {
	new := make(HeightMap, len(hm))
	for k, v := range hm {
		new[k] = v
	}
	return new
}

func (hm HeightMap) isLowpoint(c coord) int {
	height := hm[c]
	for _, ac := range c.adjacent() {
		if adjHeight, ok := hm[ac]; ok && adjHeight <= height {
			return 0
		}
	}
	return height + 1
}

func (hm HeightMap) getBasin(c coord) int {
	return getBasinHelper(hm.copy(), c.x, c.y, 0, true, true, true, true)
}

func getBasinHelper(hm HeightMap, x, y, size int, up, down, left, right bool) int {
	coord := coord{x: x, y: y}
	height, ok := hm[coord]
	if height == 9 || !ok {
		return size
	}
	size++
	delete(hm, coord)

	if up {
		size = getBasinHelper(hm, x, y-1, size, true, false, true, true)
	}
	if down {
		size = getBasinHelper(hm, x, y+1, size, false, true, true, true)
	}
	if left {
		size = getBasinHelper(hm, x-1, y, size, true, true, true, false)
	}
	if right {
		size = getBasinHelper(hm, x+1, y, size, true, true, false, true)
	}

	return size
}

func getInput(path string) HeightMap {
	output := make(HeightMap)

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
			depth, _ := strconv.Atoi(v)
			output[getCoord(x, y)] = depth
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}
