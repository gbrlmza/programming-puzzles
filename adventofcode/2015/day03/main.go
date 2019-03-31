/*
--- Day 3: Perfectly Spherical Houses in a Vacuum ---
Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and
then an elf at the North Pole calls him via radio and tells him where to move
next. Moves are always exactly one house to the north (^), south (v), east (>),
or west (<). After each move, he delivers another present to the house at his
new location.

However, the elf back at the north pole has had a little too much eggnog, and
so his directions are a little off, and Santa ends up visiting some houses
more than once. How many houses receive at least one present?

For example:

- > delivers presents to 2 houses: one at the starting location, and one to the east.
- ^>v< delivers presents to 4 houses in a square, including twice to the house
  at his starting/ending location.
- ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

--- Part Two ---
The next year, to speed up the process, Santa creates a robot version of himself,
Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the same
starting house), then take turns moving based on instructions from the elf, who
is eggnoggedly reading from the same script as the previous year.

This year, how many houses receive at least one present?

For example:

- ^v delivers presents to 3 houses, because Santa goes north, and then
	Robo-Santa goes south.
- ^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up
	back where they started.
- ^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and
	Robo-Santa going the other.
*/
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input := GetInput("input.txt")

	solution := Solve(input)
	fmt.Println("Solution One:")
	fmt.Println(solution)

	solution = SolveTwo(input)
	fmt.Println("Solution Two:")
	fmt.Println(solution)
}

// GetInput return content of file as string
func GetInput(file string) string {
	data, _ := ioutil.ReadFile(file)
	return string(data)
}

// moves modification of X,Y axis for each posible move
var moves = map[string][]int{
	">": {1, 0},
	"<": {-1, 0},
	"^": {0, 1},
	"v": {0, -1},
}

// Solve Solution for first part of the problem
func Solve(input string) int {
	var x0, y0 int             // Stating position 0,0
	houses := map[string]int{} // Visited houses

	for _, v := range input {
		move := string(v)

		// Calculate new coordinates
		x1, y1 := x0+moves[move][0], y0+moves[move][1]

		// House identifiers
		k0, k1 := fmt.Sprintf("%d,%d", x0, y0), fmt.Sprintf("%d,%d", x1, y1)

		// Increase gifts counter per house
		houses[k0]++
		houses[k1]++

		// Set current position
		x0, y0 = x1, y1
	}

	return len(houses)
}

// SolveTwo Solution for second part of the problem
func SolveTwo(input string) int {
	pos := [8]int{}
	houses := map[string]int{}

	for k, v := range input {
		move := string(v)
		offset := 0
		santaTurn := k%2 == 0
		if !santaTurn {
			offset = 4
		}

		x0, y0, x1, y1 := offset, offset+1, offset+2, offset+3

		// Calculate new coordinates
		pos[x1], pos[y1] = pos[x0]+moves[move][0], pos[y0]+moves[move][1]

		// House identifiers
		h0, h1 := fmt.Sprintf("%d,%d", pos[x0], pos[y0]), fmt.Sprintf("%d,%d", pos[x1], pos[y1])

		// Increase gifts counter per house
		houses[h0]++
		houses[h1]++

		// Update current position
		pos[x0], pos[y0] = pos[x1], pos[y1]
	}

	return len(houses)
}
