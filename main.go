package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Player int

const (
	columns           = 7
	lines             = 6
	NullPlayer Player = -1
	Player1    Player = 0
	Player2    Player = 1
)
const Reset = "\033[0m"
const Red = "\033[31m"
const Blue = "\033[34m"

type Board [lines][columns]Player
type GameState struct {
	b      Board
	toPlay Player
}

func main() {
	var gs GameState
	for i := 0; i < lines; i++ {
		for j := 0; j < columns; j++ {
			gs.b[i][j] = NullPlayer
		}
	}

	for {
		n := gameIsOver(gs.b)
		if n != NullPlayer {
			printBoard(gs.b)
			fmt.Printf("Le joueur %d a gagné", n+1)

			break
		}
		if isNull(gs.b) {
			printBoard(gs.b)
			println("Partie nulle")
			break
		}
	TurnLoop:
		for {
			fmt.Println("État actuel du plateau :")
			printBoard(gs.b)
			switch takeTurn(&gs) {
			case -1:
				continue
			case 1:
				break TurnLoop
			}

		}

		gs.toPlay = 1 - gs.toPlay
	}
}

func gameIsOver(b Board) Player {
	if hasWon(Player1, b) {
		return Player1
	}
	if hasWon(Player2, b) {
		return Player2
	}
	return NullPlayer
}

func isNull(b Board) bool {
	for i := 0; i < lines; i++ {
		for j := 0; j < columns; j++ {
			if b[i][j] == NullPlayer {
				return false
			}
		}
	}
	return true
}
func printBoard(b Board) {
	println("| 0 | 1 | 2 | 3 | 4 | 5 | 6 |\n" + strings.Repeat("-", 30))
	for i := 0; i < lines; i++ {
		for j := 0; j < columns; j++ {
			c := ""
			switch b[i][j] {
			case NullPlayer:
				c = "."
				break
			case Player1:
				c = Red + "X" + Reset
				break
			case Player2:
				c = Blue + "O" + Reset
				break
			}
			fmt.Printf("| %s ", c)
		}
		print("|\n")
	}
}

func takeTurn(gs *GameState) int {
	fmt.Printf("Joueur %d, entrez un numéro de colonne :\n", gs.toPlay+1)
	var rep string
	_, err := fmt.Scanln(&rep)
	if err != nil {
		return -1
	}
	column, err := strconv.Atoi(rep)
	if err != nil {
		return -1
	}
	if column >= 0 && (column < columns) {
		for i := lines - 1; i >= 0; i-- {
			if gs.b[i][column] == NullPlayer {
				gs.b[i][column] = gs.toPlay
				return 1
			}
		}
	}
	return -1
}

func hasWon(p Player, b Board) bool {
	//check horizontal win
	for i := 0; i < lines; i++ {
	HorizontalScroll:
		for j := 0; j < columns-3; j++ {
			for k := 0; k < 4; k++ {
				if b[i][j] == p {
					j += 1
				} else {
					continue HorizontalScroll
				}
			}
			return true
		}
	}

	//check vertical win
	for i := 0; i < columns; i++ {
	VerticalScroll:
		for j := 0; j < lines-3; j++ {
			for k := 0; k < 4; k++ {
				if b[j][i] == p {
					j += 1
				} else {
					continue VerticalScroll
				}
			}
			return true
		}
	}

	//check diagonal win
	for i := 0; i < lines-3; i++ {
	DiagonalScroll1:
		for j := 0; j < columns-3; j++ {
			for k := 0; k < 4; k++ {
				if b[i+k][j+k] != p {
					continue DiagonalScroll1
				}
			}
			return true

		}
	}
	for i := lines - 1; i > 2; i-- {
	DiagonalScroll2:
		for j := 0; j < columns-3; j++ {
			for k := 0; k < 4; k++ {
				if b[i-k][j+k] != p {
					continue DiagonalScroll2
				}
			}
			return true
		}
	}

	return false
}
