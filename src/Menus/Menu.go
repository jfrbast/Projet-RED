package projetred

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Menu() {
	var ans int
	for {
		color.Blue("Choisissez ce que vous voulez faire.")
		fmt.Println("1.Regarder votre Inventaire.")
		fmt.Println("2.Regarder vos Stats.")
		color.Red("3.Quitter")
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
			fmt.Println("J'espère à bientôt !")
			os.Exit(1)
		default:
			fmt.Println("T'as pas compris la consigne, fais un effort stp.")
		}
	}
}
