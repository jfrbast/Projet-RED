package stats

import (
	"fmt"
	"projetred/player"
	"time"
)

func ShowStats() {
	player := player.GetPlayer()
	fmt.Println("Vous regardez vos Stats.")
	fmt.Printf("Nom: %s\n", player.Name)
	fmt.Printf("Agent: %s\n", player.Agent)
	fmt.Printf("Santé: %d\n", player.Health)
	fmt.Printf("Mana: %d\n", player.Mana)
	fmt.Printf("Attaque: %d\n", player.Attack)
	fmt.Printf("Xp: %d\n", player.XP)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Println("———————————————————————————————————————————————")
	time.Sleep(5 * time.Second)
}
