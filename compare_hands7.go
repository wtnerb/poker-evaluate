package main

func holdemCompare(tableCards []card, holes [][]card) (int, bestHand) {
	return 0, bestHand{}
}

const (
	leftWins = iota
	rightWins
	tie
)

func sevenCardCompare(left, right []card) (int, bestHand) {
	l := buildBestHand(left)
	r := buildBestHand(right)
	if l.rank > r.rank {
		return leftWins, l
	}
	if r.rank > l.rank {
		return rightWins, r
	}
	return tie, l
}
