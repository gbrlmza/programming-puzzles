/*
--- Day 6: Probably a Fire Hazard ---
Because your neighbors keep defeating you in the holiday house decorating contest year after year,
you've decided to deploy one million lights in a 1000x1000 grid.

Furthermore, because you've been especially nice this year, Santa has mailed you instructions on
how to display the ideal lighting configuration.

Lights in your grid are numbered from 0 to 999 in each direction; the lights at each corner are
at 0,0, 0,999, 999,999, and 999,0. The instructions include whether to turn on, turn off, or
toggle various inclusive ranges given as coordinate pairs. Each coordinate pair represents
opposite corners of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore
refers to 9 lights in a 3x3 square. The lights all start turned off.

To defeat your neighbors this year, all you have to do is set up your lights by doing the
instructions Santa sent you in order.

For example:

- turn on 0,0 through 999,999 would turn on (or leave on) every light.
- toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off the ones
  that were on, and turning on the ones that were off.
- turn off 499,499 through 500,500 would turn off (or leave off) the middle four lights.

After following the instructions, how many lights are lit?

--- Part Two ---
You just finish implementing your winning light pattern when you realize you mistranslated
Santa's message from Ancient Nordic Elvish.

The light grid you bought actually has individual brightness controls; each light can have
a brightness of zero or more. The lights all start at zero.

The phrase turn on actually means that you should increase the brightness of thoselights by 1.

The phrase turn off actually means that you should decrease the brightness of those lights by 1,
to a minimum of zero.

The phrase toggle actually means that you should increase the brightness of those lights by 2.

What is the total brightness of all lights combined after following Santa's instructions?

For example:

turn on 0,0 through 0,0 would increase the total brightness by 1.
toggle 0,0 through 999,999 would increase the total brightness by 2000000.

*/
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

func Solve(input string) int {
	steps := parseInpt(input)
	grid := [1000][1000]bool{}

	for _, step := range steps {
		for x := step[1]; x <= step[3]; x++ {
			for y := step[2]; y <= step[4]; y++ {
				if step[0] == 0 {
					grid[x][y] = false
				} else if step[0] == 1 {
					grid[x][y] = true
				} else {
					grid[x][y] = !grid[x][y]
				}
			}
		}
	}

	lightsOn := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] {
				lightsOn++
			}
		}
	}

	return lightsOn
}

func SolveTwo(input string) int {
	steps := parseInpt(input)
	grid := [1000][1000]int{}

	for _, step := range steps {
		for x := step[1]; x <= step[3]; x++ {
			for y := step[2]; y <= step[4]; y++ {
				if step[0] == 0 && grid[x][y] > 0 {
					grid[x][y]--
				} else if step[0] == 1 {
					grid[x][y]++
				} else if step[0] == 2 {
					grid[x][y] += 2
				}
			}
		}
	}

	totalBrightness := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			totalBrightness += grid[x][y]
		}
	}

	return totalBrightness
}

func parseInpt(input string) [][5]int64 {
	lines := strings.Split(input, "\n")
	steps := make([][5]int64, len(lines))

	for i, line := range lines {
		line = strings.Replace(line, " through ", ",", -1)
		line = strings.Replace(line, "turn ", "", -1)
		line = strings.Replace(line, " ", ",", -1)
		parts := strings.Split(line, ",")

		if parts[0] == "off" {
			steps[i][0] = 0
		} else if parts[0] == "on" {
			steps[i][0] = 1
		} else {
			steps[i][0] = 2
		}
		x0, _ := strconv.ParseInt(parts[1], 10, 64)
		y0, _ := strconv.ParseInt(parts[2], 10, 64)
		x1, _ := strconv.ParseInt(parts[3], 10, 64)
		y1, _ := strconv.ParseInt(parts[4], 10, 64)
		steps[i][1], steps[i][2], steps[i][3], steps[i][4] = x0, y0, x1, y1
	}

	return steps
}
