package main

import "sort"

const (
	leftWins verdict = iota
	rightWins
	tie
)

type verdict int

func sevenCardCompare(left, right []card) verdict {
	l := buildBestHand(left)
	r := buildBestHand(right)
	if l.rank > r.rank {
		return leftWins
	}
	if r.rank > l.rank {
		return rightWins
	}
	return compareBest(l.cards, r.cards, l.rank)
}

func compareBest(l, r [5]card, rank handRank) verdict {
	sort.Sort(h(l[:]))
	sort.Sort(h(r[:]))
	switch rank {
	case highcard, straight, straightFlush, flush:
		return highest(l[:], r[:])
	case pair:
		return comparePair(l[:], r[:])
	}
	return tie
}

// highest is given a sorted left and right hand and returns a verdict
// on which hand wins
func highest(l, r []card) verdict {
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

// comparePair is given a sorted left and right hand and returns
// a verdict on which hand wins
func comparePair(l, r []card) verdict {
	lpair, rpair := pairValue(l), pairValue(r)
	if lpair < rpair {
		return rightWins
	}
	if rpair < lpair {
		return leftWins
	}
	return highest(l, r)
}

// pairValue is an inefficient double for loop, which could be
// made more efficient because only sorted hands should be given.
// However, with hands of ~5 cards, I am not concerned with performance
// at the asymtote and this is easy to read.
func pairValue(hand []card) int {
	for i := range hand {
		for j := range hand {
			if i == j {
				continue
			}
			if hand[i].Value == hand[j].Value {
				return int(hand[i].Value)
			}
		}
	}
	return 0
}
