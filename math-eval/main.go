package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mhaxanali/go-math-eval/math-eval/utils"
)

func main() {
	var evaluated float64
	var expr string

	args := os.Args

	if len(args) == 2 {
		expr = args[1]
	} else if len(args) == 1 {
		fmt.Print("Error: no expression provided")
		return
	} else {
		fmt.Printf("Error: Invalid number of arguments")
	}

	result, err := utils.Evaluate(expr)

	if err != nil {
		fmt.Print("Error: ")
		fmt.Print(err)
		return
	}

	evaluated, _ = strconv.ParseFloat(result[0], 64)

	fmt.Printf("Answer: %.3f", evaluated)

}
