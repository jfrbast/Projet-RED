package menu

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	src "projetred"
	"projetred/Menus/Forge"
	"projetred/Menus/Shop"
	"projetred/Menus/inventory"
	"projetred/Menus/stats"
	"projetred/combat"
	"time"
)

func Menu() {
	var ans int
	for {
		src.ClearScreen()
		color.Cyan("    _/      _/                               \n   _/_/  _/_/    _/_/    _/_/_/    _/    _/  \n  _/  _/  _/  _/_/_/_/  _/    _/  _/    _/   \n _/      _/  _/        _/    _/  _/    _/    \n_/      _/    _/_/_/  _/    _/    _/_/_/     ")
		color.White("Choisissez ce que vous voulez faire.")
		color.White("———————————————————————————————————————————————")
		color.Cyan("1. Regarder votre sac.")
		color.Green("2. Regarder vos Stats.")
		color.Yellow("3. Aller voir le marché noir.")
		fmt.Println("4. Aller voir Cypher.")
		color.White("5. Aller au Front !")
		color.White("6.Qui sont-ils ?")
		color.Red("7. Quitter.")
		color.White("———————————————————————————————————————————————")
		for {
			_, err := fmt.Scan(&ans)
			if err != nil || !(1 <= ans && ans <= 7) {
				fmt.Println("Chef ?")
				continue
			}
			break
		}
		switch ans {
		case 1:
			inventory.ShowInventory()
		case 2:
			stats.ShowStats()
		case 3:
			Shop.VisitShop()
		case 4:
			forge.VisitForge()
		case 5:
			combat.StartCombat()
		case 7:

			fmt.Println("à la revoyure !")
			color.White("———————————————————————————————————————————————")
			os.Exit(1)
		case 6:
			Menusecrect()
		default:
			fmt.Println("T'as pas compris la consigne, fais un effort stp.")
			color.White("———————————————————————————————————————————————")
		}
	}
}

func Menusecrect() {
	fmt.Println("les artistes cachés dans le pdf sont : partie 2- ABBA |||| partie 3-Steven Spielberg !")
	time.Sleep(5 * time.Second)
	return

}
