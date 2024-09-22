package combat

import (
	"fmt"
	"github.com/fatih/color"
	src "projetred"
	"projetred/Menus/inventory"
	"projetred/player"
	"time"
)

func StartCombat() {
	src.ClearScreen()
	color.Cyan("     _/_/_/                            _/                    _/         _/  \n  _/          _/_/    _/_/_/  _/_/    _/_/_/      _/_/_/  _/_/_/_/     _/   \n _/        _/    _/  _/    _/    _/  _/    _/  _/    _/    _/         _/    \n_/        _/    _/  _/    _/    _/  _/    _/  _/    _/    _/                \n _/_/_/    _/_/    _/    _/    _/  _/_/_/      _/_/_/      _/_/     _/      ")
	fmt.Println("Vous voulez vous battre, battez-vous !")
	time.Sleep(3 * time.Second)
	p := player.GetPlayer()
	var enemy player.Player
	if p.Level < 3 {
		enemy = player.Player{
			Name:       "Kèr et dîne le duo légendaire",
			Health:     50,
			Attack:     5,
			Initiative: 8,
		}
	}
	if p.Level >= 3 && p.Level < 6 {
		enemy = player.Player{
			Name:       "Antonain",
			Health:     30,
			Attack:     30,
			Initiative: 16,
		}
	}
	if p.Level >= 6 && p.Level < 10 {
		enemy = player.Player{
			Name:       "Phabieau",
			Health:     70,
			Attack:     13,
			Initiative: 20,
		}
	}
	if p.Level >= 10 {
		enemy = player.Player{
			Name:       "Lillhian",
			Health:     100,
			Attack:     19,
			Initiative: 30,
		}
	}
	fmt.Printf("Votre Initiative: %d | Initiative de l'ennemi (%s): %d\n", p.Initiative, enemy.Name, enemy.Initiative)
	if enemy.Initiative > p.Initiative {
		if enemy.Name == "Kèr et dîne" {
			color.Red("Kèr attaque et vous inflige %d dégâts.\n", enemy.Attack)
			time.Sleep(300 * time.Millisecond)
			color.Red("Dîne attaque et vous inflige %d dégâts.\n", enemy.Attack)
			time.Sleep(500 * time.Millisecond)
			p.Health -= enemy.Attack
			p.Health -= enemy.Attack
		} else {
			color.Red("%s attaque et vous inflige %d dégâts.\n", enemy.Name, enemy.Attack)
			time.Sleep(500 * time.Millisecond)
			p.Health -= enemy.Attack
		}
	}
	for enemy.Health > 0 && p.Health > 0 {
		fmt.Printf("Ennemi: %s - Santé: %d\n", enemy.Name, enemy.Health)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Votre Santé: %d, Mana: %d\n", p.Health, p.Mana)
		time.Sleep(700 * time.Millisecond)
		color.Blue("Choisissez une action :\n")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("1. Attaque basique, coup de couteau")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("2. Utiliser Pogo(25 mana)")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("3.Utiliser Grenade (30 mana)")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("4. Utiliser une potion de heal")
		time.Sleep(500 * time.Millisecond)
		color.Yellow("5.Appeler G-Rémy au combat !\n ")

		var action int
		_, err := fmt.Scan(&action)
		if err != nil {
			fmt.Println("Erreur de saisie, réessaye.")
			continue
		}

		switch action {
		case 1:
			color.Green("Vous attaquez %s et lui inflige %d dégâts.\n", enemy.Name, p.Attack)
			enemy.Health -= p.Attack

		case 2:
			if p.Mana >= 25 {
				isFind := false
				for index, spell := range p.Spells {
					if spell.SpellName == "Pogo" && spell.Quantity > 0 {
						isFind = true
						p.Spells[index].Quantity -= 1
						player.UseItem("Pogo")
						fmt.Printf("Vous lancez %s et brûlez %s pour %v dégâts .\n", spell.SpellName, enemy.Name, (spell.Damage))
						enemy.Health -= (spell.Damage / 2)
						time.Sleep(time.Second * 1)
						fmt.Printf("vous brûlez %s pour %v dégâts .\n", enemy.Name, (spell.Damage))
						enemy.Health -= (spell.Damage / 2)
						time.Sleep(time.Second * 1)
						fmt.Printf("Vous brûlez %s pour %v dégâts .\n", enemy.Name, (spell.Damage))
						enemy.Health -= (spell.Damage / 2)
						time.Sleep(time.Second * 1)
						p.Mana -= 25

					}
					break
				}
				if !isFind {
					color.Red("Vous ne possédez pas cette compétence!\n")
					continue
				}
			} else {
				fmt.Println("Pas assez de mana!")
				continue
			}
		case 3:
			if p.Mana >= 30 {
				isFind := false
				for index, spell := range p.Spells {
					if spell.SpellName == "Grenade" && spell.Quantity > 0 {
						isFind = true
						p.Spells[index].Quantity -= 1
						player.UseItem("Grenade")
						fmt.Printf("Vous lancez une %s et infligez %d de dégats.\n", spell.SpellName, spell.Damage)
						enemy.Health -= spell.Damage
						p.Mana -= 30
						break
					}
				}
				if !isFind {
					color.Red("Vous ne possédez pas cette compétence!")
					time.Sleep(500 * time.Millisecond)
					continue
				}
			} else {
				color.Red("Pas assez de mana!")
				time.Sleep(500 * time.Millisecond)
				continue
			}

		case 4:
			if player.UseItem("Potion de Soin") {
				player.UseItem("Potion de Soin ")
				time.Sleep(500 * time.Millisecond)
				inventory.UseItemFromInventory("Potion de Soin")
				time.Sleep(500 * time.Millisecond)
			} else {
				color.Red("Vous n'avez pas cette potion")
				continue
			}

		case 5:
			if player.UseItem("Jérémie l'intrépide") {
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Gé-ré-mi vient à votre secours !")
				time.Sleep(1500 * time.Millisecond)
				fmt.Println("Jérem semble avoir peur ?!?")
				time.Sleep(1500 * time.Millisecond)
				fmt.Println("G-rem-i prend la fuite ?")
				time.Sleep(1500 * time.Millisecond)
				color.Red("JérémY est parti tout comme vos 5000 crédits...\n")
				time.Sleep(500 * time.Millisecond)
			} else {
				color.Red("Vous ne possedez pas cette compétence.\n")
				continue
			}

		default:
			fmt.Println("Action non reconnu, essaye encore.")
			time.Sleep(500 * time.Millisecond)
			continue
		}

		if enemy.Health > 0 {

			if enemy.Name == "Kèr et dîne le duo légendaire" {
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
			color.Green("Vous avez vaincu le sous-fifre!\n")
			time.Sleep(500 * time.Millisecond)
			color.Yellow("vous avez gangné 300 crédits, 10 d'XP et 2 de Radianite !")
			time.Sleep(5 * time.Second)
			p.Credits += 300
			p.XP += 10
			p.Initiative += 1
			if p.XP >= p.Xplvl {
				p.Level++
				p.XP = 0
				p.Xplvl += 7
				color.Yellow("Vous passez au niveau supérieur! \n")
			}
			player.ItemToInventory("Radianite", 2)
			break
		}
	}

	fmt.Println("———————————————————————————————————————————————")
}
