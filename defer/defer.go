package main

import "fmt"

func main() {
	//without defer should print world hello
	fmt.Print(" world ")
	fmt.Print(" hello ")

	fmt.Println()

	//with defer should print hello world
	defer fmt.Print(" world ")
	fmt.Print(" hello ")
}
