package main

import (
	"testing"
)

func TestSevenCardCompare(t *testing.T) {
	tests := []struct {
		left   []card
		right  []card
		winner int
		desc   string
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
			rightWins,
			"pair beats highcard",
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
			leftWins,
			"two pair beats pair",
		},
	}

	for _, test := range tests {
		result, _ := sevenCardCompare(test.left, test.right)

		if result != test.winner {
			t.Error(test.desc, "compare failed! recieved index", result, "\nfrom input", test)
		}
	}
}
