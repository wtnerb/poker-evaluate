package main

import "testing"

func TestHoldemCompare(t *testing.T) {
	test := struct {
		tableCards []card
		holes      [][]card
		expected   struct {
			winner int
			hand   bestHand
		}
	}{
		[]card{
			card{ACE, SPADE},
			card{KING, DIAMOND},
			card{TWO, HEART},
			card{SEVEN, HEART},
			card{TEN, CLUB},
		},
		[][]card{
			[]card{card{FOUR, SPADE}, card{QUEEN, CLUB}},
			[]card{card{THREE, SPADE}, card{KING, SPADE}},
		},
		struct {
			winner int
			hand   bestHand
		}{
			1,
			bestHand{
				[5]card{
					card{ACE, SPADE},
					card{KING, DIAMOND},
					card{KING, SPADE},
					card{TEN, CLUB},
					card{SEVEN, HEART},
				},
				pair,
			},
		},
	}
	winner, best := holdemCompare(test.tableCards, test.holes)
	if winner != test.expected.winner {
		t.Error("the wrong person won")
	}
	if len(best.cards) != 5 || len(test.expected.hand.cards) != 5 {
		t.Error("wrong number of cards")
	}
}
