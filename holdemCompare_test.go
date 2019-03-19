package main

import (
	"testing"
)

// func TestHoldemCompare(t *testing.T) {
// 	test := struct {
// 		tableCards []card
// 		holes      [][]card
// 		expected   struct {
// 			winner int
// 			hand   bestHand
// 		}
// 	}{
// 		[]card{
// 			newCard(ACE, SPADE),
// 			newCard(KING, DIAMOND),
// 			newCard(TWO, HEART),
// 			newCard(SEVEN, HEART),
// 			newCard(TEN, CLUB),
// 		},
// 		[][]card{
// 			[]card{newCard(FOUR, SPADE}, newCard(QUEEN, CLUB}),
// 			[]card{newCard(THREE, SPADE}, newCard(KING, SPADE}),
// 		},
// 		struct {
// 			winner int
// 			hand   bestHand
// 		}{
// 			1,
// 			bestHand{
// 				[5]models.NewCar)(
// 					newCard(ACE, SPADE),
// 					newCard(KING, DIAMOND),
// 					newCard(KING, SPADE),
// 					newCard(TEN, CLUB),
// 					newCard(SEVEN, HEART),
// 				},
// 				pair,
// 			},
// 		},
// 	}
// 	winner, best := holdemCompare(test.tableCards, test.holes)
// 	if winner != test.expected.winner {
// 		t.Error("the wrong person won")
// 	}
// 	if len(best.cards) != 5 || len(test.expected.hand.cards) != 5 {
// 		t.Error("wrong number of cards")
// 	}
// }

func TestSevenCardCompare(t *testing.T) {
	tests := []struct {
		left   []card
		right  []card
		winner int
	}{
		{
			[]card{
				newCard(ACE, SPADE),
				newCard(KING, DIAMOND),
				newCard(TWO, HEART),
				newCard(SEVEN, HEART),
				newCard(TEN, CLUB),
				newCard(FOUR, SPADE),
				newCard(QUEEN, CLUB),
			},
			[]card{
				newCard(ACE, SPADE),
				newCard(KING, DIAMOND),
				newCard(TWO, HEART),
				newCard(SEVEN, HEART),
				newCard(TEN, CLUB),
				newCard(EIGHT, DIAMOND),
				newCard(KING, CLUB),
			},
			RIGHT,
		},
		{
			[]card{
				newCard(ACE, SPADE),
				newCard(SEVEN, DIAMOND),
				newCard(TWO, HEART),
				newCard(SEVEN, HEART),
				newCard(TEN, CLUB),
				newCard(FOUR, SPADE),
				newCard(ACE, CLUB),
			},
			[]card{
				newCard(ACE, SPADE),
				newCard(KING, DIAMOND),
				newCard(TWO, HEART),
				newCard(SEVEN, HEART),
				newCard(TEN, CLUB),
				newCard(EIGHT, DIAMOND),
				newCard(KING, CLUB),
			},
			LEFT,
		},
	}

	for _, test := range tests {
		result, _ := sevenCardCompare(test.left, test.right)

		if result != test.winner {
			t.Error("compare failed! recieved index", result, "\nfrom input", test)
		}
	}
}
