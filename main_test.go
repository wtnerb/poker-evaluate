package main

import (
	"poker/models"
	"testing"
)

type card = models.Card

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
