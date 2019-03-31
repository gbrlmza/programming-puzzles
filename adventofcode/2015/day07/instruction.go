package main

const (
	Literal = "literal"
	And     = "and"
	Not     = "not"
	Or      = "or"
	Rshifth = "rshifth"
	Lshifth = "lshifth"
)

type Instruction struct {
	Operator string
	Sources  []int
	Output   int
	Target   string
}
