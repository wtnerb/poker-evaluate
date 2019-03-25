package main

import (
	"errors"

	models "github.com/wtnerb/poker-models"
)

// TODO: Replace this with interface. Additionally, logic will flex when
// model changes to new form
func evaluateWinner(table models.Table) ([]byte, error) {
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
		return nil, errors.New("there are no active players")
	}
	if len(activePlayers) == 1 {
		return []byte(activePlayers[0].ID), nil
	}

	// Create a list of players with the cards required to do the comparison.
	var players []p
	for _, player := range activePlayers {
		hand := buildBestHand(append(table.FaceUp, player.Cards...))
		players = append(players, p{hand, []byte(player.ID)})
	}

	best := players[0]
	for _, player := range players {
		if winner := sevenCardCompare(player.best, best.best); leftWins == winner {
			best = player
		}
	}
	return best.id, nil
}

type p struct {
	best bestHand
	id   []byte
}
