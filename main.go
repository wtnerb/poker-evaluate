package main

import (
	"errors"
	"fmt"
	"poker/models"
)

func main() {
	fmt.Println("hello world!")
}

type card = models.Card

func evaluateSingleWinner(table models.Table) (models.TablePlayer, error) {
	// if table.Players[0] == nil {
	// 	panic("there are no players at the table!")
	// }

	var activePlayers []models.TablePlayer

	for _, player := range table.Players {
		if !player.Folded {
			activePlayers = append(activePlayers, player)
		}
	}
	if len(activePlayers) == 0 {
		return models.TablePlayer{}, errors.New("there are no active players")
	}
	if len(activePlayers) == 1 {
		return activePlayers[0], nil
	}
	return models.TablePlayer{}, errors.New("the winner could not be determined")
}

func evaluateHands(faceUp []card, holeCards [][2]card) int {
	return 1
}
