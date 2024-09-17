package main

import (
	"fmt"
	"github.com/fatih/color"
	src "projetred"
	"projetred/Menus/menu"
	"projetred/player"
	"strconv"
	"strings"
	"time"
)

func main() {
	var agentName string
	var agentChoice string
	fmt.Println('\n',
		`
									( ___ )                                                ( ___ )
									 |   |~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|   | 
									 |   |     __     __    _                       _       |   | 
									 |   |     \ \   / /_ _| | ___  _ __ __ _ _ __ | |_     |   | 
									 |   |      \ \ / / _  | |/ _ \|  __/ _  | '_ \| __|    |   |
									 |   |       \ V / (_| | | (_) | | | (_| | | | | |_     |   |
									 |   |        \_/ \__,_|_|\___/|_|  \__,_|_| |_|\__|    |   |
									 |   |                  _____ ____ _____                |   |
									 |   |                 |_   _| __ )_   _|               |   |
									 |   |                   | | |  _ \ | |                 |   |
									 |   |                   | | | |_) || |                 |   |
									 |   |                   |_| |____/ |_|                 |   |
									 |___|~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~|___|
									(_____)                                                (_____)`)

	fmt.Println("Quel est votre nom  ?")
	_, err := fmt.Scan(&agentName)
	if err != nil {
		return
	}

	if player.IsAlpha(agentName) == false {
		main()
	}
	agentName = strings.ToUpper(agentName)
	agentName = player.ToLower(agentName)
	fmt.Println("Bienvenue agent", agentName)

	var agentChoiceNbr int
	for {
		color.White("———————————————————————————————————————————————")
		fmt.Println("Veuillez choisir votre agent:")
		color.Yellow("1. BRIM")
		color.Cyan("2. Sage")
		color.Green("3. Skye")
		color.Red("4. Phoenix")
		color.White("———————————————————————————————————————————————")
		_, err = fmt.Scan(&agentChoice)
		color.Green("VOus avez choisis", agentChoice)
		time.Sleep(2 * time.Second)
		agentChoiceNbr, err = strconv.Atoi(agentChoice)
		if err != nil || !(1 <= agentChoiceNbr && agentChoiceNbr <= 4) {
			fmt.Println("Chef ?")
			time.Sleep(2 * time.Second)
			agentChoice = ""
			src.ClearScreen()
			continue
		}
		break
	}

	player.InitializePlayer(agentName, agentChoiceNbr)
	menu.Menu()
}
