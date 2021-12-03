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
	Input    []int
	Output   int
	Target   string
}

// Execute instruction
func (i Instruction) Execute() (out int) {
	// if i.Operator == Literal {
	// 	out = i.Input[0]
	// } else if i.Operator == Not {
	// 	out = !i
	// }
	// return out
	return 0
}
