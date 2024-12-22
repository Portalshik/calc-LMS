package calculator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	if expression == "" {
		return 0.0, errors.New("empty expression")
	}
	if containsLetters(expression) {
		return 0.0, errors.New("invalid expression")
	}

	for strings.Contains(expression, "(") {
		r, err := regexp.Compile(`\([^()]*\)`)
		if err != nil {
			return 0.0, err
		}

		matches := r.FindAllString(expression, -1)
		if len(matches) == 0 {
			return 0.0, errors.New("mismatched parentheses")
		}

		for _, match := range matches {
			inner := match[1 : len(match)-1]
			res, err := Calc(inner)
			if err != nil {
				return 0.0, err
			}
			expression = strings.Replace(expression, match, fmt.Sprintf("%f", res), 1)
		}
	}

	return evaluate(expression)
}

func evaluate(expression string) (float64, error) {
	r := regexp.MustCompile(`[-+*/()]|\d*\.?\d+`)
	tokens := r.FindAllString(expression, -1)

	if len(tokens) == 0 {
		return 0.0, errors.New("invalid expression")
	}

	rpn, err := toRPN(tokens)
	if err != nil {
		return 0.0, err
	}

	return evaluateRPN(rpn)
}

func toRPN(tokens []string) ([]string, error) {
	precedence := map[string]int{ "+": 1, "-": 1, "*": 2, "/": 2 }
	var output []string
	var operators []string
	unary := true

	for _, token := range tokens {
		switch {
		case isNumber(token):
			output = append(output, token)
			unary = false
		case token == "(":
			operators = append(operators, token)
			unary = true
		case token == ")":
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, errors.New("mismatched parentheses")
			}
			operators = operators[:len(operators)-1]
			unary = false
		case precedence[token] > 0:
			if unary {
				output = append(output, "0")
			}
			for len(operators) > 0 && precedence[operators[len(operators)-1]] >= precedence[token] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
			unary = true
		default:
			return nil, fmt.Errorf("invalid token: %s", token)
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, errors.New("mismatched parentheses")
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

func evaluateRPN(tokens []string) (float64, error) {
	var stack []float64

	for _, token := range tokens {
		if isNumber(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0.0, errors.New("invalid expression")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					panic("division by zero")
				}
				result = a / b
			default:
				panic(fmt.Sprintf("invalid operator: %s", token))
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0.0, errors.New("invalid expression")
	}

	return stack[0], nil
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func containsLetters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}