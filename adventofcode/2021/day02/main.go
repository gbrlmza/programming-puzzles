package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
--- Description ---
*/

func main() {
	input := GetInput("input.txt")

	result := Solve(input)
	fmt.Printf("Result: %v\n", result)
}

func Solve(input []int) (result int) {
	return 0
}

// GetInput return content of file as string
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
