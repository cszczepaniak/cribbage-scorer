package main

import (
	"fmt"

	"github.com/cszczepaniak/cribbage-scorer/comb"
)

func main() {
	fmt.Println(comb.Factorial(5))
	fmt.Println(comb.Factorial(6))
	fmt.Println(comb.Factorial(7))
	fmt.Println(comb.Nchoosek(24, 6))
}
