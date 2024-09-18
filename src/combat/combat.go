package combat

import (
	"fmt"
	"github.com/fatih/color"
	"projetred/Menus/inventory"
	"projetred/player"
	"time"
)

func StartCombat() {
	fmt.Println("Vous voulez vous battre, battez-vous !")
	time.Sleep(1 * time.Second)
	p := player.GetPlayer()
	var enemy player.Player
	if p.Level < 3 {
		enemy = player.Player{
			Name:   "Kèr et dîne",
			Health: 50,
			Attack: 5,
		}
	}
	if p.Level > 3 && p.Level < 6 {
		enemy = player.Player{
			Name:   "Antonain",
			Health: 30,
			Attack: 30,
		}
	}
	if p.Level > 6 && p.Level < 10 {
		enemy = player.Player{
			Name:   "Phabieau",
			Health: 70,
			Attack: 13,
		}
	}
	if p.Level > 10 {
		enemy = player.Player{
			Name:   "Lillhian",
			Health: 100,
			Attack: 19,
		}
	}

	for enemy.Health > 0 && p.Health > 0 {
		fmt.Printf("Ennemi: %s - Santé: %d\n", enemy.Name, enemy.Health)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Votre Santé: %d, Mana: %d\n", p.Health, p.Mana)
		time.Sleep(700 * time.Millisecond)
		fmt.Println("Choisissez une action :")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("1. Attaque basique,classique")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("2. Utiliser un sort (25 mana)")
		time.Sleep(500 * time.Millisecond)
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
			time.Sleep(500 * time.Millisecond)
			enemy.Health -= p.Attack

		case 2:
			if p.Mana >= 25 {
				fmt.Printf("Vous utilisez un sort et infligez %d dégâts à %s.\n", p.Attack+10, enemy.Name)
				time.Sleep(500 * time.Millisecond)
				enemy.Health -= (p.Attack + 10)
				p.Mana -= 25
			} else {
				fmt.Println("Pas assez de mana haha.")
				time.Sleep(500 * time.Millisecond)
			}
		case 3:
			inventory.UseItemFromInventory("Potion de Soin")
			time.Sleep(500 * time.Millisecond)
			player.UseItem("Potion de Soin ")
			time.Sleep(500 * time.Millisecond)
		default:
			fmt.Println("Action non reconnu, essaye encore.")
			time.Sleep(500 * time.Millisecond)
		}

		if enemy.Health > 0 {

			if enemy.Name == "Kèr et dîne" {
				fmt.Printf("Kèr attaque et vous inflige %d dégâts.\n", enemy.Attack)
				time.Sleep(300 * time.Millisecond)
				fmt.Printf("Dîne attaque et vous inflige %d dégâts.\n", enemy.Attack)
				time.Sleep(500 * time.Millisecond)
				p.Health -= enemy.Attack
				p.Health -= enemy.Attack
			} else {
				fmt.Printf("%s attaque et vous inflige %d dégâts.\n", enemy.Name, enemy.Attack)
				time.Sleep(500 * time.Millisecond)
				p.Health -= enemy.Attack
			}

		}

		if p.Health <= 0 {
			fmt.Println("Tu es mort !")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Aziz vous ressucite avec 50% de point de vie.")
			time.Sleep(500 * time.Millisecond)
			p.Health += p.MAX_Health / 2
			p.Credits += 100
			break
		} else if enemy.Health <= 0 {
			fmt.Println("Vous avez vaincu le sous-fifre!")
			time.Sleep(500 * time.Millisecond)
			color.Yellow("vous avez gangné 300 crédits et 10 d'XP !")
			time.Sleep(5 * time.Second)
			p.Credits += 300
			p.XP += 10
			break
		}
	}

	fmt.Println("———————————————————————————————————————————————")
}
