package main

import (
	"fmt"

	"github.com/steve-westwood/goTesting/util"
)

func main() {
	resultOfAddition := util.Add(1, 2)
	resultOfMultiplication := util.Multiple(5, 4)
	fmt.Printf("Result from addition (1 + 2): %s/n", resultOfAddition)
	fmt.Printf("Result from multiplication (5 x 4): %s/n", resultOfMultiplication)
}
