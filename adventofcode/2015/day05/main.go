/*
Santa needs help figuring out which strings in his text file are naughty or nice.

A nice string is one with all of the following properties:

- It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
- It contains at least one letter that appears twice in a row, like xx, abcdde (dd),
or aabbccdd (aa, bb, cc, or dd).
- It does not contain the strings ab, cd, pq, or xy, even if they are part of one of
the other requirements.

For example:

- ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double
	letter (...dd...), and none of the disallowed substrings.
- aaa is nice because it has at least three vowels and a double letter, even though the
	letters used by different rules overlap.
- jchzalrnumimnmhp is naughty because it has no double letter.
- haegwjzuvuyypxyu is naughty because it contains the string xy.
- dvszwmarrgswjxmb is naughty because it contains only one vowel.

How many strings are nice?
*/
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	input := GetInput("input.txt")

	solution := Solve(input)
	fmt.Println("Solution One:")
	fmt.Println(solution)
}

// GetInput return content of file as string
func GetInput(file string) string {
	data, _ := ioutil.ReadFile(file)
	return string(data)
}

// Solve Return number of nice words in input
func Solve(input string) int {
	words := strings.Fields(input)
	count := 0

	for _, word := range words {
		word := []byte(strings.ToLower(word))
		vowels, _ := regexp.Match(`(?:[aeiou].*){3}`, word)
		// The standard regex engine of Go (RE2) doesn't support backrefereces:
		double, _ := regexp.Match(`aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz`, word)
		exclude, _ := regexp.Match(`(?:ab|cd|pq|xy)`, word)
		if vowels && double && !exclude {
			count++
		}
	}

	return count
}
