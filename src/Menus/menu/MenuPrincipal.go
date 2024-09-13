package menu

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"projetred/Menus/Forge"
	"projetred/Menus/Shop"
	"projetred/Menus/inventory"
	"projetred/Menus/stats"
	"projetred/combat"
)

func Menu() {
	var ans int
	for {
		color.White("Choisissez ce que vous voulez faire.")
		color.White("———————————————————————————————————————————————")
		color.Cyan("1. Regarder votre .")
		color.Green("2. Regarder vos Stats.")
		color.Yellow("3. Aller voir le marchand à metro.")
		fmt.Println("4. Aller voir le Forgeron.")
		color.White("5. Aller au Front !")
		color.Red("6. Quitter.")
		color.White("———————————————————————————————————————————————")

		_, err := fmt.Scan(&ans)
		if err != nil {
			return
		}

		switch ans {
		case 1:
			inventory.ShowInventory()
		case 2:
			stats.ShowStats()
		case 3:
			Shop.VisitShop()
		case 4:
			Forge.VisitForge()
		case 5:
			combat.StartCombat()
		case 6:
			fmt.Println("J'espère à bientôt !")
			color.White("———————————————————————————————————————————————")
			os.Exit(1)
		default:
			fmt.Println("T'as pas compris la consigne, fais un effort stp.")
			color.White("———————————————————————————————————————————————")
		}
	}
}
