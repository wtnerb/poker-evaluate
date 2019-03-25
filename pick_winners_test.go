package main

import (
	"bytes"
	"testing"

	"gopkg.in/mgo.v2/bson"

	models "github.com/wtnerb/poker-models"
)

func TestWinner(t *testing.T) {
	// At the moment, this is a dummy test. Started working, realized
	// this was too big to be a single unit
	tests := []struct {
		table    models.Table
		expected models.TablePlayer
	}{
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e375"),
						Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: true,
					},
					models.TablePlayer{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e333"),
						Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
					},
				},
				FaceUp: []card{newCard(FIVE, SPADE), newCard(JACK, HEART), newCard(KING, CLUB), newCard(SEVEN, DIAMOND), newCard(NINE, DIAMOND)},
			},
			models.TablePlayer{
				ID:   bson.ObjectIdHex("5c6482805508c93011b4e333"),
				Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
			},
		},
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e375"),
						Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: false,
					},
					models.TablePlayer{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e333"),
						Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
					},
				},
				FaceUp: []card{newCard(FIVE, SPADE), newCard(JACK, HEART), newCard(KING, CLUB), newCard(SEVEN, DIAMOND), newCard(NINE, DIAMOND)},
			},
			models.TablePlayer{
				ID:   bson.ObjectIdHex("5c6482805508c93011b4e375"),
				Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: false,
			},
		},
	}
	for _, test := range tests {
		actual, err := evaluateWinner(test.table)
		if err != nil {
			t.Error(err)
		}
		if diff := bytes.Compare([]byte(test.expected.ID), actual); diff != 0 {
			t.Error("wrong person won. Expected", test.expected.ID, "got", actual, "\n", test.table)
		}
	}
}
