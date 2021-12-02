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
	input := GetInput("input.txt")

	resultOne := PartOne(input)
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := PartTwo(input)
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

func PartOne(input []int) (result int) {
	return 0
}

func PartTwo(input []int) (result int) {
	return 0
}

func GetInput(path string) []int {
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
