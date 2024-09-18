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
	fmt.Println("------------------------------------------------------------------")
	player := player.GetPlayer()
	fmt.Println("1.Vous êtes chez le Forgeron.")
	fmt.Println("                                        ")
	fmt.Println("4.casque de voyageur (2 Radianites + 50 crédits)")
	fmt.Println("                                        ")
	fmt.Println("5.plasetron de voyageur(4 radianites + 100 crédits)")
	fmt.Println("                                        ")
	fmt.Println("6.jambieres de voyageur(3 radianites + 50 crédits) ")
	fmt.Println("                                        ")
	fmt.Println("7. bottes de voyageur(2 radianites + 50 crédits) ")
	fmt.Println("                                        ")
	fmt.Println("8. Retour")
	fmt.Println("-------------------------------------------------------------------")
	var choice int
	for {
		_, err := fmt.Scan(&choice)
		if err != nil || !(0 <= choice && choice <= 9) {
			fmt.Println("Chef ?")
			time.Sleep(1 * time.Second)
			src.ClearScreen()
			continue
		}
		break
	}

	switch choice {
	case 4:
		if player.Radianite >= 2 && player.Credits >= 50 {
			player.UseItem("Radianite")
			player.Credits -= 50
			fmt.Println("\n")
			color.Green("Tu as obtenu ton casque !")
			player.ItemToInventory("Casque")
			time.Sleep(3 * time.Second)
		} else {
			fmt.Println("\n")
			color.Red("Tu n'as pas assez de radianite ou de crédits.")
			time.Sleep(3 * time.Second)
		}

	case 5:
		if player.Radianite >= 4 && player.Credits >= 100 {
			player.Radianite -= 4
			player.Credits -= 100
			fmt.Println("\n")
			color.Green("Tu as obtenu ton plastron !")
			time.Sleep(3 * time.Second)
		} else {
			fmt.Println("\n")
			color.Red("Tu n'as pas assez de radianite ou de crédits.")
			time.Sleep(3 * time.Second)
		}
	case 6:
		if player.Radianite >= 3 && player.Credits >= 50 {
			player.Radianite -= 3
			player.Credits -= 50
			fmt.Println("\n")
			color.Green("Tu as obtenu tes jambières !")
			time.Sleep(4 * time.Second)
		} else {
			fmt.Println("\n")
			color.Red("Tu n'as pas assez de radianite ou de crédits.")
			time.Sleep(4 * time.Second)
		}

	case 7:
		if player.Radianite >= 2 && player.Credits >= 50 {
			player.Radianite -= 2
			player.Credits -= 50
			fmt.Println("\n")
			color.Green("Tu as obtenu tes bottes !")
			time.Sleep(4 * time.Second)

		} else {
			fmt.Println("\n")
			color.Red("Tu n'as pas assez de radianite ou de crédits.")
			time.Sleep(4 * time.Second)
		}

		VisitForge()
	}

	fmt.Println("———————————————————————————————————————————————")
}
