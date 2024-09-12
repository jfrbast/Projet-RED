package projetred

import "fmt"

type Item struct {
	Name string
}

var inventory = make([]Item, 10)

// AddItem adds an item to the inventory
func AddItem(index int, name string) {
	if index >= 0 && index < len(inventory) {
		inventory[index] = Item{Name: name}
		fmt.Printf("Added %s to slot %d\n", name, index)
	} else {
		fmt.Println("Index out of range!")
	}
}

func RemoveItem(index int) {
	if index >= 0 && index < len(inventory) {
		fmt.Printf("Removed %s from slot %d\n", inventory[index].Name, index)
		inventory[index] = Item{} // Clear the slot
	} else {
		fmt.Println("Index out of range!")
	}
}

func ShowInventory() {
	for i, item := range inventory {
		if item.Name != "" {
			fmt.Printf("Item %d: %s\n", i, item.Name)
		} else {
			fmt.Printf("Item %d: Vide\n", i)
		}
	}
}

//1. Tableau dont la taile ne varie pas 10 max  + Si inventaire est vide print vide sinon print tout l inventaire
