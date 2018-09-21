package main

import (
	"fmt"
)

//
type ErrorSrruct struct {
	Msg  string
	Type string
}

func (e ErrorSrruct) Error() string {
	return fmt.Sprintf("%v", e.Msg)
}

//
func ErrorHandler(t string) error {
	switch t {
	case "internal":
		return ErrorSrruct{
			"bleeding internally",
			t,
		}
	case "thirdparty":
		return ErrorSrruct{
			"another player has joined",
			t,
		}
	case "other":
		return ErrorSrruct{
			"who cares",
			t,
		}
	}
	return ErrorSrruct{
		"who cares",
		t,
	}
}

func main() {

	err := ErrorHandler("internal")
	fmt.Println(err)

}
