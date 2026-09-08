package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrDivByZero = errors.New("division by zero")

func EvalPostfix(tokens []string) (int, error) {
	stack := make([]int, 0, len(tokens))

	for _, tok := range tokens {
		switch tok {
		case "+", "-", "*", "/":
			length := len(stack)
			if length < 2 {
				return 0, fmt.Errorf("not enough operands for operator %q", tok)
			}

			b := stack[length-1]
			a := stack[length-2]
			stack = stack[:length-2]

			res, err := calculate(a, b, tok)
			if err != nil {
				return 0, err
			}
			stack = append(stack, res)

		default:
			n, err := strconv.Atoi(tok)
			if err != nil {
				return 0, fmt.Errorf("unknown token %q", tok)
			}
			stack = append(stack, n)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}

func calculate(a int, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivByZero
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator %q", op)
	}
}

func main() {}
