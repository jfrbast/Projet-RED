package forge

import (
	"fmt"
	"projetred/player"
)

func VisitForge() {
	p := player.GetPlayer()

	fmt.Println("Bienvenue à la Forge !")
	fmt.Println("Que voulez-vous forger ?")
	fmt.Println("1. Casque (100 crédits + 2 unités de radianite)")
	fmt.Println("2. Plastron (200 crédits + 4 unités de radianite)")
	fmt.Println("3. Bottes (150 crédits + 3 unités de radianite)")
	fmt.Println("4. Retour")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Erreur lors de la saisie. Veuillez réessayer.")
		return
	}

	switch choice {
	case 1:
		if p.Credits >= 100 && player.HasEnoughRadianite(2) && !p.Helmet {
			p.Credits -= 100
			player.UseRadianite(2)
			p.Helmet = true
			p.MAX_Health += 15
			fmt.Println("Vous avez forgé un casque.")

		} else {
			if p.Helmet {
				fmt.Println("Vous possédez déjà un casque.")
			} else if !player.HasEnoughRadianite(2) {
				fmt.Println("Vous n'avez pas assez de radianite.")
			} else {
				fmt.Println("Vous n'avez pas assez de crédits.")
			}
		}
	case 2:
		if p.Credits >= 200 && player.HasEnoughRadianite(4) && !p.Chestplate {
			p.Credits -= 200
			player.UseRadianite(4)
			p.Chestplate = true
			p.MAX_Health += 30
			fmt.Println("Vous avez forgé un plastron.")

		} else {
			if p.Chestplate {
				fmt.Println("Vous possédez déjà un plastron.")
			} else if !player.HasEnoughRadianite(4) {
				fmt.Println("Vous n'avez pas assez de radianite.")
			} else {
				fmt.Println("Vous n'avez pas assez de crédits.")
			}
		}
	case 3:
		if p.Credits >= 150 && player.HasEnoughRadianite(3) && !p.Boots {
			p.Credits -= 150
			player.UseRadianite(3)
			p.Boots = true
			fmt.Println("Vous avez forgé des bottes.")
			p.MAX_Health += 20

		} else {
			if p.Boots {
				fmt.Println("Vous possédez déjà des bottes.")
			} else if !player.HasEnoughRadianite(3) {
				fmt.Println("Vous n'avez pas assez de radianite.")
			} else {
				fmt.Println("Vous n'avez pas assez de crédits.")
			}
		}
	case 4:
		fmt.Println("Vous quittez la Forge.")
		return
	default:
		fmt.Println("Choix non valide.")
	}
}
