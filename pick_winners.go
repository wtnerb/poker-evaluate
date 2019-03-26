package main

import (
	"errors"

	models "github.com/wtnerb/poker-models"
)

// TODO: Replace this with interface. Additionally, logic will flex when
// model changes to new form
func evaluateWinner(table models.Table) (RespObj, error) {
	// if table.Players[0] == nil {
	// 	panic("there are no players at the table!")
	// }
	var activePlayers []models.TablePlayer
	var w RespObj

	for _, player := range table.Players {
		if !player.Folded {
			activePlayers = append(activePlayers, player)
		}
	}
	if len(activePlayers) == 0 {
		return w, errors.New("there are no active players")
	}
	if len(activePlayers) == 1 {
		w.Ids = [][]byte{[]byte(activePlayers[0].ID)}
		return w, nil
	}

	// Create a list of players with the cards required to do the comparison.
	pb := makeBest(table.FaceUp, table.Players)
	best := pickWinners(pb)
	w.Ids = listIds(best)
	return w, nil
}

type p struct {
	best bestHand
	id   []byte
}

func makeBest(faceUp []card, players []models.TablePlayer) []p {
	var playersBest []p
	for _, player := range players {
		hand := buildBestHand(append(faceUp, player.Cards...))
		playersBest = append(playersBest, p{hand, []byte(player.ID)})
	}
	return playersBest
}

func listIds(players []p) (ids [][]byte) {
	for _, player := range players {
		ids = append(ids, player.id)
	}
	return
}

func pickWinners(players []p) []p {
	winners := []p{p{best: bestHand{rank: handRank(-1)}}}
	for _, player := range players {
		winner := sevenCardCompare(player.best, winners[0].best)
		switch winner {
		case leftWins:
			winners = []p{player}
		case tie:
			winners = append(winners, player)
		}
	}
	return winners
}
