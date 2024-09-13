package inventory

import (
	"fmt"
	src "projetred"
	"projetred/player"
	"time"
)

func ShowInventory() {
	src.ClearScreen()
	player := player.GetPlayer()
	fmt.Println("Vous regardez votre Inventaire.")
	fmt.Printf("Potions: %d\n", player.Potions)
	fmt.Printf("Armure: %s\n", player.Armor)
	fmt.Println("———————————————————————————————————————————————")
	time.Sleep(3 * time.Second)
}
