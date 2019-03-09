package main

import (
	"fmt"
	"poker/models"
)

func main() {
	fmt.Println("hello world!")
}

func evaluateSingleWinner(table models.Table) models.TablePlayer {
	// if table.Players[0] == nil {
	// 	panic("there are no players at the table!")
	// }
	return table.Players[0]
}
