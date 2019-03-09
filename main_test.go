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
	//c := models.Card

	table1 := models.Table{
		Players: []models.TablePlayer{
			{
				Name: "Brent", Cards: []models.Card{models.Card{2, 2}, models.Card{2, 3}}, Folded: true,
			},
			{
				Name: "Devin", Cards: []models.Card{models.Card{3, 1}, models.Card{4, 2}}, Folded: false,
			},
		},
		FaceUp: []models.Card{models.Card{5, 1}, models.Card{11, 2}, models.Card{13, 3}, models.Card{7, 4}, models.Card{9, 4}},
	}

	winner := evaluateSingleWinner(table1)
	if winner.Name != "Devin" {
		t.Error("Everyone except Devin folded, how did he not win?")
	}
}
