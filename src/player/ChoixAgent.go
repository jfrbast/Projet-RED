package player

import "fmt"

type Item struct {
	Name     string
	Quantity int
}
type Spells struct {
	SpellName string
	Quantity  int
}

type Player struct {
	Name       string
	Classe     string
	Health     int
	MAX_Health int
	Mana       int
	Attack     int
	Credits    int
	Agent      string
	Inventory  []Item
	Spells     []Spells
	helmet     bool
	boots      bool
	plastron   bool
}

var player Player

func InitializePlayer(name string, agentChoice int) {
	player.Name = name
	player.Credits = 300
	player.Inventory = []Item{{Name: "Potion de Soin", Quantity: 2}}
	player.Inventory = []Item{{Name: "Spell 1", Quantity: 0}}
	player.Inventory = []Item{{Name: "Spell 2", Quantity: 0}}
	player.boots = false
	player.plastron = false
	player.helmet = false

	switch agentChoice {
	case 1:
		player.Agent = "BRIM"
		player.Classe = "Smoker"
		player.Health = 120
		player.MAX_Health = 120
		player.Mana = 50
		player.Attack = 10
	case 2:
		player.Agent = "Sage"
		player.Health = 100
		player.MAX_Health = 100
		player.Mana = 70
		player.Attack = 8
	case 3:
		player.Agent = "Skye"
		player.Health = 110
		player.MAX_Health = 110
		player.Mana = 60
		player.Attack = 9
	case 4:
		player.Agent = "Phoenix"
		player.Health = 90
		player.MAX_Health = 90
		player.Mana = 80
		player.Attack = 12
	default:
		player.Agent = "Inconnu"
	}

	fmt.Printf("Vous avez choisi l'agent %s avec le nom %s.\n", player.Agent, player.Name)
}

func GetPlayer() *Player {
	return &player
}

func UseItem(itemName string) bool {
	for i, item := range player.Inventory {
		if item.Name == itemName && item.Quantity > 0 {
			player.Inventory[i].Quantity -= 1
			return true
		}
	}
	return false
}

func AddItemToInventory(itemName string, quantity int) {
	for i, item := range player.Inventory {
		if item.Name == itemName {
			player.Inventory[i].Quantity += quantity
			return
		}
	}
	player.Inventory = append(player.Inventory, Item{Name: itemName, Quantity: quantity})
}
