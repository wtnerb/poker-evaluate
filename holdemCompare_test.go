package main

import "testing"

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
// 			card{ACE, SPADE},
// 			card{KING, DIAMOND},
// 			card{TWO, HEART},
// 			card{SEVEN, HEART},
// 			card{TEN, CLUB},
// 		},
// 		[][]card{
// 			[]card{card{FOUR, SPADE}, card{QUEEN, CLUB}},
// 			[]card{card{THREE, SPADE}, card{KING, SPADE}},
// 		},
// 		struct {
// 			winner int
// 			hand   bestHand
// 		}{
// 			1,
// 			bestHand{
// 				[5]card{
// 					card{ACE, SPADE},
// 					card{KING, DIAMOND},
// 					card{KING, SPADE},
// 					card{TEN, CLUB},
// 					card{SEVEN, HEART},
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
				card{ACE, SPADE},
				card{KING, DIAMOND},
				card{TWO, HEART},
				card{SEVEN, HEART},
				card{TEN, CLUB},
				card{FOUR, SPADE},
				card{QUEEN, CLUB},
			},
			[]card{
				card{ACE, SPADE},
				card{KING, DIAMOND},
				card{TWO, HEART},
				card{SEVEN, HEART},
				card{TEN, CLUB},
				card{EIGHT, DIAMOND},
				card{KING, CLUB},
			},
			RIGHT,
		},
		{
			[]card{
				card{ACE, SPADE},
				card{SEVEN, DIAMOND},
				card{TWO, HEART},
				card{SEVEN, HEART},
				card{TEN, CLUB},
				card{FOUR, SPADE},
				card{ACE, CLUB},
			},
			[]card{
				card{ACE, SPADE},
				card{KING, DIAMOND},
				card{TWO, HEART},
				card{SEVEN, HEART},
				card{TEN, CLUB},
				card{EIGHT, DIAMOND},
				card{KING, CLUB},
			},
			LEFT,
		},
	}

	for _, test := range tests {
		result := sevenCardCompare(test.left, test.right)

		if result != test.winner {
			t.Error("compare failed! recieved index", result, "\nfrom input", test)
		}
	}
}
