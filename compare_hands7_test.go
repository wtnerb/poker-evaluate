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
		result := sevenCardCompare(test.left, test.right)

		if result != test.winner {
			t.Error(test.desc, "compare failed! recieved index", result, "\nfrom input", test)
		}
	}
}

func TestCompareBest(t *testing.T) {
	tests := []struct {
		left   [5]card
		right  [5]card
		rank   handRank
		winner verdict
		desc   string
	}{
		{
			[5]card{
				newCard(KING, SPADE),
				newCard(JACK, HEART),
				newCard(TEN, SPADE),
				newCard(TWO, SPADE),
				newCard(FOUR, CLUB),
			},

			[5]card{
				newCard(JACK, SPADE),
				newCard(TEN, CLUB),
				newCard(EIGHT, SPADE),
				newCard(TWO, DIAMOND),
				newCard(FOUR, HEART),
			},
			highcard,
			leftWins,
			"King is larger highcard than jack",
		},
		{
			[5]card{
				newCard(KING, SPADE),
				newCard(JACK, SPADE),
				newCard(TEN, CLUB),
				newCard(TWO, DIAMOND),
				newCard(FOUR, HEART),
			},

			[5]card{
				newCard(JACK, HEART),
				newCard(TEN, DIAMOND),
				newCard(EIGHT, SPADE),
				newCard(KING, DIAMOND),
				newCard(FOUR, CLUB),
			},
			highcard,
			rightWins,
			"Eight is larger than four",
		},
		{
			[5]card{
				newCard(KING, SPADE),
				newCard(JACK, SPADE),
				newCard(TEN, CLUB),
				newCard(TEN, DIAMOND),
				newCard(FOUR, HEART),
			},

			[5]card{
				newCard(JACK, HEART),
				newCard(TEN, DIAMOND),
				newCard(EIGHT, SPADE),
				newCard(EIGHT, DIAMOND),
				newCard(FOUR, CLUB),
			},
			pair,
			leftWins,
			"ten pair is greater eight pair",
		},
		{
			[5]card{
				newCard(KING, SPADE),
				newCard(JACK, SPADE),
				newCard(TEN, CLUB),
				newCard(TEN, HEART),
				newCard(FOUR, HEART),
			},

			[5]card{
				newCard(JACK, HEART),
				newCard(TEN, DIAMOND),
				newCard(TEN, SPADE),
				newCard(EIGHT, DIAMOND),
				newCard(FOUR, CLUB),
			},
			pair,
			leftWins,
			"Equal pairs, king highcard wins",
		},
		{
			[5]card{
				newCard(KING, SPADE),
				newCard(JACK, SPADE),
				newCard(TEN, CLUB),
				newCard(TEN, HEART),
				newCard(FOUR, HEART),
			},

			[5]card{
				newCard(JACK, HEART),
				newCard(TEN, DIAMOND),
				newCard(TEN, SPADE),
				newCard(EIGHT, DIAMOND),
				newCard(KING, CLUB),
			},
			pair,
			rightWins,
			"Equal pairs, eight beats four",
		},
		{
			[5]card{
				newCard(KING, SPADE),
				newCard(TEN, SPADE),
				newCard(TEN, CLUB),
				newCard(TEN, HEART),
				newCard(FOUR, HEART),
			},

			[5]card{
				newCard(TEN, HEART),
				newCard(TEN, DIAMOND),
				newCard(TEN, SPADE),
				newCard(EIGHT, DIAMOND),
				newCard(KING, CLUB),
			},
			threeOfAKind,
			rightWins,
			"Equal Three of a kind, eight beats four",
		},
		{
			[5]card{
				newCard(KING, SPADE),
				newCard(TEN, SPADE),
				newCard(TEN, CLUB),
				newCard(TEN, HEART),
				newCard(KING, HEART),
			},

			[5]card{
				newCard(TEN, HEART),
				newCard(TEN, DIAMOND),
				newCard(TEN, SPADE),
				newCard(EIGHT, DIAMOND),
				newCard(EIGHT, CLUB),
			},
			fullHouse,
			leftWins,
			"Kings full of tens beats eights full of tens",
		},
		{
			[5]card{
				newCard(KING, SPADE),
				newCard(TEN, SPADE),
				newCard(TEN, CLUB),
				newCard(TEN, HEART),
				newCard(KING, HEART),
			},

			[5]card{
				newCard(JACK, HEART),
				newCard(JACK, DIAMOND),
				newCard(JACK, SPADE),
				newCard(EIGHT, DIAMOND),
				newCard(EIGHT, CLUB),
			},
			fullHouse,
			rightWins,
			"eights full of jacks beats kings full of tens",
		},
		{
			[5]card{
				newCard(KING, SPADE),
				newCard(TEN, SPADE),
				newCard(NINE, CLUB),
				newCard(TEN, HEART),
				newCard(KING, HEART),
			},

			[5]card{
				newCard(JACK, HEART),
				newCard(JACK, DIAMOND),
				newCard(NINE, SPADE),
				newCard(TEN, DIAMOND),
				newCard(TEN, CLUB),
			},
			twoPair,
			leftWins,
			"Kings are better than jacks",
		},
		{
			[5]card{
				newCard(NINE, SPADE),
				newCard(TEN, SPADE),
				newCard(NINE, CLUB),
				newCard(TEN, HEART),
				newCard(KING, HEART),
			},

			[5]card{
				newCard(EIGHT, HEART),
				newCard(EIGHT, DIAMOND),
				newCard(NINE, SPADE),
				newCard(TEN, DIAMOND),
				newCard(TEN, CLUB),
			},
			twoPair,
			leftWins,
			"Nines are better than eights",
		},
		{
			[5]card{
				newCard(NINE, SPADE),
				newCard(TEN, SPADE),
				newCard(NINE, CLUB),
				newCard(TEN, HEART),
				newCard(FIVE, HEART),
			},

			[5]card{
				newCard(NINE, HEART),
				newCard(EIGHT, DIAMOND),
				newCard(NINE, SPADE),
				newCard(TEN, DIAMOND),
				newCard(TEN, CLUB),
			},
			twoPair,
			rightWins,
			"Eight is better than five",
		},
		{
			[5]card{
				newCard(KING, SPADE),
				newCard(JACK, SPADE),
				newCard(QUEEN, CLUB),
				newCard(TEN, HEART),
				newCard(NINE, HEART),
			},

			[5]card{
				newCard(ACE, HEART),
				newCard(TWO, DIAMOND),
				newCard(THREE, SPADE),
				newCard(FOUR, DIAMOND),
				newCard(FIVE, CLUB),
			},
			straight,
			leftWins,
			"A-5 straight less than 9-K straight",
		},
	}

	for _, test := range tests {
		result := compareBest(test.left, test.right, test.rank)
		if result != test.winner {
			t.Error(test.desc, "\nwrong person won. got:", result, "expected:", test.winner, "\n", test.left, "\n", test.right)
		}
	}
}
