package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/13

const (
	foldYPrefix = "fold along y="
	foldXPrefix = "fold along x="
)

type grid struct {
	points map[loc]int
	folds  []loc
	width  int
	height int
}
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
	for _, loc := range g.folds {
		g.fold(loc)
		break
	}
	return len(g.points)
}

// part two solution
func partTwo(g grid) int {
	for _, loc := range g.folds {
		g.fold(loc)
	}
	g.draw()
	return len(g.points)
}

// fold folds the paper on the provided location (x or y)
func (g *grid) fold(l loc) {
	var newX, newY int
	// move points
	for y := l.y; y <= g.height; y++ {
		for x := l.x; x <= g.width; x++ {
			pLoc := loc{x: x, y: y}
			if _, ok := g.points[pLoc]; !ok {
				continue
			}
			newX, newY = pLoc.x, pLoc.y
			if l.x > 0 {
				newX = g.width - pLoc.x
			}
			if l.y > 0 {
				newY = g.height - pLoc.y
			}
			newLoc := loc{x: newX, y: newY}
			delete(g.points, pLoc)
			g.points[newLoc] = 1
		}
	}
	// set new paper width & height
	if l.x > 0 {
		g.width = l.x - 1
	}
	if l.y > 0 {
		g.height = l.y - 1
	}
}

// draw prints the current points in the paper
func (g *grid) draw() {
	for y := 0; y <= g.height; y++ {
		fmt.Printf("%02d| ", y+1)
		for x := 0; x <= g.width; x++ {
			c := "."
			if g.points[loc{x: x, y: y}] == 1 {
				c = "#"
			}
			fmt.Print(c)
		}
		fmt.Print("\n")
	}
}

func getInput(path string) grid {
	g := grid{
		points: make(map[loc]int),
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
		if strings.HasPrefix(text, foldYPrefix) {
			y, _ := strconv.Atoi(strings.ReplaceAll(text, foldYPrefix, ""))
			g.folds = append(g.folds, loc{x: 0, y: y})
			continue
		}
		if strings.HasPrefix(text, foldXPrefix) {
			x, _ := strconv.Atoi(strings.ReplaceAll(text, foldXPrefix, ""))
			g.folds = append(g.folds, loc{x: x, y: 0})
			continue
		}

		coord := strings.Split(text, ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		g.points[loc{x: x, y: y}] = 1
		if x > g.width {
			g.width = x
		}
		if y > g.height {
			g.height = y
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return g
}
