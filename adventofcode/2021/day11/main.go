package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/11

const flashThreshold = 10

type grid map[loc]int
type loc struct {
	x int
	y int
}

func main() {
	resultOne := partOne(getInput("input.txt"))
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(getInput("input.txt"))
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

// part one solution
func partOne(g grid) int {
	var totalFlashes int
	for i := 0; i < 100; i++ {
		locs := g.increaseEnergy()
		for _, loc := range locs {
			g.flash(loc)
		}
		totalFlashes += g.finish()
	}

	return totalFlashes
}

// part two solution
func partTwo(g grid) int {
	step := 1
	for {
		locs := g.increaseEnergy()
		for _, loc := range locs {
			g.flash(loc)
		}
		flashes := g.finish()
		if flashes == len(g) {
			break
		}
		step++
	}
	return step
}

// increaseEnergy increases energy level of each location
// returns slice of locations that have energy to flash
func (g grid) increaseEnergy() []loc {
	var locs []loc
	for loc := range g {
		g[loc]++
		if g[loc] == flashThreshold {
			locs = append(locs, loc)
		}
	}
	return locs
}

// flash flashes the given location and propagates into adjacents locations
func (g grid) flash(l loc) {
	adjLocs := g.adjacents(l)
	for _, adjLoc := range adjLocs {
		if _, ok := g[adjLoc]; !ok {
			continue
		}
		g[adjLoc]++
		if g[adjLoc] == flashThreshold {
			g.flash(adjLoc)
		}
	}
}

// finish resets locations that flashed. returns the number of restored locations
func (g grid) finish() int {
	var flashes int
	for loc, energy := range g {
		if energy >= flashThreshold {
			flashes++
			g[loc] = 0
		}
	}
	return flashes
}

// adjacents returns adjacent locations to the provided location
func (g grid) adjacents(l loc) []loc {
	offsets := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	locations := make([]loc, len(offsets))
	for i, offset := range offsets {
		locations[i] = loc{x: l.x + offset[0], y: l.y + offset[1]}
	}
	return locations
}

func getInput(path string) grid {
	output := make(grid, 100)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 1
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "")
		for x, v := range values {
			n, _ := strconv.Atoi(v)
			output[loc{x: x + 1, y: y}] = n
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}
