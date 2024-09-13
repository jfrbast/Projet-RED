package Forge

import (
	"fmt"
	"projetred/player"
)

func VisitForge() {
	player := player.GetPlayer()
	fmt.Println("Vous êtes chez le Forgeron.")
	fmt.Println("1. Améliorer l'armure (200 crédits)")
	fmt.Println("2. Retour")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		return
	}

	switch choice {
	case 1:
		if player.Credits >= 200 {
			player.Armor = "Armure améliorée"
			player.Credits -= 200
			fmt.Println("Vous avez amélioré votre armure.")

		} else {
			fmt.Println("Vous n'avez pas assez de crédits.")

		}
		VisitForge()

	}
	fmt.Println("———————————————————————————————————————————————")
}
