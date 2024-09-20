package forge

import (
	"fmt"
	"github.com/fatih/color"
	src "projetred"
	"projetred/player"
	"time"
)

func VisitForge() {
	p := player.GetPlayer()
	src.ClearScreen()
	color.Cyan("    _/_/_/_/                                      \n   _/        _/_/    _/  _/_/    _/_/_/    _/_/   \n  _/_/_/  _/    _/  _/_/      _/    _/  _/_/_/_/  \n _/      _/    _/  _/        _/    _/  _/         \n_/        _/_/    _/          _/_/_/    _/_/_/    \n                                 _/               \n                            _/_/                  ")
	fmt.Println("Bienvenue chez Cypher!")
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
			time.Sleep(2 * time.Second)
			VisitForge()

		} else {
			if p.Helmet {
				fmt.Println("Vous possédez déjà un casque.")
				time.Sleep(1 * time.Second)
				VisitForge()
			} else if !player.HasEnoughRadianite(2) {
				fmt.Println("Vous n'avez pas assez de radianite.")
				time.Sleep(1 * time.Second)
				VisitForge()
			} else {
				fmt.Println("Vous n'avez pas assez de crédits.")
				time.Sleep(1 * time.Second)
				VisitForge()
			}
		}
	case 2:
		if p.Credits >= 200 && player.HasEnoughRadianite(4) && !p.Chestplate {
			p.Credits -= 200
			player.UseRadianite(4)
			p.Chestplate = true
			p.MAX_Health += 30
			fmt.Println("Vous avez forgé un plastron.")
			time.Sleep(2 * time.Second)
			VisitForge()

		} else {
			if p.Chestplate {
				fmt.Println("Vous possédez déjà un plastron.")
				time.Sleep(1 * time.Second)
				VisitForge()
			} else if !player.HasEnoughRadianite(4) {
				fmt.Println("Vous n'avez pas assez de radianite.")
				time.Sleep(1 * time.Second)
				VisitForge()
			} else {
				fmt.Println("Vous n'avez pas assez de crédits.")
				time.Sleep(1 * time.Second)
				VisitForge()
			}
		}
	case 3:
		if p.Credits >= 150 && player.HasEnoughRadianite(3) && !p.Boots {
			p.Credits -= 150
			player.UseRadianite(3)
			p.Boots = true
			fmt.Println("Vous avez forgé des bottes.")
			p.MAX_Health += 20
			time.Sleep(2 * time.Second)
			VisitForge()

		} else {
			if p.Boots {
				fmt.Println("Vous possédez déjà des bottes.")
				time.Sleep(1 * time.Second)
				VisitForge()
			} else if !player.HasEnoughRadianite(3) {
				fmt.Println("Vous n'avez pas assez de radianite.")
				time.Sleep(1 * time.Second)
				VisitForge()
			} else {
				fmt.Println("Vous n'avez pas assez de crédits.")
				time.Sleep(1 * time.Second)
				VisitForge()
			}
		}
	case 4:
		fmt.Println("Vous quittez la Forge.")
		time.Sleep(2 * time.Second)
		return
	default:
		color.Red("Choix non valide.")
		time.Sleep(2 * time.Second)
		VisitForge()
	}
}
