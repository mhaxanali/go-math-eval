package main

import (
	"fmt"

	"github.com/mhaxanali/go-math-eval/math-eval/utils"
)

func main() {
	str, err := utils.Evaluate("2+2")
	if err != nil {
		return
	}
	fmt.Print(str)
}
