package main

import "fmt"

type Error struct{}

func (e *Error) Error() string {
	return "my error"
}

func canError() *Error {
	return nil
}

func main() {
	var err error

	// (methods, value) -> (not nil, nil)
	if err != nil {
		fmt.Println(err)
	}
}
