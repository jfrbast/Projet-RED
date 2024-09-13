package Shop

import (
	"fmt"
	"github.com/fatih/color"
	src "projetred"
	"projetred/player"
)

func VisitShop() {
	src.ClearScreen()
	player := player.GetPlayer()
	fmt.Printf("\n")
	color.Yellow("Vous possédez %d Crédits  ", player.Credits)
	fmt.Printf("\n")
	fmt.Println("Vous êtes chez le Marchant.")
	fmt.Println("1. Acheter une potion (50 crédits)")
	fmt.Println("2. Acheter un sort (100 crédits)")
	fmt.Println("3. Retour")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		return
	}

	switch choice {
	case 1:
		if player.Credits >= 50 {
			player.Potions++
			player.Credits -= 50
			fmt.Println("Vous avez acheté une potion.")
		} else {
			fmt.Println("Vous n'avez pas assez de crédits.")
		}
		VisitShop()
	case 2:
		if player.Credits >= 100 {
			player.Mana += 25
			player.Credits -= 100
			fmt.Println("Vous avez acheté un sort, votre mana a augmenté.")
		} else {
			fmt.Println("Vous n'avez pas assez de crédits.")
		}
		VisitShop()
	}
	fmt.Println("———————————————————————————————————————————————")
}
