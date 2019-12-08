package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Opcode int

const (
	Add  Opcode = 1
	Mul         = 2
	Halt        = 99
)

type Program struct {
	data []int
	pos  int
}

func NewProgram(data []int) Program {
	prog := Program{}
	prog.data = make([]int, len(data))
	copy(prog.data, data)
	return prog
}

func (p *Program) Output() int {
	p.run()
	return p.fetch(0)
}

func (p *Program) run() {
	for op := Opcode(p.peek(0)); op != Halt; op = Opcode(p.peek(0)) {
		pos1 := p.peek(1)
		pos2 := p.peek(2)
		writePos := p.peek(3)

		operand1, operand2 := p.fetch(pos1), p.fetch(pos2)
		var result int
		switch op {
		case Add:
			result = operand1 + operand2
		case Mul:
			result = operand1 * operand2
		default:
			panic(fmt.Sprintf("unknown opcode %v", op))
		}
		p.write(writePos, result)

		p.step(4)
	}
}

func (p *Program) peek(howfar int) int {
	return p.data[p.pos+howfar]
}

func (p *Program) step(size int) {
	p.pos += size
}

func (p *Program) fetch(pos int) int {
	return p.data[pos]
}

func (p *Program) write(pos, val int) {
	p.data[pos] = val
}

func day2ProcessedInput() []int {
	input := readInput("./day2.txt")
	input = strings.Trim(input, "\n")
	fields := strings.Split(input, ",")
	var data []int
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		data = append(data, num)
	}
	return data
}

func Day2Part1() int {
	data := day2ProcessedInput()

	prog := NewProgram(data)
	// Before running the program, replace position 1 with the value 12,
	// and replace position 2 with the value 2
	prog.write(1, 12)
	prog.write(2, 2)
	// What value is left at position 0 after the program halts?
	return prog.Output()
}

func Day2Part2() int {
	data := day2ProcessedInput()

	// data[1] = noun, data[2] = verb
	// Find the input noun and verb that cause the program to produce the output 19690720
	requiredOutput := 19690720
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			prog := NewProgram(data)
			prog.write(1, noun)
			prog.write(2, verb)

			if prog.Output() == requiredOutput {
				// What is 100 * noun + verb?
				return 100*noun + verb
			}
		}
	}

	panic("no answer")
}
