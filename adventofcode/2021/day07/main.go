package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/7

func main() {
	input := getInput("input.txt")

	resultOne := partOne(input)
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(input)
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

func partOne(input []int) int {
	var cost, median int
	sort.Ints(input)
	median = input[len(input)/2]
	for _, n := range input {
		cost += int(math.Abs(float64(n - median)))
	}
	return cost
}

func partTwo(input []int) int {
	var cost, mean int
	for _, n := range input {
		mean += n
	}
	mean = mean / len(input)

	sort.Ints(input)
	maxDistance := int(math.Abs(float64(mean - input[0])))
	if math.Abs(float64(mean-input[len(input)-1])) > float64(maxDistance) {
		maxDistance = int(math.Abs(float64(mean - input[len(input)-1])))
	}
	distancesCost := make(map[int]int, maxDistance)
	for i := 0; i <= maxDistance; i++ {
		distancesCost[i] = distancesCost[i-1] + i
	}
	for _, p := range input {
		cost += distancesCost[int(math.Abs(float64(mean-p)))]
	}
	return cost
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

// func partOne(input []int) int {
// 	sort.Ints(input)
// 	positions := make(map[int]int)
// 	for i := input[0]; i <= input[len(input)-1]; i++ {
// 		positions[i] = 0
// 	}
// 	for _, v := range input {
// 		for p := range positions {
// 			positions[p] += int(math.Abs(float64(v - p)))
// 		}
// 	}
// 	cost := -1
// 	for _, pCost := range positions {
// 		if cost == -1 || pCost < cost {
// 			cost = pCost
// 		}
// 	}
// 	return cost
// }

// func partTwo(input []int) int {
// 	sort.Ints(input)
// 	positions := make(map[int]int)
// 	for i := input[0]; i <= input[len(input)-1]; i++ {
// 		positions[i] = 0
// 	}

// 	maxDistance := input[len(input)-1] - input[0]
// 	distancesCost := make(map[int]int, maxDistance)
// 	for i := 0; i <= maxDistance; i++ {
// 		distancesCost[i] = distancesCost[i-1] + i
// 	}

// 	for _, v := range input {
// 		for p := range positions {
// 			positions[p] += distancesCost[int(math.Abs(float64(v-p)))]
// 		}
// 	}

// 	cost := -1
// 	for _, pCost := range positions {
// 		if cost == -1 || pCost < cost {
// 			cost = pCost
// 		}
// 	}

// 	return cost
// }
