package player

import "fmt"

type Player struct {
	Name       string
	Classe     string
	Health     int
	MAX_Health int
	Mana       int
	Attack     int
	Credits    int
	Potions    int
	Armor      string
	Agent      string
}

var player Player

func InitializePlayer(name string, agentChoice int) {
	player.Name = name
	player.Credits = 300
	player.Potions = 3
	player.Armor = "Armure basique"

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
