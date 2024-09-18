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
	fmt.Printf("Initiative : %d \n", player.Initiative)
	fmt.Printf("Limite d'inventaire : %d \n", player.InventoryMax)
	fmt.Printf("Chestplate : %v \n", player.Chestplate)
	fmt.Printf("Helmet : %v \n", player.Helmet)
	fmt.Printf("Boots : %v \n", player.Boots)

	fmt.Println("———————————————————————————————————————————————")
	time.Sleep(5 * time.Second)
}
