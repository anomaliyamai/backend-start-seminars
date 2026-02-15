package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("=== МАССИВЫ ===")

	var heroes [5]string
	heroes[0] = "Invoker"
	heroes[1] = "Pudge"
	heroes[2] = "Crystal Maiden"
	fmt.Printf("Пики команды: %v\n", heroes)
	fmt.Printf("Количество героев: %d\n", len(heroes))

	roles := [...]string{"carry", "support", "offlane"}
	fmt.Printf("Роли: %v\n", roles)

	fmt.Println("Перебор массива ролей:")
	for i, role := range roles {
		fmt.Printf("Позиция %d: %s\n", i, role)
	}

	fmt.Println("\nСравнение массивов:")
	team1 := [3]string{"Invoker", "Pudge", "Crystal Maiden"}
	team2 := [3]string{"Invoker", "Pudge", "Crystal Maiden"}
	team3 := [3]string{"Drow Ranger", "Lion", "Axe"}

	fmt.Printf("team1 == team2: %t\n", team1 == team2)
	fmt.Printf("team1 == team3: %t\n", team1 == team3)

	fmt.Println("\n=== СЛАЙСЫ ===")


// 	type slice struct {
//     ptr *T   // указатель на массив
//     len int  // длина
//     cap int  // capacity
// }

	var items []string
	fmt.Printf("Пустой инвентарь: %v, предметов: %d, слотов: %d\n", items, len(items), cap(items))

	items = append(items, "Blink Dagger", "Black King Bar", "Aghanim's Scepter")
	items = append(items, []string{"aaa", "bbc"}...)
	fmt.Printf("После покупки: %v, предметов: %d, слотов: %d\n", items, len(items), cap(items))

	stash := make([]string, 2, 6)
	stash[0] = "Observer Ward"
	stash[1] = "Smoke of Deceit"
	fmt.Printf("Склад: %v, предметов: %d, слотов: %d\n", stash, len(stash), cap(stash))

	teammates := []string{"Dendi", "No[o]ne", "RAMZES"}
	fmt.Printf("Тиммейты: %v\n", teammates)

	fmt.Println("\nОперации со слайсами:")
	networth := []int{1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000}
	fmt.Printf("Нетворс команды: %v\n", networth)
	fmt.Printf("Нетворс 3-5 игроков: %v\n", networth[2:5])
	fmt.Printf("Нетворс первых 3: %v\n", networth[:3])
	fmt.Printf("Нетворс последних 3: %v\n", networth[7:])
	fmt.Printf("Копия всего нетворса: %v\n", networth[:])

	fmt.Println("\nДобавление героев в пул:")
	heroPool := []string{"Invoker", "Pudge"}
	heroPool = append(heroPool, "Storm Spirit")
	fmt.Printf("После добавления: %v\n", heroPool)

	moreHeroes := []string{"Juggernaut", "Lion"}
	heroPool = append(heroPool, moreHeroes...)
	fmt.Printf("После расширения пула: %v\n", heroPool)

	fmt.Println("\nКопирование слайсов:")
	source := []string{"Invoker", "Pudge", "Crystal Maiden"}
	dest := make([]string, len(source))
	copy(dest, source)
	fmt.Printf("Исходный: %v\n", source)
	fmt.Printf("Копия: %v\n", dest)

	dest[0] = "Storm Spirit"
	fmt.Printf("После изменения копии:\n")
	fmt.Printf("Исходный: %v\n", source)
	fmt.Printf("Копия: %v\n", dest)

	fmt.Println("\nУдаление элементов:")
	skills := []string{"Sunstrike", "Tornado", "EMP", "Meteor", "Deafening Blast"}
	fmt.Printf("Все скилы: %v\n", skills)

	indexToRemove := 2
	skills = append(skills[:indexToRemove], skills[indexToRemove+1:]...)
	fmt.Printf("После удаления EMP: %v\n", skills)

	fmt.Println("\nСлайсы из массивов:")
	fullTeam := [10]string{"Invoker", "Pudge", "Crystal Maiden", "Drow Ranger", "Lion", "Axe", "Juggernaut", "Storm Spirit", "Mirana", "Phantom Assassin"}
	radiant := fullTeam[0:5]
	dire := fullTeam[5:10]

	fmt.Printf("Радиант: %v\n", radiant)
	fmt.Printf("Дайр: %v\n", dire)

	fmt.Println("\nИзменение underlying array:")
	radiant[0] = "Techies"
	fmt.Printf("Полная команда после изменения: %v\n", fullTeam)
	fmt.Printf("Радиант: %v\n", radiant)

	fmt.Println(slices.Equal(radiant, dire))

	a := make([]string, 2, 6)
	a = append(a, "aa", "bb")
	b := a
	// changeSlice(&b)
	// fmt.Println(a, len(a))
	// fmt.Println(b, len(b))
	b = append(b, "aaa", "ccc", "ddd")

	changeSlice(b)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
}

func changeSlice(a []string) {
	a[2] = "zz"
}
