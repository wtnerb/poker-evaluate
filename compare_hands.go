package main

func holdemCompare(tableCards []card, holes [][]card) (int, bestHand) {
	return 0, bestHand{}
}

const (
	LEFT = iota
	RIGHT
	TIE
)

func sevenCardCompare(left, right []card) int {
	l := buildBestHand(left)
	r := buildBestHand(right)
	if l.rank > r.rank {
		return LEFT
	}
	if r.rank > l.rank {
		return RIGHT
	}
	return TIE
}
