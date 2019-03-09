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
		expected int
	}{
		{
			[][2]card{
				[2]card{card{2, 3}, card{2, 2}},
				[2]card{card{3, 1}, card{4, 2}},
				[2]card{card{4, 1}, card{3, 2}},
			},
			[]card{card{5, 1}, card{11, 2}, card{13, 3}, card{7, 4}, card{9, 4}},
			0,
		},
	}

	for _, test := range tests {
		winner := evaluateHands(test.table, test.hands)
		if winner != test.expected {
			t.Errorf("Somebody won who was not supposed to win! \nexpected %d, got %d\ntable: %v\nhands: %v", test.expected, winner, test.table, test.hands)
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
	}

	for _, test := range tests {
		actual := rankHand(test.hand)
		if actual != test.rank {
			t.Errorf("Hand ranked incorrectly. %v\nranked as %v should be %v", test.hand, actual, test.rank)
		}
	}
}
