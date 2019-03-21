package main

const (
	leftWins verdict = iota
	rightWins
	tie
)

type verdict int

func sevenCardCompare(left, right []card) (verdict, bestHand) {
	l := buildBestHand(left)
	r := buildBestHand(right)
	if l.rank > r.rank {
		return leftWins, l
	}
	if r.rank > l.rank {
		return rightWins, r
	}
	return compareBest(l, r), l
}

func compareBest(l, r bestHand) verdict {

	return tie
}
