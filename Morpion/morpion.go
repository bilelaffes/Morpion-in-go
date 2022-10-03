package Morpion

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Init() {
	var (
		morpionTable                = [3][3]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
		caseTab, joueur, winner, null int = 0, 1, 0, 0
		err                     error
	)
	scanner := bufio.NewScanner(os.Stdin)

	for winner == 0 && null == 0 {

		afficheMorpion(morpionTable)
		fmt.Printf("Joueur %d entrez un nombre compris entre 1 à 9 : ", joueur)
		scanner.Scan()
		caseTab, err = strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Printf("Entrez un nombre et non autre chose !\n")
			continue
		}

		if (caseTab < 1) || (caseTab > 9) {
			fmt.Printf("Votre nombre doit être compris entre 1 à 9 !\n")
			continue
		}

		morpionTable = remplirMorpion(caseTab, joueur, morpionTable)
		winner = checkWinner(morpionTable, joueur)
		null = checkNull(morpionTable)
		switch joueur {
		case 1:
			joueur = 2
			break
		case 2:
			joueur = 1
			break
		}
	}

}

func afficheMorpion(morpion [3][3]string) {
	var afficheTab string

	for _, value := range morpion {
		for _, j := range value {
			afficheTab += " " + j
		}
		afficheTab += "\n"
	}
	fmt.Printf(afficheTab)
}

func remplirMorpion(index int, joueur int, morpion [3][3]string) [3][3]string {

	for index1, tab1 := range morpion {
		for index2, j := range tab1 {
			if val, _ := strconv.Atoi(j); val == index {
				if joueur == 1 {
					morpion[index1][index2] = "X"
					break
				}

				if joueur == 2 {
					morpion[index1][index2] = "O"
					break
				}
			}
		}
	}
	return morpion
}

func checkWinner(morpion [3][3]string, joueur int) int {

	for i := 0; i < 3; i++ {
		if (morpion[i][0] == morpion[i][1]) && (morpion[i][1] == morpion[i][2]) {
			fmt.Printf("Joueur %d, winner\n", joueur)
			return 1
		}

		if (morpion[0][i] == morpion[1][i]) && (morpion[1][i] == morpion[2][i]) {
			fmt.Printf("Joueur %d, winner\n", joueur)
			return 1
		}
	}

	if (morpion[0][0] == morpion[1][1]) && (morpion[1][1] == morpion[2][2]) {
		fmt.Printf("Joueur %d, winner\n", joueur)
		return 1
	}

	if (morpion[2][0] == morpion[1][1]) && (morpion[1][1] == morpion[0][2]) {
		fmt.Printf("Joueur %d, winner\n", joueur)
		return 1
	}
	return 0
}

func checkNull(morpion [3][3]string) int{
	
	for _, value := range morpion {
		for _, j := range value {
			if j != "X"  && j != "O" {
				return 0
			}
		}
	}
	fmt.Printf("MATCH NULL\n")
	return 1
}