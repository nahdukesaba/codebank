package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Define operator precedence
var precedence = map[string]int{
	"+":  1,
	"-":  1,
	"*":  2,
	"/":  2,
	"**": 3,
}

// Helper function to determine if a token is an operator
func IsOperator(token string) bool {
	_, ok := precedence[token]
	return ok
}

// Shunting Yard Algorithm to convert infix to postfix
func InfixToPostfix(expression string) ([]string, error) {
	var output []string
	var operators []string

	tokens := strings.Fields(expression)

	for _, token := range tokens {
		if IsOperator(token) {
			for len(operators) > 0 && precedence[operators[len(operators)-1]] >= precedence[token] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		} else {
			output = append(output, token)
		}
	}

	for len(operators) > 0 {
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

// Function to evaluate a postfix expression
func EvaluatePostfix(postfix []string) (int, error) {
	var stack []int

	for _, token := range postfix {
		if IsOperator(token) {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				result = a / b
			}
			stack = append(stack, result)
		} else {
			value, err := strconv.Atoi(token)
			if err != nil {
				return 0, err
			}
			stack = append(stack, value)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}

// Function to evaluate an infix expression
func EvaluateExpression(expression string) (int, error) {
	postfix, err := InfixToPostfix(expression)
	if err != nil {
		return 0, err
	}
	return EvaluatePostfix(postfix)
}

//func main() {
//	expression := "2 + 1 * 3"
//	result, err := EvaluateExpression(expression)
//	if err != nil {
//		fmt.Println("Error evaluating expression:", err)
//		return
//	}
//	fmt.Printf("Result: %d\n", result)
//}
