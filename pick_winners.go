package main

import (
	"errors"

	models "github.com/wtnerb/poker-models"
)

// TODO: Replace this with interface. Additionally, logic will flex when
// model changes to new form
func evaluateWinner(table models.Table) (models.TablePlayer, error) {
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

	best := activePlayers[0]
	for _, player := range activePlayers {
		if win := sevenCardCompare(append(table.FaceUp, best.Cards...), append(table.FaceUp, best.Cards...)); win == 1 {
			best = player
		}
	}
	return best, nil
}
