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
	fmt.Println("1. Améliorer la cuirasse (200 crédits)")
	fmt.Println("2. Ameliorer le sac a dos (250 crédits) ")
	fmt.Println("3. Retour")

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
			color.Green("Vous avez amélioré votre Cuirasse.")
			time.Sleep(3 * time.Second)

		} else {
			fmt.Println("\n")
			color.Red("Ta pas assez d'argent va farm.")
			time.Sleep(3 * time.Second)

		}
	case 2:
		if player.Credits >= 250 {
			//player.ImproveBag = "sac amélioré"
			player.Credits -= 250
			fmt.Println("\n")
			color.Yellow("Ta améliorer ton sac !")
			time.Sleep(3 * time.Second)
		} else {
			fmt.Println("\n")
			color.Red("T'es pauvre")
			time.Sleep(3 * time.Second)
		}
		VisitForge()

	}
	fmt.Println("———————————————————————————————————————————————")
}
