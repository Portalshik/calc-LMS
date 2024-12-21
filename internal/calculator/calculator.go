package calculator

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	r, err := regexp.Compile(`\([^()]*\)|\(([^()]*\([^()]*\)[^()]*)+\)`)
	if err != nil {
		return 0.0, err
	}

	matches := r.FindAllString(expression, -1)
	expression = r.ReplaceAllLiteralString(expression, "%f")
	matchesResults := []any{}
	for _, exp := range matches {
		resT, _ := Calc(exp[1 : len([]rune(exp))-1])
		matchesResults = append(matchesResults, resT)
	}
	expression = fmt.Sprintf(expression, matchesResults...)

	expList := strings.Split(expression, "")
	res := []string{}
	temp := ""
	for i, elem := range expList {
		if isNumber(elem) {
			temp += elem
		} else if isDot(elem) {
			temp += elem
		} else if isAct(elem) {
			if i == len(expList)-1 || isAct(expList[i-1]) {
				return 0.0, fmt.Errorf("invalid expression")
			}
			res = append(res, temp)
			temp = ""
			res = append(res, elem)
		}
	}
	res = append(res, temp)
	return _calculate(res)
}

func isNumber(symbol string) bool {
	runes := []rune(symbol)[0]
	return runes <= 57 && runes >= 48
}

func isDot(symbol string) bool {
	runes := []rune(symbol)[0]
	return runes == 46
}

func isAct(symbol string) bool {
	runes := []rune(symbol)[0]
	return slices.Contains([]rune("+-*/"), runes)
}

func _calculate(exp []string) (float64, error) {
	for _, val := range "*/+-" {
		for i := 0; i < len(exp); i++ {
			if exp[i] == string(val) {
				if i-1 < 0 || i+1 >= len(exp) {
					return 0.0, fmt.Errorf("invalid expression")
				}
				a, _ := strconv.ParseFloat(exp[i-1], 64)
				b, _ := strconv.ParseFloat(exp[i+1], 64)
				res := _act(a, b, string(val))
				var tempExp []string
				tempExp = append(tempExp, exp[:i-1]...)
				tempExp = append(tempExp, fmt.Sprintf("%f", res))
				tempExp = append(tempExp, exp[i+2:]...)
				exp = tempExp
				i--
			}
		}
	}

	return strconv.ParseFloat(exp[0], 64)
}

func _act(a, b float64, act string) float64 {
	switch act {
	case "/":
		return a / b
	case "*":
		return a * b
	case "+":
		return a + b
	case "-":
		return a - b
	}
	return 0.0
}
