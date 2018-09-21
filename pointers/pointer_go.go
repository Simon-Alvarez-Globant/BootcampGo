package main

import "fmt"

func main() {
	a := "idk"
	ap := &a

	printref(a)
	printval(ap)
}

func printref(x string) {
	fmt.Println(x)
}
func printval(y *string) {
	fmt.Println(*y)

}

// idk  value
// idk  reference
