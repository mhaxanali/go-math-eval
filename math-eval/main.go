package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mhaxanali/go-math-eval/math-eval/utils"
)

func main() {
	var evaluated float64

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter Expression: ")
	expr, _ := reader.ReadString('\n')
	expr = strings.TrimSpace(expr)

	result, err := utils.Evaluate(expr)

	if err != nil {
		fmt.Print("Error: ")
		fmt.Print(err)
		return
	}

	for _, v := range result {
		evaluated, err = strconv.ParseFloat(v, 64)
		if err != nil {
		}
	}

	fmt.Printf("Answer: %.3f", evaluated)

}
