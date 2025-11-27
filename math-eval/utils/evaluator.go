package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var opers [7]rune = [7]rune{'+', '-', '*', '/', '^', '(', ')'}

func count(col []string, com string) int {
	var count int
	for _, v := range col {
		if v == com {
			count++
		}
	}
	return count
}

func Evaluate(expr string) ([]string, error) {
	var err error
	var result []string

	expr = strings.ReplaceAll(expr, " ", "")

	expr_tokens := tokenize(expr)
	for {
		if !(contains(expr_tokens, "(") || contains(expr_tokens, ")")) {
			result, err = solveTokens(expr_tokens)
			return result, err
		} else {
			if count(expr_tokens, "(") != count(expr_tokens, ")") {
				return []string{}, errors.New("mismatched brackets")
			} else {
				expr_tokens, err = solveInnerMostBrackets(expr_tokens)
				if err != nil {
					return []string{}, err
				}
			}
		}
	}
}

func tokenize(expr string) []string {
	for _, oper := range opers {
		expr = strings.ReplaceAll(expr, string(oper), fmt.Sprintf(",%c,", oper))
	}
	rawTokens := strings.Split(expr, ",")
	var expr_tokens []string
	for _, t := range rawTokens {
		t = strings.TrimSpace(t) // remove spaces
		if t != "" {             // skip empty tokens
			expr_tokens = append(expr_tokens, t)
		}
	}
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

func solveInnerMostBrackets(tokens []string) ([]string, error) {
	var open int
	var close int

	for i, v := range tokens {
		if v == "(" {
			open = i
		} else if v == ")" {
			close = i
			break
		}
	}
	inner := tokens[open+1 : close]
	solved, err := solveTokens(inner)

	if err != nil {
		return []string{}, err
	}

	newTokens := append(tokens[:open], append(solved, tokens[close+1:]...)...)

	return newTokens, nil
}
