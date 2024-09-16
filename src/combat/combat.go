package combat

import (
	"fmt"
	"projetred/player"
	"time"
)

func StartCombat() {
	fmt.Println("Vous voulez vous battre, battez-vous !")
	time.Sleep(1 * time.Second)
	p := player.GetPlayer() // `p` pour le joueur

	// Définition d'un ennemi de base (larbin)
	enemy := player.Player{ // On assume que `player.Player` est bien la structure de l'ennemi
		Name:   "Larbin",
		Health: 50,
		Attack: 10,
	}

	// Boucle de combat
	for enemy.Health > 0 && p.Health > 0 {
		fmt.Printf("Ennemi: %s - Santé: %d\n", enemy.Name, enemy.Health)
		fmt.Printf("Votre Santé: %d, Mana: %d\n", p.Health, p.Mana)
		fmt.Println("Choisissez une action :")
		fmt.Println("1. Attaque basique,classique")
		fmt.Println("2. Utiliser un sort (25 mana)")
		fmt.Println("3. Utiliser une potion")

		var action int
		_, err := fmt.Scan(&action)
		if err != nil {
			fmt.Println("Erreur de saisie, réessaye.")
			continue
		}

		switch action {
		case 1:
			fmt.Printf("Vous attaquez %s et lui inflige %d dégâts.\n", enemy.Name, p.Attack)
			enemy.Health -= p.Attack

		case 2:
			if p.Mana >= 25 {
				fmt.Printf("Vous utilisez un sort et infligez %d dégâts à %s.\n", p.Attack+10, enemy.Name)
				enemy.Health -= (p.Attack + 10)
				p.Mana -= 25
			} else {
				fmt.Println("Pas assez de mana haha.")
			}
		//case 3:
		//if p.Potions > 0 {
		//	fmt.Println("Vous utilisez une potion et regagnez 30 points de santé.")
		//p.Health += 30
		//	if p.Health > 100 { // Assurez-vous que la santé ne dépasse pas un maximum (ex. 100)
		//		p.Health = 100
		//	}
		//p.Potions--
		//} else {
		//	fmt.Println("Vous n'avez plus de potions.")
		//}
		default:
			fmt.Println("Action non reconnu, essaye encore.")
		}

		// Ennemi attaque après chaque action
		if enemy.Health > 0 {
			fmt.Printf("%s attaque et vous inflige %d dégâts.\n", enemy.Name, enemy.Attack)
			p.Health -= enemy.Attack
		}

		// Si le joueur ou l'ennemi meurt
		if p.Health <= 0 {
			fmt.Println("Tu es mort !")
			fmt.Println("Aziz vous ressucite avec 50% de point de vie.")
			p.Health += p.MAX_Health / 2
			break
		} else if enemy.Health <= 0 {
			fmt.Println("Vous avez vaincu le sous-fifre!")
			p.Credits += 300
			break
		}
	}

	fmt.Println("———————————————————————————————————————————————")
}
