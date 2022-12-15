package Morpion

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	table               = [3][3]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	playerSymbol        = [2]string{"X", "O"}
	cellTab, player int = 0, 0
)

func Init() {
	var (
		winner, draw bool = false, false
		err          error
	)
	scanner := bufio.NewScanner(os.Stdin)

	for !winner && !draw {

		displayTable()
		fmt.Printf("Player %d enter a number between 1 to 9 : ", player)
		scanner.Scan()
		cellTab, err = strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Enter a number and not something else !")
			continue
		}

		if (cellTab < 1) || (cellTab > 9) {
			fmt.Println("Your number must be between 1 to 9 !")
			continue
		}

		fillTable()
		winner = checkWinner()
		draw = checkDraw()
		player = (player + 1) % 2
	}

}

func displayTable() {
	var tableDisplay string

	for _, value := range table {
		for _, j := range value {
			tableDisplay += " " + j
		}
		tableDisplay += "\n"
	}
	fmt.Printf(tableDisplay)
}

func fillTable() {

	for index1, tab1 := range table {
		for index2, j := range tab1 {
			if val, _ := strconv.Atoi(j); val == cellTab {
				table[index1][index2] = playerSymbol[player]
				break // In order not to go through the whole array
			}
		}
	}
}

func checkWinner() bool {

	for i := 0; i < 3; i++ {
		if (table[i][0] == table[i][1]) && (table[i][1] == table[i][2]) {
			fmt.Printf("Player %d, winner\n", player)
			return true
		}

		if (table[0][i] == table[1][i]) && (table[1][i] == table[2][i]) {
			fmt.Printf("Player %d, winner\n", player)
			return true
		}
	}

	if (table[0][0] == table[1][1]) && (table[1][1] == table[2][2]) {
		fmt.Printf("Player %d, winner\n", player)
		return true
	}

	if (table[2][0] == table[1][1]) && (table[1][1] == table[0][2]) {
		fmt.Printf("Player %d, winner\n", player)
		return true
	}
	return false
}

func checkDraw() bool {

	for _, value := range table {
		for _, j := range value {
			if j != "X" && j != "O" {
				return false
			}
		}
	}
	fmt.Println("Draw")
	return true
}
