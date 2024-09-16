package inventory

import (
	"fmt"
	"github.com/fatih/color"
	"projetred/player"
)

func ShowInventory() {
	p := player.GetPlayer()
	var choice int

	for {
		color.Cyan("Votre inventaire :")
		color.White("———————————————————————————————————————————————")

		for i, item := range p.Inventory {
			fmt.Printf("%d. %s (Quantité: %d)\n", i+1, item.Name, item.Quantity)
		}

		color.White("———————————————————————————————————————————————")
		fmt.Println("0. Retourner au menu principal")
		fmt.Print("Choisissez un objet pour l'utiliser, ou 0 pour quitter : ")

		_, err := fmt.Scan(&choice)
		if err != nil || choice == 0 {
			return
		}

		if choice > 0 && choice <= len(p.Inventory) {
			item := p.Inventory[choice-1]
			UseItemFromInventory(item.Name)
		} else {
			color.Red("Choix non valide.")
		}
	}
}

func UseItemFromInventory(itemName string) {
	p := player.GetPlayer()

	switch itemName {
	case "Potion de Soin":
		if player.UseItem("Potion de Soin") {
			p.Health += 50
			if p.Health > 100 {
				p.Health = 100
			}
			color.Green("Vous avez utilisé une Potion de Soin. Santé actuelle : %d/%d", p.Health, p.MAX_Health)
		} else {
			color.Red("Vous n'avez plus de potions de soin.")
		}
	default:
		color.Red("Cet objet ne peut pas être utilisé.")
	}
}
