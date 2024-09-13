package main

import (
	"fmt"
	"github.com/fatih/color"
	"projetred/Menus/menu"
	"projetred/player"
)

func main() {
	var agentName string
	var agentChoice int

	fmt.Println("Bienvenue dans le jeu Projet RED!")
	color.White("———————————————————————————————————————————————")
	fmt.Println("Veuillez choisir votre agent:")
	color.Yellow("1. BRIM")
	color.Cyan("2. Sage")
	color.Green("3. Skye")
	color.Red("4. Phoenix")
	color.White("———————————————————————————————————————————————")
	_, err := fmt.Scan(&agentChoice)
	if err != nil {
		return
	}

	fmt.Println("Quel est votre nom  ?")
	_, err = fmt.Scan(&agentName)
	if err != nil {
		return
	}

	player.InitializePlayer(agentName, agentChoice)
	menu.Menu()
}
