package utils

import (
	"fmt"
	"strings"
)

var opers [7]rune = [7]rune{'+', '-', '*', '/', '^', '(', ')'}

func Evaluate(expr string) (string, error) {

	var err error = nil

	expr_tokens := tokenize(expr)
	// only for debugging purposes
	fmt.Print(expr_tokens)

	return "", err
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
