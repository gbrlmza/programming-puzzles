package main

const (
	boardSize = 5
)

type (
	cell struct {
		X      int
		Y      int
		Marked bool
	}
	board struct {
		Cells map[int]cell
		Marks map[int]int
		Won   bool
	}
	game struct {
		Numbers []int
		Boards  []board
	}
)

func (b *board) mark(num int) bool {
	if b.Won { // already won
		return true
	}
	c, ok := b.Cells[num]
	if !ok { // missed number
		return false
	}
	if c.Marked { // already marked
		return false
	}
	c.Marked = true
	vkey := c.Y * -1
	hKey := c.X
	b.Marks[vkey]++
	b.Marks[hKey]++
	b.Cells[num] = c
	if b.Marks[vkey] >= boardSize || b.Marks[hKey] >= boardSize {
		b.Won = true
	}
	return b.Won
}

func (b *board) score() int {
	var score int
	for n, c := range b.Cells {
		if c.Marked {
			continue
		}
		score += n
	}
	return score
}
