package main

import "sort"

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
	return compareBest(l.cards, r.cards, l.rank), l
}

func compareBest(l, r [5]card, rank handRank) verdict {
	sort.Sort(h(l[:]))
	sort.Sort(h(r[:]))
	switch rank {
	case highcard:
		return highest(l[:], r[:])
	}
	return tie
}

func highest(l, r []card) verdict {
	sort.Sort(h(l))
	sort.Sort(h(r))
	for i := range l {
		if l[i].Value > r[i].Value {
			return leftWins
		}
		if l[i].Value < r[i].Value {
			return rightWins
		}
	}
	return tie
}
