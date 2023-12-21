// Package ch10_solution provides utility functions to
// add two numbers together.
package ch10_solution

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two numbers as parameters and returns the sum
// of those two numbers, as demonstrated in https://www.mathsisfun.com/numbers/addition.html.
func Add[T Number](val1, val2 T) T {
	return val1 + val2
}
