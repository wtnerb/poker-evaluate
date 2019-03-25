package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/wtnerb/poker-models"
)

const port = ":4002"

func main() {
	fmt.Println("hello world!")
	http.HandleFunc("/", recieveTable)

	http.ListenAndServe(port, nil)
}

func recieveTable(w http.ResponseWriter, r *http.Request) {
	table := models.Table{}
	err := json.NewDecoder(r.Body).Decode(&table)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("json is malformed"))
		return
	}

	winner, err := evaluateWinner(table)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid game state " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(winner)
}

// makes refering to the card type much easier
type card = models.Card
