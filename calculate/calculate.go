package calculate

import (
	"errors"
	"strconv"
	"strings"
)

func applyOp(op rune, b, a float64) (float64, error) {
	switch op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	}
	return 0, errors.New("unknown operator")
}

func Calc(expression string) (float64, error) {

	expression = strings.ReplaceAll(expression, " ", "")

	numStack := []float64{}
	opStack := []rune{}

	priority := func(op rune) int {
		if op == '+' || op == '-' {
			return 1
		}
		if op == '*' || op == '/' {
			return 2
		}
		return 0
	}

	var err error

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])
		if ch >= '0' && ch <= '9' {
			j := i
			for j < len(expression) && ((expression[j] >= '0' && expression[j] <= '9') || expression[j] == '.') {
				j++
			}
			num, err := strconv.ParseFloat(expression[i:j], 64)
			if err != nil {
				return 0, errors.New("invalid number format")
			}
			numStack = append(numStack, num)
			i = j - 1
		} else if ch == '(' {
			opStack = append(opStack, ch)
		} else if ch == ')' {

			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				numStack, opStack, err = processTopOperation(numStack, opStack)
				if err != nil {
					return 0, err
				}
			}
			if len(opStack) == 0 {
				return 0, errors.New("mismatched parentheses")
			}

			opStack = opStack[:len(opStack)-1]
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {

			for len(opStack) > 0 && priority(opStack[len(opStack)-1]) >= priority(ch) {
				numStack, opStack, err = processTopOperation(numStack, opStack)
				if err != nil {
					return 0, err
				}
			}
			opStack = append(opStack, ch)
		} else {
			return 0, errors.New("invalid character in expression")
		}
	}

	for len(opStack) > 0 {
		numStack, opStack, err = processTopOperation(numStack, opStack)
		if err != nil {
			return 0, err
		}
	}

	if len(numStack) != 1 {
		return 0, errors.New("invalid expression")
	}
	return numStack[0], nil
}

func processTopOperation(numStack []float64, opStack []rune) ([]float64, []rune, error) {
	if len(numStack) < 2 {
		return numStack, opStack, errors.New("not enough operands")
	}
	if len(opStack) == 0 {
		return numStack, opStack, errors.New("operator stack is empty")
	}

	b := numStack[len(numStack)-1]
	a := numStack[len(numStack)-2]
	op := opStack[len(opStack)-1]

	numStack = numStack[:len(numStack)-2]
	opStack = opStack[:len(opStack)-1]

	result, err := applyOp(op, b, a)
	if err != nil {
		return numStack, opStack, err
	}

	numStack = append(numStack, result)
	return numStack, opStack, nil
}
