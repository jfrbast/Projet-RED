package projetred

import (
	"fmt"
)

func main() {
	inventory := NewInventory()

	for {

		fmt.Println("\nMenu principal :")
		fmt.Println("1. Ajouter un objet à l'inventaire")
		fmt.Println("2. Retirer un objet de l'inventaire")
		fmt.Println("3. Afficher le contenu de l'inventaire")
		fmt.Println("4. Quitter")

		var choice int
		fmt.Print("Choisissez une option : ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Entrée invalide. Veuillez entrer un nombre.")
			continue
		}

		switch choice {
		case 1:
			// Ajouter un objet
			var name, description string
			fmt.Print("Entrez le nom de l'objet : ")
			fmt.Scan(&name)
			fmt.Print("Entrez la description de l'objet : ")
			fmt.Scan(&description)
			item := NewItem(name, description)
			inventory.AddItem(item)

		case 2:
			var name string
			fmt.Print("Entrez le nom de l'objet à retirer : ")
			fmt.Scan(&name)
			inventory.RemoveItem(name)

		case 3:

			inventory.Display()

		case 4:

			fmt.Println("Quitter le programme.")
			return

		default:

			fmt.Println("Choix non valide. Essayez encore.")
		}
	}
}
