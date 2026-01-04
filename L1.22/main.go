package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	a.SetString("1000000000000000000000000", 10)
	b.SetString("10000000000000000000000", 10)
	ans := new(big.Int)

	fmt.Println("Sum = ", ans.Add(a, b))

	fmt.Println("Mul = ", ans.Mul(a, b))

	fmt.Println("Div = ", ans.Div(a, b))

	fmt.Println("Sub = ", ans.Sub(a, b))
}
