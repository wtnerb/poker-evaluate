package main

import (
	"poker/models"
	"testing"
)

const (
	SPADE = 1 + iota
	HEART
	CLUB
	DIAMOND
)

const (
	ACE = models.ACE - iota
	KING
	QUEEN
	JACK
	TEN
	NINE
	EIGHT
	SEVEN
	SIX
	FIVE
	FOUR
	THREE
	TWO
)

func TestSingleWinner(t *testing.T) {
	// At the moment, this is a dummy test. Started working, realized
	// this was too big to be a single unit
	// tests := []struct{
	// 	table models.Table,
	// 	expected models.TablePlayer
	// }{
	// 	{
	// 		{
	// 			{
	// 				Name
	// 			}
	// 		}
	// 	}
	// }

	table1 := models.Table{
		Players: []models.TablePlayer{
			{
				Name: "Brent", Cards: []card{card{TWO, HEART}, card{TWO, CLUB}}, Folded: true,
			},
			{
				Name: "Devin", Cards: []card{card{THREE, SPADE}, card{FOUR, HEART}}, Folded: false,
			},
		},
		FaceUp: []card{card{FIVE, SPADE}, card{JACK, HEART}, card{KING, CLUB}, card{SEVEN, DIAMOND}, card{NINE, DIAMOND}},
	}

	winner, _ := evaluateSingleWinner(table1)
	if winner.Name != "Devin" {
		t.Error("Everyone except Devin folded, how did he not win?")
	}
}

func TestEvaluateHands(t *testing.T) {
	tests := []struct {
		hands    [][2]card
		table    []card
		expected []handRank
	}{
		{
			[][2]card{
				[2]card{card{TWO, CLUB}, card{TWO, HEART}},
				[2]card{card{THREE, SPADE}, card{FOUR, HEART}},
				[2]card{card{FOUR, SPADE}, card{models.ACE, HEART}},
			},
			[]card{
				card{FIVE, SPADE},
				card{JACK, HEART},
				card{SIX, CLUB},
				card{SEVEN, DIAMOND},
				card{NINE, DIAMOND}},
			[]handRank{pair, straight, highcard},
		},
	}

	for _, test := range tests {
		ranks := rankHands(test.table, test.hands)
		for i := range test.expected {
			if ranks[i] != test.expected[i] {
				t.Errorf("A hand was not ranked correctly! hand number %d, expected %v, got %v\ntable: %v\nhands: %v", i, test.expected[i], ranks[i], test.table, test.hands)
			}

		}
	}
}

func TestRankHand(t *testing.T) {
	tests := []struct {
		hand []card
		rank handRank
	}{
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{FIVE, SPADE},
				card{SEVEN, SPADE},
				card{JACK, SPADE},
				card{QUEEN, HEART},
			}, pair,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{FIVE, SPADE},
				card{QUEEN, SPADE},
				card{JACK, SPADE},
				card{TWO, DIAMOND},
			}, threeOfAKind,
		},
		{
			[]card{
				card{TWO, SPADE},
				card{TWO, HEART},
				card{THREE, SPADE},
				card{FIVE, SPADE},
				card{QUEEN, SPADE},
				card{JACK, SPADE},
				card{TWO, DIAMOND},
			}, flush,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{FIVE, SPADE},
				card{FIVE, CLUB},
				card{JACK, SPADE},
				card{JACK, HEART},
			}, twoPair,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{FOUR, HEART},
				card{THREE, DIAMOND},
				card{SIX, SPADE},
				card{FIVE, CLUB},
				card{JACK, SPADE},
				card{JACK, HEART},
			}, straight,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{FOUR, HEART},
				card{THREE, DIAMOND},
				card{SIX, SPADE},
				card{FIVE, CLUB},
				card{JACK, SPADE},
				card{FIVE, HEART},
			}, straight,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{FOUR, CLUB},
				card{THREE, CLUB},
				card{SIX, CLUB},
				card{FIVE, CLUB},
				card{JACK, SPADE},
				card{FIVE, DIAMOND},
			}, straightFlush,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{FOUR, CLUB},
				card{THREE, CLUB},
				card{SIX, SPADE},
				card{FIVE, CLUB},
				card{JACK, CLUB},
				card{FIVE, DIAMOND},
			}, flush,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{FOUR, HEART},
				card{THREE, DIAMOND},
				card{models.ACE, SPADE}, //testing that Ace's will switch to a one when needed
				card{FIVE, CLUB},
				card{JACK, SPADE},
				card{FIVE, HEART},
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

