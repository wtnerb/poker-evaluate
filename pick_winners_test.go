package main

import (
	"testing"

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
						Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: true,
					},
					models.TablePlayer{
						Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
					},
				},
				FaceUp: []card{newCard(FIVE, SPADE), newCard(JACK, HEART), newCard(KING, CLUB), newCard(SEVEN, DIAMOND), newCard(NINE, DIAMOND)},
			},
			models.TablePlayer{
				Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
			},
		},
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: false,
					},
					models.TablePlayer{
						Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
					},
				},
				FaceUp: []card{newCard(FIVE, SPADE), newCard(JACK, HEART), newCard(KING, CLUB), newCard(SEVEN, DIAMOND), newCard(NINE, DIAMOND)},
			},
			models.TablePlayer{
				Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: false,
			},
		},
	}
	for _, test := range tests {
		actual, err := evaluateWinner(test.table)
		if err != nil {
			t.Error(err)
		}
		if actual.Name != test.expected.Name {
			t.Error("wrong person won. Expected", test.expected.Name, "got", actual.Name, "\n", test.table)
		}
	}
}
