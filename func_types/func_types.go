package main

import (
	"fmt"
)

func main() {
	check("asdasdad")
	check(2)
	check(2.4)
	check(true)
}

func check(i interface{}) {
	switch t := i.(type) {
	case float32:
		fmt.Printf("%T\n", t)
	case int:
		fmt.Printf("%T\n", t)
	case string:
		fmt.Printf("%T\n", t)
	case bool:
		fmt.Printf("%T\n", t)
	default:
		fmt.Printf("unsupported type: %T\n", t)
	}
}
