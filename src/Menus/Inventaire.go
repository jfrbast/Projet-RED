package Menus

import "fmt"

type Item struct {
	Name string
}

func ShowInventory() {

	inventory := []Item{
		{Name: "casque"},
		{Name: "placetron"},
		{Name: "Jambieres"},
		{Name: "bottes"},
		{Name: "potions de soins"},
		{Name: "potions de poisons"},
		{Name: ""},
		{Name: ""},
		{Name: ""},
		{Name: ""},
	}

	for i, item := range inventory {
		if item.Name != "" {
			fmt.Printf("Item %d: %s\n", i, item.Name)
		} else {
			fmt.Printf("Item %d: Vide\n", i)
		}
	}
}

//1. Tableau dont la taille ne varie pas 10 max  + Si inventaire est vide print vide sinon print tout l inventaire
