package player

import (
	"fmt"
)

type Item struct {
	Name     string
	Quantity int
}
type Spells struct {
	SpellName string
	Quantity  int
}

type Player struct {
	Name           string
	Classe         string
	Health         int
	MAX_Health     int
	Mana           int
	Attack         int
	Credits        int
	Agent          string
	Inventory      []Item
	Spells         [2]Spells
	helmet         bool
	boots          bool
	chestplate     bool
	PotionGratuite bool
	XP             int
	Level          int
	Xplvl          int
	InventoryMax   int
	UpInv          int
	Initiative     int
}

var player Player

func InitializePlayer(name string, agentChoice int) {
	player.Name = name
	player.Credits = 3000000000
	ItemToInventory("Potion de Soin", 2)
	ItemToInventory("Spell 1", 1)
	ItemToInventory("Spell 2", 1)
	player.boots = false
	player.chestplate = false
	player.helmet = false
	player.PotionGratuite = true
	player.InventoryMax = 10
	player.XP = 0
	player.Level = 1
	player.Xplvl = 10
	player.UpInv = 0
	player.Initiative = 10
	if player.XP == player.Xplvl {

		player.Xplvl += 5
		player.Level += 1
		player.XP = 0
	}

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

func ItemToInventory(itemName string, quantity int) {
	for i, item := range player.Inventory {
		if item.Name == itemName {
			player.Inventory[i].Quantity += quantity
			return
		}
	}
	player.Inventory = append(player.Inventory, Item{Name: itemName, Quantity: quantity})
}
func (p *Player) CheckInventory() bool {
	count := 0
	for _, items := range p.Inventory {
		count = count + items.Quantity
		if count >= p.InventoryMax {
			return true
		}
	}
	fmt.Println(count, "/", p.InventoryMax)
	return false
}
func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 97 && s[i] <= 122 || s[i] >= 65 && s[i] <= 90) {
			return false
		}
	}
	return true
}
func ToLower(s string) string {
	result := ""
	for i, char := range s {
		if i >= 1 {
			if char >= 'A' && char <= 'Z' {
				result += string(char + 32)
			} else {
				result += string(char)
			}
		} else {
			result += string(char)
		}

	}
	return result
}
