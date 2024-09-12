package projetred

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Menu() {
	var ans int
	for {
		color.White("Choisissez ce que vous voulez faire.")
		color.Blue("1.Regarder votre Inventaire.")
		color.Green("2.Regarder vos Stats.")
		color.Yellow("3.Aller voir le Marchant.")
		fmt.Println("4.Aller voir le Forgeron.")
		color.White("5.Aller au Combat !")
		color.Red("6.Quitter.")
		_, err := fmt.Scan(&ans)
		if err != nil {
			return
		}
		switch ans {
		case 1:
			fmt.Println("Vous regarder votre Inventaire.")
		case 2:
			fmt.Println("Vous regarder Vos Stats.")
		case 3:
			fmt.Println("Vous êtes chez le Marchant.")
		case 4:
			fmt.Println("Vous êtes chez le Forgeron.")
		case 6:
			fmt.Println("J'espère à bientôt !")
			os.Exit(1)
		default:
			fmt.Println("T'as pas compris la consigne, fais un effort stp.")
		}
	}
}
