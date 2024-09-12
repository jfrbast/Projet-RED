package Menus

import (
	"fmt"
)

func ChoixAgent() {
	fmt.Println("Choisissez qui vous voulez être.")
	fmt.Println("1.Brimstone.")
	fmt.Println("2.Skye.")
	fmt.Println("3.Sage.")
	fmt.Println("4.Phoenix.")
	_, err := fmt.Scan(&ans)
	if err != nil {
		return
	}
	switch ans {
	case '1':

		fmt.Println("Vous avez choisis Brimstone.")
	case '2':

		fmt.Println("Vous avez choisis Skye.")
	case '3':
		fmt.Println("Vous avez choisis Sage.")
	case '4':
		fmt.Println("Vous avez choisis Phoenix.")
	default:
		fmt.Println("T'as pas compris la consigne, fais un effort stp.")
		ChoixAgent()
	}

}
