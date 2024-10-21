package main

import (
	"fmt"
	"slices"
	"strings"
)

type Token struct {
	str string
}

func main() {
	Calc("123+21-1.1/3")
}

func Calc(expression string) (float64, error) {
	expList := strings.Split(expression, "")
	res := []string{}
	temp := ""
	// _type := ""
	for _, elem := range expList {
		if isNumber(elem) {
			temp += elem
		} else if isDot(elem) {
			temp += elem
			// _type = "float64"
		} else if isAct(elem) {
			res = append(res, temp)
			temp = ""
			res = append(res, elem)
		}
	}
	res = append(res, temp)
	fmt.Println(res)
	return 0, nil
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
	return slices.Contains([]rune("+-*/()"), runes)
}
