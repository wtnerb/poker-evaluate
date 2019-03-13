package main

import (
	"poker/models"
	"testing"
)

func TestSingleWinner(t *testing.T) {
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
				Name: "Brent", Cards: []card{card{2, 2}, card{2, 3}}, Folded: true,
			},
			{
				Name: "Devin", Cards: []card{card{3, 1}, card{4, 2}}, Folded: false,
			},
		},
		FaceUp: []card{card{5, 1}, card{11, 2}, card{13, 3}, card{7, 4}, card{9, 4}},
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
				[2]card{card{2, 3}, card{2, 2}},
				[2]card{card{3, 1}, card{4, 2}},
				[2]card{card{4, 1}, card{models.ACE, 2}},
			},
			[]card{
				card{5, 1},
				card{11, 2},
				card{6, 3},
				card{7, 4},
				card{9, 4}},
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
				card{2, 3},
				card{2, 2},
				card{3, 4},
				card{5, 1},
				card{7, 1},
				card{11, 1},
				card{12, 2},
			}, pair,
		},
		{
			[]card{
				card{2, 3},
				card{2, 2},
				card{3, 4},
				card{5, 1},
				card{12, 1},
				card{11, 1},
				card{2, 4},
			}, threeOfAKind,
		},
		{
			[]card{
				card{2, 1},
				card{2, 2},
				card{3, 1},
				card{5, 1},
				card{12, 1},
				card{11, 1},
				card{2, 4},
			}, flush,
		},
		{
			[]card{
				card{2, 3},
				card{2, 2},
				card{3, 4},
				card{5, 1},
				card{5, 3},
				card{11, 1},
				card{11, 2},
			}, twoPair,
		},
		{
			[]card{
				card{2, 3},
				card{4, 2},
				card{3, 4},
				card{6, 1},
				card{5, 3},
				card{11, 1},
				card{11, 2},
			}, straight,
		},
		{
			[]card{
				card{2, 3},
				card{4, 2},
				card{3, 4},
				card{6, 1},
				card{5, 3},
				card{11, 1},
				card{5, 2},
			}, straight,
		},
		{
			[]card{
				card{2, 3},
				card{4, 3},
				card{3, 3},
				card{6, 3},
				card{5, 3},
				card{11, 1},
				card{5, 4},
			}, straightFlush,
		},
		{
			[]card{
				card{2, 3},
				card{4, 3},
				card{3, 3},
				card{6, 1},
				card{5, 3},
				card{11, 3},
				card{5, 4},
			}, flush,
		},
		{
			[]card{
				card{2, 3},
				card{4, 2},
				card{3, 4},
				card{models.ACE, 1}, //testing that Ace's will switch to a one when needed
				card{5, 3},
				card{11, 1},
				card{5, 2},
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
		best best_hand
	}{
		{
			[]card{
				card{2, 3},
				card{2, 2},
				card{3, 4},
				card{5, 1},
				card{7, 1},
				card{11, 1},
				card{12, 2},
			}, best_hand{
				[5]card{
					card{2, 3},
					card{2, 2},
					card{12, 2},
					card{11, 1},
					card{7, 1},
				},
				pair,
			},
		},
		{
			[]card{
				card{2, 3},
				card{2, 2},
				card{3, 4},
				card{5, 1},
				card{7, 1},
				card{11, 1},
				card{11, 2},
			}, best_hand{
				[5]card{
					card{2, 3},
					card{2, 2},
					card{11, 2},
					card{11, 1},
					card{7, 1},
				},
				pair,
			},
		},
		{
			[]card{
				card{2, 3},
				card{2, 2},
				card{3, 4},
				card{7, 3},
				card{7, 1},
				card{11, 1},
				card{11, 2},
			}, best_hand{
				[5]card{
					card{3, 4},
					card{7, 1},
					card{7, 3},
					card{11, 2},
					card{11, 1},
				},
				pair,
			},
		},
		{
			[]card{
				card{2, 3},
				card{8, 2},
				card{3, 4},
				card{11, 3},
				card{7, 1},
				card{2, 1},
				card{2, 2},
			}, best_hand{
				[5]card{
					card{8, 2},
					card{11, 3},
					card{2, 3},
					card{2, 2},
					card{2, 1},
				},
				threeOfAKind,
			},
		},
		{
			[]card{
				card{2, 3},
				card{8, 2},
				card{2, 4},
				card{11, 3},
				card{7, 1},
				card{2, 1},
				card{2, 2},
			}, best_hand{
				[5]card{
					card{2, 4},
					card{11, 3},
					card{2, 3},
					card{2, 2},
					card{2, 1},
				},
				threeOfAKind,
			},
		},
	}
	for _, test := range tests {
		actual := buildBestHand(test.hand)
		for _, c := range actual.cards {
			if !containsCard(test.best.cards[:], c) {
				t.Error("there is a missing card!", test.best.cards, actual.cards)
			}
		}
	}
}
