package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("Привет, обычный текст")
	color.Red("Это текст красного цвета")
	color.Green("Это текст зелёного цвета")
}
