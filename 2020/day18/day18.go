package day18

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jdhenke/advent-of-code/input"
)

/*
Part1 Prompt

--- Day 18: Operation Order ---
As you look out the window and notice a heavily-forested continent slowly
appear over the horizon, you are interrupted by the child sitting next to you.
They're curious if you could help them with their math homework.

Unfortunately, it seems like this "math" follows different rules than you
remember.

The homework (your puzzle input) consists of a series of expressions that
consist of addition (+), multiplication (*), and parentheses ((...)). Just like
normal math, parentheses indicate that the expression inside must be evaluated
before it can be used by the surrounding expression. Addition still finds the
sum of the numbers on both sides of the operator, and multiplication still
finds the product.

However, the rules of operator precedence have changed. Rather than evaluating
multiplication before addition, the operators have the same precedence, and are
evaluated left-to-right regardless of the order in which they appear.

For example, the steps to evaluate the expression 1 + 2 * 3 + 4 * 5 + 6 are as
follows:

    1 + 2 * 3 + 4 * 5 + 6
      3   * 3 + 4 * 5 + 6
          9   + 4 * 5 + 6
             13   * 5 + 6
                 65   + 6
                     71

Parentheses can override this order; for example, here is what happens if
parentheses are added to form 1 + (2 * 3) + (4 * (5 + 6)):

    1 + (2 * 3) + (4 * (5 + 6))
    1 +    6    + (4 * (5 + 6))
         7      + (4 * (5 + 6))
         7      + (4 *   11   )
         7      +     44
                51

Here are a few more examples:

- 2 * 3 + (4 * 5) becomes 26.
- 5 + (8 * 3 + 9 + 3 * 4 * 3) becomes 437.
- 5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4)) becomes 12240.
- ((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 becomes 13632.

Before you can help with the homework, you need to understand it yourself.
Evaluate the expression on each line of the homework; what is the sum of the
resulting values?
*/
func Part1(r io.Reader) (answer int, err error) {
	return day18(r)
}

func Part2(r io.Reader) (answer int, err error) {
	return day18(r)
}

func day18(r io.Reader) (answer int, err error) {
	if err := input.ForEachLine(r, func(line string) error {
		e, err := Parse(line)
		if err != nil {
			return err
		}
		answer += eval(e)
		return nil
	}); err != nil {
		return 0, err
	}
	return answer, nil
}

func eval(e *Expression) int {
	if e.Value != nil {
		return *e.Value
	}
	current := eval(e.Expressions[0])
	for i := 0; i < len(e.Operators); i++ {
		if e.Operators[i] == "+" {
			current += eval(e.Expressions[i+1])
		} else {
			current *= eval(e.Expressions[i+1])
		}
	}
	return current
}

type Expression struct {
	Value       *int
	Expressions []*Expression
	Operators   []string
}

func Parse(line string) (*Expression, error) {
	scan := bufio.NewScanner(strings.NewReader(line))
	scan.Split(splitFunc)
	return parse(scan, false)
}

// lexer acts as an untyped lexer, each token should be a paren, a value, or an operator.
func parse(scan *bufio.Scanner, inParens bool) (*Expression, error) {
	current := &Expression{}
	for scan.Scan() {
		switch t := scan.Text(); t {
		case "(":
			sub, err := parse(scan, true)
			if err != nil {
				return nil, err
			}
			if len(current.Expressions) != len(current.Operators) {
				return nil, fmt.Errorf("unexpected expression: %v", sub)
			}
			current.Expressions = append(current.Expressions, sub)
		case ")":
			if !inParens {
				return nil, fmt.Errorf("unexpected closing paren")
			}
			return current, nil
		case "+", "*":
			if len(current.Expressions) != len(current.Operators)+1 {
				return nil, fmt.Errorf("unexpected operator: %v", t)
			}
			current.Operators = append(current.Operators, t)
		default: // value
			if len(current.Expressions) != len(current.Operators) {
				return nil, fmt.Errorf("unexpected value: %v", t)
			}
			d, err := strconv.Atoi(t)
			if err != nil {
				return nil, err
			}
			current.Expressions = append(current.Expressions, &Expression{
				Value: &d,
			})
		}
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	if inParens {
		return nil, fmt.Errorf("detected unclosed paren")
	}
	if len(current.Expressions) != len(current.Operators)+1 {
		return nil, fmt.Errorf("unfinished expression")
	}
	return current, nil
}

var _ bufio.SplitFunc = splitFunc

var digits = map[byte]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
}

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// because: https://stackoverflow.com/a/19941087, it seems the best thing to do is always keep data[0] the start of
	// a new token, so always take up any trailing whitespace.
	defer func() {
		if err != nil {
			return
		}
		if advance < len(data) && data[advance] == ' ' {
			advance++
		}
	}()
	if len(data) == 0 {
		return 0, nil, nil
	}
	switch c := data[0]; c {
	case '(', ')', '+', '*':
		return 1, data[0:1], nil
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		i := 0
		for ; i < len(data) && digits[data[i]]; i++ {
		}
		if i == len(data) && !atEOF {
			return 0, nil, nil
		}
		return i, data[:i], nil
	default:
		return 0, nil, fmt.Errorf("unexpected character: %s", string(c))
	}

}
