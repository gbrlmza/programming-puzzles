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
type table struct {
	rows      map[string]row
	unvisited map[string]int
}
type row struct {
	node     string
	distance int
	previous string
	visited  bool
}

func main() {
	resultOne := partOne(getInput("input.txt", 0))
	fmt.Printf("PartOne: %v\n", resultOne)

	resultTwo := partTwo(getInput("input.txt", 4))
	fmt.Printf("PartTwo: %v\n", resultTwo)
}

// part one solution
func partOne(g graph) int {
	return shortestPath(g)
}

// part two solution
func partTwo(g graph) int {
	return shortestPath(g)
}

// Shortest path. Dijkstra's algorithm:
// - https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
// - https://www.youtube.com/watch?v=pVfj6mxhdMw
func shortestPath(g graph) int {
	tbl := netTable()
	tbl.set(row{node: getNodeName(1, 1)})
	visited := len(g.nodes)

	for visited != 0 {
		row := tbl.nextNode()
		row.visited = true
		currNode := g.nodes[row.node]
		for _, connName := range currNode.connections {
			conn := g.nodes[connName]
			connRow := tbl.rows[connName]
			if connRow.visited {
				continue
			}
			newDist := row.distance + conn.risk
			if connRow.node == "" || newDist < connRow.distance {
				connRow.node = connName
				connRow.distance = newDist
				connRow.previous = currNode.name
				tbl.set(connRow)
			}
		}
		tbl.set(row)
		visited--
	}

	return tbl.rows[g.end].distance
}

func (t table) nextNode() row {
	// TODO: This can be improved. I could maintain an ordered list of unvisited
	// nodes to avoid iterating over all the unvisited nodes each time.
	var nextName string
	var nextDist int
	for name, dist := range t.unvisited {
		if nextName == "" || dist < nextDist {
			nextName = name
			nextDist = dist
		}
	}
	return t.rows[nextName]
}

func (t table) set(r row) {
	if r.visited {
		delete(t.unvisited, r.node)
	} else {
		t.unvisited[r.node] = r.distance
	}
	t.rows[r.node] = r
}

func netTable() table {
	return table{
		rows:      make(map[string]row),
		unvisited: make(map[string]int),
	}
}

func getInput(path string, expand int) graph {
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
			nodes := getNodes(x+1, y, risk, len(values), len(values), expand)
			for _, n := range nodes {
				output.nodes[n.name] = n
				output.end = n.name
			}
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

func getNodes(x, y, risk, width, height, expand int) []node {
	var nodes []node
	nodes = append(nodes, node{
		name:        getNodeName(x, y),
		risk:        risk,
		connections: []string{getNodeName(x-1, y), getNodeName(x+1, y), getNodeName(x, y-1), getNodeName(x, y+1)},
	})
	// horizontal expansion
	exp := [][]int{{x, y, risk}}
	for i := 1; i <= expand; i++ {
		erisk := newRisk(risk, i)
		ex := width*i + x
		ey := y
		exp = append(exp, []int{ex, ey, erisk})
		nodes = append(nodes, node{
			name:        getNodeName(ex, ey),
			risk:        erisk,
			connections: []string{getNodeName(ex-1, ey), getNodeName(ex+1, ey), getNodeName(ex, ey-1), getNodeName(ex, ey+1)},
		})
	}
	// vertical expansion
	for _, h := range exp {
		for i := 1; i <= expand; i++ {
			ex := h[0]
			ey := height*i + h[1]
			erisk := newRisk(h[2], i)
			nodes = append(nodes, node{
				name:        getNodeName(ex, ey),
				risk:        erisk,
				connections: []string{getNodeName(ex-1, ey), getNodeName(ex+1, ey), getNodeName(ex, ey-1), getNodeName(ex, ey+1)},
			})
		}
	}
	return nodes
}

func newRisk(risk, inc int) int {
	n := risk + inc
	if n > 9 {
		n = n % 9
	}
	return n
}
