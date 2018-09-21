package main

import (
	"fmt"
)

func main() {
	maps := map[string]string{
		"rsc": "3711",
		"r":   "2138",
		"gri": "190",
		"adg": "912",
	}

	if maps["dontexist"] == "" {
		fmt.Println("no exist")
	}
}
