package main

import "fmt"

func increaseMMRByValue(mmr int) {
	mmr += 100
}

func increaseMMRByPointer(mmr *int) {
	*mmr += 100
}

func addHero(heroes []string, newHero string) {
	heroes = append(heroes, newHero)
}

func addHeroByPointer(heroes *[]string, newHero string) {
	*heroes = append(*heroes, newHero)
}

func main() {
	fmt.Println("=== УКАЗАТЕЛИ ===")

	mmr := 4200
	mmrPtr := &mmr
	fmt.Println("MMR:", mmr, "по указателю:", *mmrPtr)
	*mmrPtr = 5000
	fmt.Println("Изменено через указатель:", mmr)

	playerMMR := 3000
	increaseMMRByValue(playerMMR)
	fmt.Println("MMR после ByValue:", playerMMR)
	increaseMMRByPointer(&playerMMR)
	fmt.Println("MMR после ByPointer:", playerMMR)

	heroes := []string{"Invoker", "Pudge"}
	addHero(heroes, "Crystal Maiden")
	fmt.Println("Герои после addHero:", heroes)
	addHeroByPointer(&heroes, "Drow Ranger")
	fmt.Println("Герои после addHeroByPointer:", heroes)

	value1, value2 := 100, 100
	ptr1, ptr2 := &value1, &value2
	ptr3 := &value1
	fmt.Printf("Адрес ptr1: %p\n", ptr1)
	fmt.Printf("Адрес ptr3: %p\n", ptr3)

	fmt.Println("ptr1 == ptr2:", ptr1 == ptr2)
	fmt.Println("ptr1 == ptr3:", ptr1 == ptr3)
	fmt.Println("*ptr1 == *ptr2:", *ptr1 == *ptr2)

	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("nil указатель")
	}

	x := 42
	ptr := &x
	ptrPtr := &ptr
	fmt.Println("x через ptrPtr:", **ptrPtr)
}
