package main

import (
	"fmt"
	"os"
	"strings"
	"math/rand"
	"time"

	"jacksbrain.com/candyland/definitions"
)

func getPlayers() []definitions.Player {
	if (len(os.Args) == 1) {
		fmt.Println("Usage: ./candyland <playerOneName> [playerTwoName] [playerThreeName] [...]")
		os.Exit(1)
	}

	var players []definitions.Player

	for _, val := range os.Args[1:] {
		var playerObj =  definitions.Player{Name: val, Square: 0}
		players = append(players, playerObj)
	}

	return players
}

func GetNewDeck(deck []definitions.Card) []definitions.Card {
	cardsCopy := make([]definitions.Card, len(deck))
	copy(cardsCopy, deck)
	return cardsCopy
}

func DrawRandomCard(deck *[]definitions.Card) (definitions.Card, bool) {
	index := rand.Intn(len(*deck))
	
	// Remove and return the card at a random index
	card := (*deck)[index]
	*deck = append((*deck)[:index], (*deck)[index+1:]...)
	
	if len(*deck) == 0 {
		return card, false
	} else {
		return card, true
	}
}

func GetIndexOfNextColor(currentSquare int, color string) (int, bool) {
	if currentSquare < 0 || currentSquare >= len(definitions.BoardSquares) {
		return -1, false
	}

	for i := currentSquare; i < len(definitions.BoardSquares); i++ {
		if definitions.BoardSquares[i].Color == color {
			return i, true
		}
	}

	return -1, false
}

func GetIndexOfIcon(icon string) (int) {
	for index, element := range definitions.BoardSquares {
		if (element.IconSquare == icon) {
			return index
		}
	}

	return -1
}

func MovePlayerToNextSquareOfColor(player *definitions.Player, color string) {
	newSquare, wasFound := GetIndexOfNextColor(player.Square + 1, color)

	if (wasFound) {
		player.Square = newSquare
		fmt.Printf(" They advance to square %d.", newSquare)
	} else {
		fmt.Printf(" There are no more squares of that color, so %s WINS!\n", strings.ToUpper(player.Name))
		os.Exit(0)
	}

	if (definitions.BoardSquares[player.Square].ShortcutTarget != 0) {
		fmt.Printf(" This is a shortcut square! They take the %s to square %d.", definitions.BoardSquares[player.Square].ShortcutName, definitions.BoardSquares[player.Square].ShortcutTarget)
		player.Square = definitions.BoardSquares[player.Square].ShortcutTarget
	}
}

func main() {
	timeseed := time.Now().UnixNano()
	rand.Seed(timeseed)

	players := getPlayers()

	fmt.Println("Welcome to Candyland!\n")

	fmt.Println("Players:")
	for i, player := range players {
		fmt.Printf("%d. %s\n", i+1, player.Name)
	}

	fmt.Println("\nLet's play!")
	fmt.Println("\n=========================\n")

	hotDeck := GetNewDeck(definitions.Deck)
	turnCount := -1

	for true {
		turnCount += 1
		card, hasMore := DrawRandomCard(&hotDeck)

		fmt.Printf("Draw %d: %s has drawn a ", turnCount, players[turnCount % len(players)].Name)

		if (card.Icon != "") {
			fmt.Printf("%s card! They jump to square ", card.Icon)
			newSquare := GetIndexOfIcon(card.Icon)
			players[turnCount % len(players)].Square = newSquare
			fmt.Printf("%d. ", newSquare)
		} else {
			fmt.Printf("%s card.", card.Color)

			MovePlayerToNextSquareOfColor(&players[turnCount % len(players)], card.Color)
			
			if (card.Double) {
				fmt.Printf(" And it's a double!")
				MovePlayerToNextSquareOfColor(&players[turnCount % len(players)], card.Color)
			}
		}

		fmt.Printf("\n")

		if (!hasMore) {
			fmt.Printf("We've hit the end of the deck. Shuffling it again!\n")
			hotDeck = GetNewDeck(definitions.Deck)
		}
	}
}
