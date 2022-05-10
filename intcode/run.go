package intcode

import (
	"fmt"
)

type Code []int

func Run(code Code) (ans int) {
	ip := 0
	for {
		switch op := code[ip]; op {
		case 1:
			code[code[ip+3]] = code[code[ip+1]] + code[code[ip+2]]
			ip += 4
		case 2:
			code[code[ip+3]] = code[code[ip+1]] * code[code[ip+2]]
			ip += 4
		case 99:
			return code[0]
		default:
			panic(fmt.Sprintf("unknown opcode: %v", op))
		}
	}
}
