package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/4

func main() {
	resultOne := partOne(getInput("input.txt"))
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(getInput("input.txt"))
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

func partOne(g game) int {
	for _, n := range g.Numbers {
		for i := range g.Boards {
			b := &g.Boards[i]
			if b.mark(n) {
				return b.score() * n
			}
		}
	}
	return 0
}

func partTwo(g game) int {
	var lastWin board
	var lastNum int
	for _, n := range g.Numbers {
		for i := range g.Boards {
			b := &g.Boards[i]
			if b.Won {
				continue
			}
			if b.mark(n) {
				lastWin = *b
				lastNum = n
			}
		}
	}
	return lastWin.score() * lastNum
}

func getInput(path string) game {
	g := game{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := -1
	b := board{}
	x := 1
	y := 1
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if lineNum == 0 {
			nums := strings.Split(line, ",")
			for _, nStr := range nums {
				n, _ := strconv.Atoi(nStr)
				g.Numbers = append(g.Numbers, n)
			}
			continue
		}

		if line == "" { // board separator
			if len(b.Cells) > 0 {
				g.Boards = append(g.Boards, b)
			}
			b = board{ // initialize new board
				Cells: make(map[int]cell, boardSize),
				Marks: make(map[int]int, boardSize),
			}
			x = 1
			continue
		}

		nums := strings.Split(line, " ")
		y = 1
		for _, nStr := range nums {
			n, err := strconv.Atoi(nStr)
			if err != nil {
				continue
			}
			b.Cells[n] = cell{X: x, Y: y}
			y++
		}
		x++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return g
}
