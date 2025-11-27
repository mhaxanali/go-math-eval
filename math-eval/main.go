package main

import (
	"fmt"

	"github.com/mhaxanali/go-math-eval/math-eval/utils"
)

func main() {
	result, err := utils.Evaluate("0/0")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(result)
}
