package inventory

import (
	"fmt"
	"github.com/fatih/color"
	src "projetred"
	"projetred/player"
	"time"
)

func ShowInventory() {
	p := player.GetPlayer()
	var choice int

	for {
		src.ClearScreen()
		color.Cyan("    _/_/_/                                              _/                                   \n     _/    _/_/_/    _/      _/    _/_/    _/_/_/    _/_/_/_/    _/_/    _/  _/_/  _/    _/  \n    _/    _/    _/  _/      _/  _/_/_/_/  _/    _/    _/      _/    _/  _/_/      _/    _/   \n   _/    _/    _/    _/  _/    _/        _/    _/    _/      _/    _/  _/        _/    _/    \n_/_/_/  _/    _/      _/        _/_/_/  _/    _/      _/_/    _/_/    _/          _/_/_/     \n                                                                                     _/      \n                                                                                _/_/         ")
		color.Cyan("Votre sac a dos :")
		color.White("———————————————————————————————————————————————")

		for i, item := range p.Inventory {
			fmt.Printf("%d. %s (Quantité: %d)\n", i+1, item.Name, item.Quantity)
		}

		color.White("———————————————————————————————————————————————")
		fmt.Println("0. Retourner au lobby")
		fmt.Print("Choisissez un objet pour l'utiliser, ou 0 pour quitter : ")

		for {
			choice2, err := fmt.Scan(&choice)
			if err != nil || !(1 <= choice2 && choice2 <= 4) {
				fmt.Println("Chef ?")
				choice2 = 0
				continue
			}
			break
		}
		if choice == 0 {
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
			if p.Health > p.MAX_Health {
				p.Health = p.MAX_Health
			}
			color.Green("Vous avez utilisé une Potion de Soin. Santé actuelle : %d/%d", p.Health, p.MAX_Health)
		} else {
			color.Red("Tu peux plus te soigner frerot.")
		}
	case "Potion de Mana":
		if player.UseItem("Potion de Mana") {
			p.Mana += 50
			color.Green("Vous avez utilisé une Potion de Mana. Mana actuelle : %d", p.Mana)
			time.Sleep(3 * time.Second)
		}
	default:
		color.Red("tu peux pas l'utiliser !.")

	}

}
