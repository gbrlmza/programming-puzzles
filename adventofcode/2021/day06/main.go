package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/6

func main() {
	resultOne := evolve(getInput("input.txt"), 80)
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := evolve(getInput("input.txt"), 256)
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

func evolve(initialState []int, days int) int {
	countByDays := make([]int, 9)
	for _, d := range initialState {
		countByDays[d]++
	}

	var count6, count8, day, count int
	for i := 0; i < days; i++ {
		count6, count8 = countByDays[6], countByDays[8]
		for day, count = range countByDays {
			switch day {
			case 0:
				countByDays[8] += count
				countByDays[6] += count
				countByDays[0] = 0
			case 8:
				countByDays[8] -= count8
				countByDays[7] += count8
			case 6:
				countByDays[6] -= count6
				countByDays[5] += count6
			default:
				countByDays[day] = 0
				countByDays[day-1] += count
			}
		}
	}

	var total int
	for _, c := range countByDays {
		total += c
	}
	return total
}

func getInput(path string) []int {
	var output []int

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numStr := strings.Split(scanner.Text(), ",")
		output = make([]int, len(numStr))
		for i, str := range numStr {
			n, _ := strconv.Atoi(str)
			output[i] = n
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}
