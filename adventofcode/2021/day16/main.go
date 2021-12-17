package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

// https://adventofcode.com/2021/day/16

type input string
type packet struct {
	version     int
	ptype       int
	value       int
	lengthType  int
	lengthValue int
	subPackets  []packet
}

func main() {
	resultOne := partOne(getInput("input.txt"))
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(getInput("input.txt"))
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

// part one solution
func partOne(in input) int {
	_, sum := decode(&in, 0)
	return sum
}

// part two solution
func partTwo(in input) int {
	packets, _ := decode(&in, 0)
	return calculate(packets)
}

func decode(in *input, packetLimit int) ([]packet, int) {
	var sum, subSum int
	var packets []packet
	var p packet
	for binStrToDec(string(*in)) != 0 {
		if packetLimit != 0 && len(packets) >= packetLimit {
			break
		}
		if p.version == 0 {
			p.version = binStrToDec(in.consume(3))
			p.ptype = binStrToDec(in.consume(3))
			sum += p.version
			if p.ptype != 4 {
				p.lengthType = binStrToDec(in.consume(1))
				if p.lengthType == 0 {
					p.lengthValue = binStrToDec(in.consume(15))
				} else {
					p.lengthValue = binStrToDec(in.consume(11))
				}
			}
		}
		if p.ptype == 4 {
			var strValue string
			for {
				group := in.consume(5)
				strValue += group[1:]
				if group[0] == '0' {
					break
				}
			}
			mod := len(strValue) % 4
			if mod != 0 {
				in.consume(4 - mod)
			}
			p.value = binStrToDec(strValue)
		} else {
			if p.lengthType == 0 {
				subInput := input(in.consume(p.lengthValue))
				p.subPackets, subSum = decode(&subInput, 0)
			} else {
				p.subPackets, subSum = decode(in, p.lengthValue)
			}
			sum += subSum
		}
		packets = append(packets, p)
		p = packet{}
	}
	return packets, sum
}

func calculate(packets []packet) int {
	var total int
	for _, p := range packets {
		var partial int
		switch p.ptype {
		case 4: // literal value
			return p.value
		case 0: // sum
			for _, sp := range p.subPackets {
				partial += calculate([]packet{sp})
			}
		case 1: // product
			partial = 1
			for _, sp := range p.subPackets {
				partial *= calculate([]packet{sp})
			}
		case 2: // minimum
			values := packetValues(p.subPackets)
			sort.Ints(values)
			partial = values[0]
		case 3: // maximun
			values := packetValues(p.subPackets)
			sort.Ints(values)
			partial = values[len(values)-1]
		case 5: // greater than
			values := packetValues(p.subPackets)
			if values[0] > values[1] {
				return 1
			}
			return 0
		case 6: // less than
			values := packetValues(p.subPackets)
			if values[0] < values[1] {
				return 1
			}
			return 0
		case 7: // equal to
			values := packetValues(p.subPackets)
			if values[0] == values[1] {
				return 1
			}
			return 0
		}
		total += partial
	}
	return total
}

func packetValues(packets []packet) []int {
	var values []int
	for _, p := range packets {
		values = append(values, calculate([]packet{p}))
	}
	return values
}

func (i *input) consume(len int) string {
	str := *i
	val := str[:len]
	str = str[len:]
	*i = str
	return string(val)
}

func binStrToDec(binStr string) int {
	dec, _ := strconv.ParseInt(binStr, 2, 64)
	return int(dec)
}

func getInput(path string) input {
	var bin string
	bytes, _ := ioutil.ReadFile(path)
	for _, hex := range bytes {
		dec, _ := strconv.ParseUint(string(hex), 16, 64)
		bin += fmt.Sprintf("%04b", dec)
	}
	return input(bin)
}
