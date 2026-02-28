package unittests

// import "errors"

func Div(a, b int) (int, error) {
	// if b == 0 {
	// 	return 0, errors.New("division by zero")
	// }
	return a / b, nil
}

func Add(a, b int) int {
	return a + b
}
