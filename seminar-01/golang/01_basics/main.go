package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

var (
	UINT64_MAX uint64 = math.MaxUint64
)

func main() {
	var name string = "Golang"
	age := 18 - 2               // так можно писать только внутри функций
	version := 1.25             // var version float64 = 1.25
	isDota2GoodGame := true

	// %d         decimal integer
	// %x, %o, %b integer in hexade cimal, octal, binary
	// %f, %g, %e floating-point number: 3.141593 3.141592653589793 3.141593e+00
	// %t         boolean: true or false
	// %c         rune (Unicode code point)
	// %s         string
	// %q         quoted string "abc" or rune 'c'
	// %v         any value in a natural format
	// %T         type of any value
	// %%         literal percent sign (no operand)

	fmt.Printf("Максимум uint64: %x\n", UINT64_MAX)
	fmt.Printf("Язык: %s\n", name)
	fmt.Printf("Возраст: %d лет\n", age)
	fmt.Printf("Версия: %.2f\n", version)
	fmt.Printf("Дота2 хорошая игра: %t\n", isDota2GoodGame)

	var as int
	fmt.Printf("%d\n", as) // zero value

	const pi = 3.14159
	fmt.Printf("Число Пи: %.2f\n", pi)

	var x, y int = 10, 20
	fmt.Printf("x = %d, y = %d, сумма = %d\n", x, y, x+y)

	str := "The best hero" + "🦸‍♂️"

	fmt.Printf("Строчка: %s\n", str)
	fmt.Printf("Байтов в строке: %d\n", len(str))
	fmt.Printf("Rune в строке: %d\n", utf8.RuneCountInString(str))
	fmt.Println("Символы по рунам:")
	for i, r := range str {
		fmt.Printf("Позиция %d: %c (код: %U)\n", i, r, r)
	}
}
