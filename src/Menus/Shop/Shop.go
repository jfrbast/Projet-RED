package Shop

import (
	"fmt"
	"github.com/fatih/color"
	src "projetred"
	"projetred/player"
	"time"
)

func VisitShop() {
	src.ClearScreen()
	players := player.GetPlayer()
	fmt.Printf("\n")
	color.Yellow("Vous possédez %d Crédits  ", players.Credits)
	fmt.Printf("\n")
	fmt.Println("Vous êtes chez le Marchant.")

	fmt.Print("2. Acheter une potion (50 crédits) ")
	if players.PotionGratuite == true {
		color.Green(" || La première est gratuite !")
	}
	fmt.Println("3. Acheter un sort (100 crédits)")
	fmt.Println("4.Potion de mana (100 crédits)")
	fmt.Println("5.Sakado.Augmente l'espace de votre sac à dos.(500 crédits)")
	fmt.Printf("6.acheter spell 1 (100 crédits)")
	fmt.Println("7.acheter spell 2 (150 crédits)")
	fmt.Println("8.Acheter Jeremy qui vous suivra et se battra pour vous !(5000 crédits !)")
	fmt.Println("9. Retour")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		return
	}

	switch choice {

	case 1:

		if players.CheckInventory() == true {
			if players.PotionGratuite == true {
				players.PotionGratuite = false
				player.ItemToInventory("Potion de Soin", 1)
				fmt.Println("La Première est gratuite !.")
				time.Sleep(3 * time.Second)
			} else {
				if players.Credits >= 50 {
					player.ItemToInventory("Potion de Soin", 1)
					players.Credits -= 50
					fmt.Println("Vous avez acheté une potion.")
					time.Sleep(3 * time.Second)
				} else {
					fmt.Println("Vous n'avez pas assez de crédits.")
					time.Sleep(3 * time.Second)
				}
				VisitShop()
			}
		} else {
			color.Red("Plus assez d'espace dans votre inventaire.")
		}

	case 2:
		if players.CheckInventory() == true {
			if players.Credits >= 100 {
				players.Credits -= 100
				fmt.Println("Vous avez acheté une Potion de mana.")
				player.ItemToInventory("Potion de Mana", 1)
			} else {
				fmt.Println("Vous n'avez pas assez de crédits.")
			}
			time.Sleep(3 * time.Second)
			VisitShop()
		} else {
			color.Red("Plus assez d'espace dans votre inventaire.")
		}

	case 3:
		if players.CheckInventory() == true {
			if players.Credits >= 500 && players.UpInv < 3 {
				color.Green("Vous avez acheté un sac à main flambant neuf")
				players.InventoryMax += 5
				players.UpInv++
			} else {
				fmt.Println("Vous n'avez pas assez de crédits.")
			}
			time.Sleep(3 * time.Second)
		} else {
			color.Red("Plus assez d'espace dans votre inventaire.")
		}

	case 4:
		if players.CheckInventory() == true {
			if players.Credits >= 100 {

				player.ItemToInventory("Spell 1", 1)
				players.Credits -= 100
				color.Green("Vous avez acheté votre spell 1 !")
			} else {
				fmt.Println("Vous n'avez pas assez de crédits.")
			}
			time.Sleep(3 * time.Second)
			VisitShop()
		} else {
			color.Red("Plus assez d'espace dans votre inventaire.")
		}

	case 5:
		if players.CheckInventory() == true {
			if players.Credits >= 300 {
				player.ItemToInventory("Spell 2", 1)
				players.Credits -= 300
				color.Green("Vous avez acheté spell 2 !")
			} else {
				fmt.Println("Vous n'avez pas assez de crédits.")
			}
			time.Sleep(3 * time.Second)
			VisitShop()
		} else {
			color.Red("Plus assez d'espace dans votre inventaire.")
		}

	case 6:
		if players.CheckInventory() == true {
			if players.Credits >= 5000 {
				color.Green("Jérémie vous accompagne, il vous sera de grande aide !")
				player.ItemToInventory("Jérémie l'intrépide", 1)
				players.Credits -= 5000
			} else {
				fmt.Println("Vous n'avez pas assez de crédits.")
			}
			time.Sleep(3 * time.Second)
			VisitShop()

		} else {
			color.Red("Plus assez d'espace dans votre inventaire.")
		}

	}

	fmt.Println("———————————————————————————————————————————————")
}
