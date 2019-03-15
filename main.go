package main

import (
	"errors"
	"fmt"
	"poker/models"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(card{2, 1}, card{2, 2}, card{2, 3}, card{2, 4})
}

// makes refering to the card type much easier
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

// looks at hands from texas hold'em and returns the ranks of each hand
func rankHands(faceUp []card, holeCards [][2]card) []handRank {
	ranks := make([]handRank, len(holeCards))
	for i := range holeCards {
		ranks[i] = rankHand(append(faceUp, holeCards[i][0], holeCards[i][1]))
	}
	return ranks
}
