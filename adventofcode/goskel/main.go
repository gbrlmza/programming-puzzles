package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
-- Problem description
*/

func main() {
	input := getInput("input.txt")

	resultOne := partOne(input)
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(input)
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

func partOne(input []int) (result int) {
	return 0
}

func partTwo(input []int) (result int) {
	return 0
}

func getInput(path string) []int {
	result := make([]int, 0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
