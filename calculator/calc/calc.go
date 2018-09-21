package main

import (
	"fmt"
	"os"
	"strconv"

	"../func"
)

func main() {
	args := os.Args
	x := strconv.ParseFloat(os.Args[1], 32)
	y := strconv.ParseFloat(os.Args[3], 32)

	switch op := os.Args[2]; op {
	case "+":
		fmt.Print(ident.Add(x, y))
	case "-":
		fmt.Print(ident.Subs(x, y))
	case "*":
		fmt.Print(ident.Mult(x, y))
	case "/":
		fmt.Print(ident.Div(x, y))
	default:
		fmt.Println("no operator")
	}
}
