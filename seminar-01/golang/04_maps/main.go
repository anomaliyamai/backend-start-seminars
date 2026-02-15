package main

import "fmt"

func addPlayer(m map[string]int) {
	m["Dendi"] = 9000
	fmt.Println("Внутри функции:", m)
}

func main() {
	fmt.Println("=== МАПЫ ===")

	var playerMMR map[string]int
	playerMMR = make(map[string]int)
	playerMMR["RAMZES"] = 9000
	playerMMR["Solo"] = 7000
	playerMMR["Stray228"] = 15000
	addPlayer(playerMMR)

	fmt.Printf("MMR игроков: %v\n", playerMMR)
	fmt.Printf("MMR Dendi: %d\n", playerMMR["Dendi"])

	// ключ = числовой ID героя
	heroWinrate := map[int]int{
		1: 55,
		2: 48,
		3: 51,
	}

	fmt.Println("\nВинрейт героев по ID:")
	for heroID, winrate := range heroWinrate {
		fmt.Printf("HeroID %d: %d%%\n", heroID, winrate)
	}

	id := 4
	winrate, exists := heroWinrate[id]
	if exists {
		fmt.Printf("Винрейт героя %d: %d%%\n", id, winrate)
	} else {
		fmt.Printf("Нет данных по герою с ID %d\n", id)
		fmt.Println("Как выглядит значение по несуществующему ключу:", winrate)
	}

	delete(heroWinrate, 2)
	fmt.Printf("После удаления героя 2: %v\n", heroWinrate)

	clear(heroWinrate)
	fmt.Printf("После очистки мапы: %v\n", heroWinrate)

	fmt.Println("\nВложенные мапы:")
	teamStats := map[string]map[string]int{
		"Radiant": {
			"kills":   25,
			"deaths":  15,
			"assists": 45,
		},
		"Dire": {
			"kills":   15,
			"deaths":  25,
			"assists": 30,
		},
	}

	for team, stats := range teamStats {
		fmt.Printf("%s: ", team)
		for stat, value := range stats {
			fmt.Printf("%s=%d ", stat, value)
		}
		fmt.Println()
	}

	fmt.Println("\nПроверка существования ключа:")
	radiantKills, ok := teamStats["Radiant"]["kills"]
	if ok {
		fmt.Printf("Радиант убил: %d\n", radiantKills)
	}

	fmt.Println("\n=== Преаллокация мапы через make со слайсом ===")

	players := []string{"Dendi", "RAMZES", "Solo", "Stray228", "No[o]ne"}

	preAllocMMR := make(map[string]int, len(players))

	for i, player := range players {
		preAllocMMR[player] = 5000 + i*1000
	}

	fmt.Println("MMR игроков с преаллокацией:", preAllocMMR)

	fmt.Printf("MMR Dendi: %d\n", preAllocMMR["Dendi"])

	player := "No[o]ne"
	if mmr, exists := preAllocMMR[player]; exists {
		fmt.Printf("%s найден с MMR %d\n", player, mmr)
	} else {
		fmt.Printf("%s не найден\n", player)
	}

	newPlayers := []string{"Ceb", "Miracle"}
	for _, p := range newPlayers {
		preAllocMMR[p] = 6000
	}

	fmt.Println("После добавления новых игроков:", preAllocMMR)
}
