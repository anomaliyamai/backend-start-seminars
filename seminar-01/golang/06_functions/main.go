package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return a / b, nil
}

func calculate(x, y int) (sum, product int) {
	sum = x + y
	product = x * y
	return
}

func variadic(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	result := add(5, 3)
	fmt.Printf("Сумма: %d\n", result)

	quotient, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Частное: %.2f\n", quotient)
	}

	s, p := calculate(4, 6)
	fmt.Printf("Сумма: %d, Произведение: %d\n", s, p)

	total := variadic(1, 2, 3, 4, 5)
	fmt.Printf("Общая сумма: %d\n", total)
}
