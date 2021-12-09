package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

// https://adventofcode.com/2021/day/8

type input struct {
	signals []string
	output  []string
}

func main() {
	in := getInput("input.txt")

	resultOne := partOne(in)
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(in)
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

func partOne(ins []input) int {
	var count int
	for _, in := range ins {
		for _, out := range in.output {
			if len(out) != 5 && len(out) != 6 {
				count++
			}
		}
	}
	return count
}

func partTwo(ins []input) int {
	var total, num int
	for _, in := range ins {
		num = 0
		pattern := getPattern(in.signals)
		for i, out := range in.output {
			num += int(math.Pow10(3-i)) * pattern[out]
		}
		total += num
	}
	return total
}

func getPattern(signals []string) map[string]int {
	pattern := make(map[int]string)
	unknown := make([]string, 0, 6)
	// easy patterns first
	for _, s := range signals {
		switch len(s) {
		case 2:
			pattern[1] = s
		case 3:
			pattern[7] = s
		case 4:
			pattern[4] = s
		case 7:
			pattern[8] = s
		default:
			unknown = append(unknown, s)
		}
	}

	// of the unkown signals, we can figure out 3 by looking at signals with
	// length 5 and segments A & C(number 1)
	for i, signal := range unknown {
		if len(signal) == 5 && complement(pattern[1], signal) == "" {
			pattern[3] = signal
			unknown = append(unknown[:i], unknown[i+1:]...)
			break
		}
	}

	// 3 complement 7 give us segments D/G
	sDG := complement(pattern[3], pattern[7])

	// knowing segments D/G we can find 0
	for i, signal := range unknown {
		comp := complement(sDG, signal)
		if len(signal) == 6 && len(comp) == 1 {
			pattern[0] = signal
			unknown = append(unknown[:i], unknown[i+1:]...)
			break
		}
	}

	// we can figure out 6/9 comparing against A & C(number 1)
	for _, signal := range unknown {
		if len(signal) == 6 {
			if complement(pattern[1], signal) == "" {
				pattern[9] = signal
			} else {
				pattern[6] = signal
			}
		}
	}

	// finally we can get 2/5 comparing with 6
	for _, signal := range unknown {
		if len(signal) == 5 {
			if len(complement(pattern[6], signal)) == 1 {
				pattern[5] = signal
			} else {
				pattern[2] = signal
			}
		}
	}

	output := make(map[string]int, len(pattern))
	for k, v := range pattern {
		output[v] = k
	}
	return output
}

// gets everything that is in A except for anything in its overlap with B
// info: https://www.purplemath.com/modules/venndiag2.htm
func complement(a, b string) string {
	chars := make(map[string]struct{}, 6)
	for _, c := range a {
		chars[string(c)] = struct{}{}
	}
	for _, c := range b {
		delete(chars, string(c))
	}
	var comp string
	for c := range chars {
		comp += c
	}
	return comp
}

func getInput(path string) []input {
	var in []input

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	v := make([]string, 14)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&v[0], &v[1], &v[2], &v[3], &v[4], &v[5], &v[6], &v[7], &v[8], &v[9], &v[10], &v[11], &v[12], &v[13])
		for i := range v {
			s := strings.Split(v[i], "")
			sort.Strings(s)
			v[i] = strings.Join(s, "")
		}
		i := input{
			signals: make([]string, 10),
			output:  make([]string, 4),
		}
		copy(i.signals, v[:10])
		copy(i.output, v[10:])
		in = append(in, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return in
}
