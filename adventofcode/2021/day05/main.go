package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/5

func main() {
	input := getInput("input.txt")

	resultOne := partOne(input)
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(input)
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

func partOne(input [][]int) int {
	points := make(map[string]int)
	for _, l := range input {
		if l[0] != l[2] && l[1] != l[3] {
			// we only process vertical & horizontal lines in part one
			continue
		}
		for _, p := range getPoints(l[0], l[1], l[2], l[3]) {
			coord := fmt.Sprintf("%d,%d", p[0], p[1])
			points[coord]++
		}
	}
	result := 0
	for _, v := range points {
		if v >= 2 {
			result++
		}
	}
	return result
}

func partTwo(input [][]int) int {
	points := make(map[string]int)
	for _, l := range input {
		for _, p := range getPoints(l[0], l[1], l[2], l[3]) {
			coord := fmt.Sprintf("%d,%d", p[0], p[1])
			points[coord]++
		}
	}
	result := 0
	for _, v := range points {
		if v >= 2 {
			result++
		}
	}
	return result
}

func getInput(path string) [][]int {
	result := make([][]int, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, " -> ")
		p1 := strings.Split(points[0], ",")
		p2 := strings.Split(points[1], ",")
		p1x, _ := strconv.Atoi(p1[0])
		p1y, _ := strconv.Atoi(p1[1])
		p2x, _ := strconv.Atoi(p2[0])
		p2y, _ := strconv.Atoi(p2[1])
		result = append(result, []int{p1x, p1y, p2x, p2y})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

// getPoints return points between two coordinates (x1,y1) - > (x2,y2)
// it only works with horizontal, vertical and 45° lines
// getPoints(0,0,5,0) returns [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}}
// getPoints(3,10,0,7) returns [][]int{{3, 10}, {2, 9}, {1, 8}, {0, 7}}
func getPoints(x1, y1, x2, y2 int) [][]int {
	if x1 != x2 && y2 != y1 && math.Abs(float64(x2-x1)) != math.Abs(float64(y2-y1)) {
		// safety check to avoid processing non-supported line types
		return nil
	}

	xStep := 1
	if x2 < x1 {
		xStep = -1
	} else if x2 == x1 {
		xStep = 0
	}
	yStep := 1
	if y2 < y1 {
		yStep = -1
	} else if y2 == y1 {
		yStep = 0
	}
	result := [][]int{}
	morePoints := true
	i := 0
	for morePoints {
		x := x1 + i*xStep
		y := y1 + i*yStep
		result = append(result, []int{x, y})
		if x == x2 && y == y2 {
			morePoints = false
		}
		i++
	}
	return result
}
