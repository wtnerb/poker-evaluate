package main

import (
	"testing"
)

func (v verdict) String() string {
	if v == leftWins {
		return "left won"
	}
	if v == rightWins {
		return "right won"
	}
	if v == tie {
		return "tie"
	}
	return "invalid verdict"
}

func TestSevenCardCompare(t *testing.T) {
	tests := []struct {
		left   []card
		right  []card
		winner verdict
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

func TestCompareBest(t *testing.T) {
	tests := []struct {
		left   bestHand
		right  bestHand
		winner verdict
		desc   string
	}{
		{
			bestHand{
				[5]card{
					newCard(KING, SPADE),
					newCard(JACK, SPADE),
					newCard(TEN, CLUB),
					newCard(TWO, DIAMOND),
					newCard(FOUR, HEART),
				},
				highcard,
			},
			bestHand{
				[5]card{
					newCard(JACK, SPADE),
					newCard(TEN, CLUB),
					newCard(EIGHT, SPADE),
					newCard(TWO, DIAMOND),
					newCard(FOUR, HEART),
				},
				highcard,
			},
			leftWins,
			"King is larger highcard than jack",
		},
	}

	for _, test := range tests {
		result := compareBest(test.left, test.right)
		if result != test.winner {
			t.Error(test.desc, "\nwrong person won. got:", result, "expected:", test.winner, "\n", test.left.cards, "\n", test.right.cards)
		}
	}
}
