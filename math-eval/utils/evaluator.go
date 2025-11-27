package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var opers [7]rune = [7]rune{'+', '-', '*', '/', '^', '(', ')'}

func Evaluate(expr string) ([]string, error) {

	var err error = nil

	expr_tokens := tokenize(expr)

	if !(contains(expr_tokens, "(") || contains(expr_tokens, ")")) {
		result, err := solveTokens(expr_tokens)
		return result, err
	}

	return []string{}, err
}

func tokenize(expr string) []string {

	for _, oper := range opers {
		expr = strings.ReplaceAll(expr, string(oper), fmt.Sprintf(",%c,", oper))
	}
	expr_tokens := strings.Split(expr, ",")

	return expr_tokens

}

func contains(comp_slice []string, comp_str string) bool {
	for _, s := range comp_slice {
		if s == comp_str {
			return true
		}
	}
	return false
}

func solveTokens(tokens []string) ([]string, error) {
	var err error
	for contains(tokens, "^") {
		for i, v := range tokens {
			if v == "^" {
				if i+2 >= len(tokens) {
					tokens, err = solve(tokens, i)
					if err != nil {
						return []string{}, err
					}
					break

				} else if tokens[i+2] == "^" {

				} else {
					tokens, err = solve(tokens, i)
					if err != nil {
						return []string{}, err
					}
					break
				}
			}
		}
	}
	for contains(tokens, "/") || contains(tokens, "*") {
		for i, v := range tokens {
			if v == "*" {
				tokens, err = solve(tokens, i)
				if err != nil {
					return []string{}, err
				}
				break
			} else if v == "/" {
				tokens, err = solve(tokens, i)
				if err != nil {
					return []string{}, err
				}
				break
			}
		}
	}
	for contains(tokens, "+") || contains(tokens, "-") {
		for i, v := range tokens {
			if v == "+" {
				tokens, err = solve(tokens, i)
				if err != nil {
					return []string{}, err
				}
				break
			} else if v == "-" {
				tokens, err = solve(tokens, i)
				if err != nil {
					return []string{}, err
				}
				break
			}
		}
	}

	return tokens, nil
}

func toFloat64(s1 string, s2 string) (float64, float64, error) {
	f1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		return 1, 1, errors.New("invalid characters passed")
	}
	f2, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		return 1, 1, errors.New("invalid characters passed")
	}
	return f1, f2, nil
}

func solve(tokens []string, i int) ([]string, error) {
	var tmp float64
	hlp1, hlp2, err := toFloat64(tokens[i-1], tokens[i+1])
	if err != nil {
		return []string{}, errors.New("invalid characters passed")
	}
	switch tokens[i] {
	case "^":
		tmp = math.Pow(hlp1, hlp2)
	case "/":
		if hlp2 == 0 {
			return []string{}, errors.New("cannot divide by zero")
		} else {
			tmp = hlp1 / hlp2
		}
	case "*":
		tmp = hlp1 * hlp2
	case "+":
		tmp = hlp1 + hlp2
	case "-":
		tmp = hlp1 - hlp2
	}

	tokens = append(tokens[:i-1], append([]string{fmt.Sprintf("%.2f", tmp)}, tokens[i+2:]...)...)
	return tokens, err
}
