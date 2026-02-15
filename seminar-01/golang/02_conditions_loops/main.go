package main

import "fmt"

func main() {
	fmt.Println("=== УСЛОВНЫЕ КОНСТРУКЦИИ ===")
	fmt.Println()

	mmr := 4200

	if mmr <= 3000 {
		fmt.Println("Уровень игры: Stray228 - Титан")
	} else {
		fmt.Println("Ваш уровень игры: рыцарь")
	}

	if mmr >= 3000 && mmr < 7000 {
		fmt.Println("Вы - страж")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "отрицательное")
	} else if num < 10 {
		fmt.Println(num, "с одной цифрой")
	} else {
		fmt.Println(num, "из нескольких цифр")
	}

	gpm := 0
	switch {
	case gpm >= 800:
		fmt.Println("Фарм машина! (GPM: отлично)")
	case gpm >= 600:
		fmt.Println("Фарм хороший (GPM: хорошо)")
	case gpm >= 400:
		fmt.Println("Фарм средний (GPM: удовлетворительно)")
	default:
		fmt.Println("Фарм слабый (GPM: неудовлетворительно)")
	}

	day := "суббота"
	switch day {
	case "понедельник", "вторник", "среда", "четверг", "пятница":
		fmt.Println("Рабочий день — катать только вечером")
	case "суббота", "воскресенье":
		fmt.Println("Выходной — можно устроить марафон по доте")
	default:
		fmt.Println("Неизвестный день — но доту никто не отменял")
	}

	fmt.Println("\n=== ЦИКЛЫ ===")

	fmt.Println("\nЦикл for (как while):")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("\nЦикл for (классический):")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("\nЦикл for по range:")
	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Println("\nБесконечный цикл с break:")
	for {
		fmt.Println("loop")
	}

	fmt.Println("\nЦикл for с continue:")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
