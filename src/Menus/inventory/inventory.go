package inventory

import (
	"fmt"
	"projetred/player"
)

func ShowInventory() {
	player := player.GetPlayer()
	fmt.Println("Vous regardez votre Inventaire.")
	fmt.Printf("Potions: %d\n", player.Potions)
	fmt.Printf("Armure: %s\n", player.Armor)
	fmt.Println("———————————————————————————————————————————————")
}
