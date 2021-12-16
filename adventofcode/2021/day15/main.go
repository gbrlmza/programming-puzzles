package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/15

type graph struct {
	nodes map[string]node
	start string
	end   string
}
type node struct {
	name        string
	risk        int
	connections []string
}
type table []row
type row struct {
	node     string
	distance int
	previous string
	visited  bool
}

func main() {
	input := getInput("input.txt")

	resultOne := partOne(input)
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(input)
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

// part one solution. Dijkstra's algorithm:
// - https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
// - https://www.youtube.com/watch?v=pVfj6mxhdMw
func partOne(g graph) int {
	tbl := table{{node: getNodeName(1, 1)}}

	var newDist int
	var currNode node
	for tbl.hasUnvisited() {
		row := tbl.nextNode()
		row.visited = true
		currNode = g.nodes[row.node]
		for _, connName := range currNode.connections {
			conn := g.nodes[connName]
			connRow := tbl.get(connName)
			if connRow.visited {
				continue
			}
			newDist = row.distance + conn.risk
			if connRow.node == "" || newDist < connRow.distance {
				connRow.node = connName
				connRow.distance = newDist
				connRow.previous = currNode.name
				tbl = tbl.set(connRow)
			}
		}
		tbl = tbl.set(row)
	}

	return tbl.get(g.end).distance
}

// part two solution
func partTwo(g graph) int {
	return 0
}

func (t table) nextNode() row {
	var next row
	for _, r := range t {
		if r.visited {
			continue
		}
		if next.node == "" || r.distance < next.distance {
			next = r
		}
	}
	return next
}

func (t table) get(name string) row {
	for _, r := range t {
		if r.node == name {
			return r
		}
	}
	return row{}
}

func (t table) set(newRow row) table {
	for i, r := range t {
		if r.node == newRow.node {
			t[i] = newRow
			return t
		}
	}
	t = append(t, newRow)
	return t
}

func (t table) hasUnvisited() bool {
	for _, r := range t {
		if !r.visited {
			return true
		}
	}
	return false
}

func getInput(path string) graph {
	output := graph{
		nodes: make(map[string]node),
		start: getNodeName(1, 1),
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 1
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "")
		for x, str := range values {
			risk, _ := strconv.Atoi(str)
			nodeName := getNodeName(x+1, y)
			output.nodes[nodeName] = node{
				name: nodeName,
				risk: risk,
				connections: []string{
					getNodeName(x, y),
					getNodeName(x+2, y),
					getNodeName(x+1, y-1),
					getNodeName(x+1, y+1),
				},
			}
			output.end = nodeName
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// removed non-valid node connections
	for nodeName, node := range output.nodes {
		for i := len(node.connections) - 1; i >= 0; i-- {
			connName := node.connections[i]
			if _, ok := output.nodes[connName]; !ok {
				node.connections = append(node.connections[:i], node.connections[i+1:]...)
			}
		}
		output.nodes[nodeName] = node
	}

	return output
}

func getNodeName(x, y int) string {
	return fmt.Sprintf("%d.%d", x, y)
}
