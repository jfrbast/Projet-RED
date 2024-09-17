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
	fmt.Println("1.Vous êtes chez le Forgeron.")
	fmt.Println("2. Améliorer la cuirasse (200 crédits)")
	fmt.Println("3. Ameliorer le sac a dos (250 crédits) ")
	fmt.Println("4.casque de voyageur (2 Radianites + 50 crédits)")
	fmt.Println("5.plasetron de voyageur(4 radianites + 100 crédits)")
	fmt.Println("6.jambieres de voyageur(3 radianites + 50 crédits)")
	fmt.Println("7. bottes de voyageur(2 radianites + 50 crédits)")
	fmt.Println("8. Retour")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		return
	}

	switch choice {
	case 2:
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
	case 3:
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
	case 4:
		if player.Radianite+player.Credits >= 2+50 {
			player.Radianite += player.Credits - 50
			fmt.Println("\n")
			color.Green("Tu a obtenu ton casque !")
			time.Sleep(3 * time.Second)
		} else {
			fmt.Println("\n")
			color.Red("Ta pas assez d'argent ")
			time.Sleep(3 * time.Second)
		}

	case 5:
		if player.Radianite+player.Credits >= 4+100 {
			player.Radianite += player.Credits - 100
			fmt.Println("\n")
			color.Green("Tu a obtenu ton plastron !")
			time.Sleep(3 * time.Second)
		} else {
			fmt.Println("\n")
			color.Red("Ta pas assez d'argent ")
			time.Sleep(3 * time.Second)
		}
	case 6:
		if player.Radianite+player.Credits >= 3+50 {
			player.Radianite += player.Credits - 50
			fmt.Println("\n")
			color.Green("Tu a obtenu tes jambières !")
			time.Sleep(4 * time.Second)
		} else {
			fmt.Println("\n")
			color.Red("Ta pas assez d'argent ")
			time.Sleep(4 * time.Second)
		}
	case 7:
		if player.Radianite+player.Credits >= 2+50 {
			player.Radianite += player.Credits - 50
			fmt.Println("\n")
			color.Green("Tu a obtenu tes bottes !")
			time.Sleep(4 * time.Second)
		} else {
			fmt.Println("\n")
			color.Red("Ta pas assez d'argent ")
			time.Sleep(4 * time.Second)
		}
		VisitForge()
	}

	fmt.Println("———————————————————————————————————————————————")
}
