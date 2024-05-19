package stringify

import (
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func GenerateTokensWithSpaces(str string) []string {
	tokens := strings.Split(str, " ")
	return tokens
}

func GenerateTokens(str string) []string {
	tokens := []string{}
	for i := 0; i < len(str); {
		if strings.Contains("0123456789", string(str[i])) {
			token := ""
			for strings.Contains("0123456789", string(str[i])) {
				token += string(str[i])
				i++
				if i+1 >= len(str) {
					break
				}
			}
			tokens = append(tokens, token)
		} else if strings.Contains("+-*/", string(str[i])) {
			token := ""
			for strings.Contains("+-*/", string(str[i])) {
				token += string(str[i])
				i++
				if i+1 >= len(str) {
					break
				}
			}
			if token == "**" {
				token = "^"
			}
			tokens = append(tokens, token)
		}
	}
	return tokens
}

//TODO implement GenerateTokensWithBrackets

func SimpleCalculateWithSpaces(exp string) int {
	if len(exp) == 0 {
		return -1
	}
	stack := &Stack{}
	tokens := GenerateTokensWithSpaces(exp)
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "+" || token == "-" || token == "*" || token == "/" || token == "**" || token == "^" {
			op1 := stack.Pop()
			op2, _ := strconv.Atoi(tokens[i+1])
			i++
			var res int
			switch token {
			case "+":
				res = op1 + op2
			case "-":
				res = op1 - op2
			case "*":
				res = op1 * op2
			case "/":
				res = op1 / op2
			case "**":
				res = op1 ^ op2
			case "^":
				res = op1 ^ op2
			}
			stack.Push(res)
		} else {
			num, _ := strconv.Atoi(token)
			stack.Push(num)
		}
	}

	res := stack.Pop()
	return res
}

func SimpleCalculate(exp string) int {
	if len(exp) == 0 {
		return -1
	}
	stack := &Stack{}
	tokens := GenerateTokens(exp)
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "+" || token == "-" || token == "*" || token == "/" || token == "**" || token == "^" {
			op1 := stack.Pop()
			op2, _ := strconv.Atoi(tokens[i+1])
			i++
			var res int
			switch token {
			case "+":
				res = op1 + op2
			case "-":
				res = op1 - op2
			case "*":
				res = op1 * op2
			case "/":
				res = op1 / op2
			case "**":
				res = op1 ^ op2
			case "^":
				res = op1 ^ op2
			}
			stack.Push(res)
		} else {
			num, _ := strconv.Atoi(token)
			stack.Push(num)
		}
	}

	res := stack.Pop()
	return res
}
