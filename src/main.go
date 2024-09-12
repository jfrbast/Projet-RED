package main

import (
	"fmt"
	"os"
)

func main() {
	var ans int
	for {
		fmt.Println("Choisissez ce que vous voulez faire.")
		fmt.Println("1.Regarder votre Inventaire.")
		fmt.Println("2.Regarder vos Stats.")
		fmt.Println("3.Quitter.")
		fmt.Scanf("%c", &ans)
		switch ans {
		case '1':
			fmt.Println("Vous regarder votre Inventaire.")

		case '2':

			fmt.Println("Vous regarder Vos Stats.")
		case '3':
			fmt.Println("J'espère à bientôt !")
			os.Exit(1)
		}

	}
}
