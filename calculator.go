// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Add takes some numbers and returns the result of adding them together.
func Add(inputs ...float64) (float64, error) {
	res := inputs[0]

	if len(inputs) < 2 {
		return 0, fmt.Errorf("bad input: %f (only one operand)", res)
	}

	for _, n := range inputs[1:] {
		res += n
	}
	return res, nil
}

// Subtract takes some numbers and returns the result of subtracting them
// together.
func Substract(inputs ...float64) (float64, error) {
	res := inputs[0]

	if len(inputs) < 2 {
		return 0, fmt.Errorf("bad input: %f (only one operand)", res)
	}

	for _, n := range inputs[1:] {
		res -= n
	}
	return res, nil
}

// Multiply takes some numbers and returns the result of multiplying them
// together.
func Multiply(inputs ...float64) (float64, error) {
	res := inputs[0]

	if len(inputs) < 2 {
		return 0, fmt.Errorf("bad input: %f (only one operand)", res)
	}

	for _, n := range inputs[1:] {
		res *= n
	}
	return res, nil
}

// Divide takes two numbers and returns the result of dividing them
// together.
func Divide(inputs ...float64) (float64, error) {
	res := inputs[0]

	if len(inputs) < 2 {
		return 0, fmt.Errorf("bad input: %f (only one operand)", res)
	}

	for _, n := range inputs[1:] {
		if n == 0 {
			return 0, fmt.Errorf("bad input: %f, %f (division by zero is not allowed)", res, n)
		}
		res /= n
	}
	return res, nil
}

// Sqrt takes a number and returns its square root.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input: %f (square root of a negative number is not allowed)", a)
	}
	return math.Sqrt(a), nil
}

func Calculate(s string) (float64, error) {
	// Remove whitespaces (simply)
	s = strings.Replace(s, " ", "", -1)

	// Split between + - * /
	var symbol rune
	f := func(c rune) bool {
		symbols := []rune{'+', '-', '*', '/'}
		for _, r := range symbols {
			if c == r {
				symbol = r
				return true
			}
		}
		return false
	}

	fields := strings.FieldsFunc(s, f)
	if len(fields) == 1 {
		return 0, fmt.Errorf("bad input: %s (unrecognized operation)", s)
	}

	a, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, fmt.Errorf("bad input: %s", fields[0])
	}
	b, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return 0, fmt.Errorf("bad input: %s", fields[1])
	}

	var res float64
	switch symbol {
	case '+':
		res, err = Add(a, b)
	case '-':
		res, err = Substract(a, b)
	case '*':
		res, err = Multiply(a, b)
	case '/':
		res, err = Divide(a, b)
	}

	return res, err
}
