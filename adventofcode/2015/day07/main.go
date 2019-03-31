// Day 7: Some Assembly Required
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input := GetInput("input.txt")
	fmt.Println(input)
}

// GetInput return content of file as string
func GetInput(file string) string {
	data, _ := ioutil.ReadFile(file)
	return string(data)
}

// ParseInput return instructions
func ParseInput(input string) {
	regex := `(\w)? ?(NOT|AND|OR|LSHIFT|RSHIFT) (\w) -> (\w)|(\d+) +(->) (\w)`
	fmt.Println(regex)
}
