package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"poker/models"
)

func main() {
	fmt.Println("hello world!")
}

func recieveTable(w http.ResponseWriter, r *http.Request) {

	table := models.Table{}
	err := json.NewDecoder(r.Body).Decode(&table)
	if err != nil {
		fmt.Println("there was an error reading the json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	winner, err := evaluateWinner(table)
	if err != nil {
		w.Write([]byte("invalid game state" + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte(`["` + winner.Name + `"]`))

	w.WriteHeader(http.StatusOK)
}

// makes refering to the card type much easier
type card = models.Card

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
		if win, _ := sevenCardCompare(append(table.FaceUp, best.Cards...), append(table.FaceUp, best.Cards...)); win == 1 {
			best = player
		}
	}
	return best, nil
}
