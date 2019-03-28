package main

import "testing"

func TestRankHand(t *testing.T) {
	tests := []struct {
		hand []card
		rank handRank
	}{
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(TWO, HEART),
				newCard(THREE, DIAMOND),
				newCard(FIVE, SPADE),
				newCard(SEVEN, SPADE),
				newCard(JACK, SPADE),
				newCard(QUEEN, HEART),
			}, pair,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(TWO, HEART),
				newCard(THREE, DIAMOND),
				newCard(FIVE, SPADE),
				newCard(QUEEN, SPADE),
				newCard(JACK, SPADE),
				newCard(TWO, DIAMOND),
			}, threeOfAKind,
		},
		{
			[]card{
				newCard(TWO, SPADE),
				newCard(TWO, HEART),
				newCard(THREE, SPADE),
				newCard(FIVE, SPADE),
				newCard(QUEEN, SPADE),
				newCard(JACK, SPADE),
				newCard(TWO, DIAMOND),
			}, flush,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(TWO, HEART),
				newCard(THREE, DIAMOND),
				newCard(FIVE, SPADE),
				newCard(FIVE, CLUB),
				newCard(JACK, SPADE),
				newCard(JACK, HEART),
			}, twoPair,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(FOUR, HEART),
				newCard(THREE, DIAMOND),
				newCard(SIX, SPADE),
				newCard(FIVE, CLUB),
				newCard(JACK, SPADE),
				newCard(JACK, HEART),
			}, straight,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(FOUR, HEART),
				newCard(THREE, DIAMOND),
				newCard(SIX, SPADE),
				newCard(FIVE, CLUB),
				newCard(JACK, SPADE),
				newCard(FIVE, HEART),
			}, straight,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(FOUR, CLUB),
				newCard(THREE, CLUB),
				newCard(SIX, CLUB),
				newCard(FIVE, CLUB),
				newCard(JACK, SPADE),
				newCard(FIVE, DIAMOND),
			}, straightFlush,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(FOUR, CLUB),
				newCard(THREE, CLUB),
				newCard(SIX, SPADE),
				newCard(FIVE, CLUB),
				newCard(JACK, CLUB),
				newCard(FIVE, DIAMOND),
			}, flush,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(FOUR, HEART),
				newCard(THREE, DIAMOND),
				newCard(ACE, SPADE), //testing that Ace's will switch to a one when needed
				newCard(FIVE, CLUB),
				newCard(JACK, SPADE),
				newCard(FIVE, HEART),
			}, straight,
		},
	}

	for _, test := range tests {
		actual := rankHand(test.hand)
		if actual != test.rank {
			t.Errorf("Hand ranked incorrectly. %v\nranked as %v should be %v", test.hand, actual, test.rank)
		}
	}
}
