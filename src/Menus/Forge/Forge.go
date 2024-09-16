package Forge

import (
	"fmt"
	"github.com/fatih/color"
	src "projetred"
	"projetred/player"
	"time"
)

func VisitForge() {
	src.ClearScreen()
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
			//player.Armor = "Armure améliorée"
			player.Credits -= 200
			fmt.Println("\n")
			color.Green("Vous avez amélioré votre armure.")
			time.Sleep(3 * time.Second)

		} else {
			fmt.Println("\n")
			color.Red("Vous n'avez pas assez de crédits.")
			time.Sleep(3 * time.Second)

		}
		VisitForge()

	}
	fmt.Println("———————————————————————————————————————————————")
}