func TestBestHand(t *testing.T) {
	tests := []struct {
		hand []card
		best bestHand
	}{
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{FIVE, SPADE},
				card{SEVEN, SPADE},
				card{JACK, SPADE},
				card{QUEEN, HEART},
			}, bestHand{
				[5]card{
					card{TWO, CLUB},
					card{TWO, HEART},
					card{QUEEN, HEART},
					card{JACK, SPADE},
					card{SEVEN, SPADE},
				},
				pair,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{FIVE, SPADE},
				card{SEVEN, SPADE},
				card{JACK, SPADE},
				card{JACK, HEART},
			}, bestHand{
				[5]card{
					card{TWO, CLUB},
					card{TWO, HEART},
					card{JACK, HEART},
					card{JACK, SPADE},
					card{SEVEN, SPADE},
				},
				pair,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{SEVEN, CLUB},
				card{SEVEN, SPADE},
				card{JACK, SPADE},
				card{JACK, HEART},
			}, bestHand{
				[5]card{
					card{THREE, DIAMOND},
					card{SEVEN, SPADE},
					card{SEVEN, CLUB},
					card{JACK, HEART},
					card{JACK, SPADE},
				},
				pair,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{EIGHT, HEART},
				card{THREE, DIAMOND},
				card{JACK, CLUB},
				card{SEVEN, SPADE},
				card{TWO, SPADE},
				card{TWO, HEART},
			}, bestHand{
				[5]card{
					card{EIGHT, HEART},
					card{JACK, CLUB},
					card{TWO, CLUB},
					card{TWO, HEART},
					card{TWO, SPADE},
				},
				threeOfAKind,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{EIGHT, HEART},
				card{THREE, DIAMOND},
				card{JACK, CLUB},
				card{THREE, SPADE},
				card{TWO, SPADE},
				card{TWO, HEART},
			}, bestHand{
				[5]card{
					card{THREE, DIAMOND},
					card{JACK, CLUB},
					card{THREE, SPADE},
					card{TWO, HEART},
					card{TWO, SPADE},
				},
				fullHouse,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{EIGHT, HEART},
				card{TWO, DIAMOND},
				card{JACK, CLUB},
				card{SEVEN, SPADE},
				card{TWO, SPADE},
				card{TWO, HEART},
			}, bestHand{
				[5]card{
					card{TWO, DIAMOND},
					card{JACK, CLUB},
					card{TWO, CLUB},
					card{TWO, HEART},
					card{TWO, SPADE},
				},
				threeOfAKind,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{THREE, HEART},
				card{FOUR, DIAMOND},
				card{JACK, CLUB},
				card{SEVEN, SPADE},
				card{FIVE, SPADE},
				card{SIX, HEART},
			}, bestHand{
				[5]card{
					card{SEVEN, SPADE},
					card{FIVE, SPADE},
					card{SIX, HEART},
					card{THREE, HEART},
					card{FOUR, DIAMOND},
				},
				straight,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{THREE, CLUB},
				card{FOUR, DIAMOND},
				card{JACK, CLUB},
				card{SEVEN, CLUB},
				card{FIVE, SPADE},
				card{SIX, CLUB},
			}, bestHand{
				[5]card{
					card{TWO, CLUB},
					card{THREE, CLUB},
					card{JACK, CLUB},
					card{SEVEN, CLUB},
					card{SIX, CLUB},
				},
				flush,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{THREE, CLUB},
				card{FOUR, DIAMOND},
				card{JACK, CLUB},
				card{SEVEN, CLUB},
				card{FIVE, CLUB},
				card{SIX, CLUB},
			}, bestHand{
				[5]card{
					card{FIVE, CLUB},
					card{THREE, CLUB},
					card{JACK, CLUB},
					card{SEVEN, CLUB},
					card{SIX, CLUB},
				},
				flush,
			},
		},
		{
			[]card{
				card{EIGHT, SPADE},
				card{THREE, CLUB},
				card{SEVEN, CLUB},
				card{JACK, DIAMOND},
				card{FOUR, CLUB},
				card{FIVE, CLUB},
				card{SIX, CLUB},
			}, bestHand{
				[5]card{
					card{FIVE, CLUB},
					card{THREE, CLUB},
					card{FOUR, CLUB},
					card{SEVEN, CLUB},
					card{SIX, CLUB},
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
