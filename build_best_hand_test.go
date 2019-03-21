func TestBestHand(t *testing.T) {
	tests := []struct {
		hand []card
		best bestHand
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
			}, bestHand{
				[5]card{
					newCard(TWO, CLUB),
					newCard(TWO, HEART),
					newCard(QUEEN, HEART),
					newCard(JACK, SPADE),
					newCard(SEVEN, SPADE),
				},
				pair,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(TWO, HEART),
				newCard(THREE, DIAMOND),
				newCard(FIVE, SPADE),
				newCard(SEVEN, SPADE),
				newCard(JACK, SPADE),
				newCard(JACK, HEART),
			}, bestHand{
				[5]card{
					newCard(TWO, CLUB),
					newCard(TWO, HEART),
					newCard(JACK, HEART),
					newCard(JACK, SPADE),
					newCard(SEVEN, SPADE),
				},
				pair,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(TWO, HEART),
				newCard(THREE, DIAMOND),
				newCard(SEVEN, CLUB),
				newCard(SEVEN, SPADE),
				newCard(JACK, SPADE),
				newCard(JACK, HEART),
			}, bestHand{
				[5]card{
					newCard(THREE, DIAMOND),
					newCard(SEVEN, SPADE),
					newCard(SEVEN, CLUB),
					newCard(JACK, HEART),
					newCard(JACK, SPADE),
				},
				pair,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(EIGHT, HEART),
				newCard(THREE, DIAMOND),
				newCard(JACK, CLUB),
				newCard(SEVEN, SPADE),
				newCard(TWO, SPADE),
				newCard(TWO, HEART),
			}, bestHand{
				[5]card{
					newCard(EIGHT, HEART),
					newCard(JACK, CLUB),
					newCard(TWO, CLUB),
					newCard(TWO, HEART),
					newCard(TWO, SPADE),
				},
				threeOfAKind,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(EIGHT, HEART),
				newCard(THREE, DIAMOND),
				newCard(JACK, CLUB),
				newCard(THREE, SPADE),
				newCard(TWO, SPADE),
				newCard(TWO, HEART),
			}, bestHand{
				[5]card{
					newCard(THREE, DIAMOND),
					newCard(TWO, CLUB),
					newCard(THREE, SPADE),
					newCard(TWO, HEART),
					newCard(TWO, SPADE),
				},
				fullHouse,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(EIGHT, HEART),
				newCard(TWO, DIAMOND),
				newCard(JACK, CLUB),
				newCard(SEVEN, SPADE),
				newCard(TWO, SPADE),
				newCard(TWO, HEART),
			}, bestHand{
				[5]card{
					newCard(TWO, DIAMOND),
					newCard(JACK, CLUB),
					newCard(TWO, CLUB),
					newCard(TWO, HEART),
					newCard(TWO, SPADE),
				},
				threeOfAKind,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(THREE, HEART),
				newCard(FOUR, DIAMOND),
				newCard(JACK, CLUB),
				newCard(SEVEN, SPADE),
				newCard(FIVE, SPADE),
				newCard(SIX, HEART),
			}, bestHand{
				[5]card{
					newCard(SEVEN, SPADE),
					newCard(FIVE, SPADE),
					newCard(SIX, HEART),
					newCard(THREE, HEART),
					newCard(FOUR, DIAMOND),
				},
				straight,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(THREE, HEART),
				newCard(FOUR, DIAMOND),
				newCard(JACK, CLUB),
				newCard(SEVEN, SPADE),
				newCard(FIVE, SPADE),
				newCard(ACE, HEART),
			}, bestHand{
				[5]card{
					newCard(TWO, CLUB),
					newCard(FIVE, SPADE),
					newCard(ACE, HEART),
					newCard(THREE, HEART),
					newCard(FOUR, DIAMOND),
				},
				straight,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(THREE, CLUB),
				newCard(FOUR, DIAMOND),
				newCard(JACK, CLUB),
				newCard(SEVEN, CLUB),
				newCard(FIVE, SPADE),
				newCard(SIX, CLUB),
			}, bestHand{
				[5]card{
					newCard(TWO, CLUB),
					newCard(THREE, CLUB),
					newCard(JACK, CLUB),
					newCard(SEVEN, CLUB),
					newCard(SIX, CLUB),
				},
				flush,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(THREE, CLUB),
				newCard(FOUR, DIAMOND),
				newCard(JACK, CLUB),
				newCard(SEVEN, CLUB),
				newCard(FIVE, CLUB),
				newCard(SIX, CLUB),
			}, bestHand{
				[5]card{
					newCard(FIVE, CLUB),
					newCard(THREE, CLUB),
					newCard(JACK, CLUB),
					newCard(SEVEN, CLUB),
					newCard(SIX, CLUB),
				},
				flush,
			},
		},
		{
			[]card{
				newCard(EIGHT, SPADE),
				newCard(THREE, CLUB),
				newCard(SEVEN, CLUB),
				newCard(JACK, DIAMOND),
				newCard(FOUR, CLUB),
				newCard(FIVE, CLUB),
				newCard(SIX, CLUB),
			}, bestHand{
				[5]card{
					newCard(FIVE, CLUB),
					newCard(THREE, CLUB),
					newCard(FOUR, CLUB),
					newCard(SEVEN, CLUB),
					newCard(SIX, CLUB),
				},
				straightFlush,
			},
		},
	}
	for _, test := range tests {
		actual := buildBestHand(test.hand)
		for _, c := range actual.cards {
			if !containsCard(test.best.cards[:], c) {
				t.Error("there is at least one missing/incorrect card!\n",
					test.best.rank, test.best.cards, "\n",
					actual.rank, actual.cards)
				break
			}
		}
	}
}

