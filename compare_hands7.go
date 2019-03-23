package main

import (
	"sort"

	models "github.com/wtnerb/poker-models"
)

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
	case highcard, flush:
		return highest(l[:], r[:])
	case pair:
		return comparePair(l[:], r[:])
	case straight, straightFlush:
		return compareStraight(l[:], r[:])
	case threeOfAKind, fullHouse:
		return compareThreeOfAKind(l[:], r[:])
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

func compareStraight(l, r []card) verdict {
	if l[0].Value == models.ACE && l[len(l)-1].Value == 2 {
		flipAce(l)
		defer revertAce(l)
	}
	if r[0].Value == models.ACE && r[len(r)-1].Value == 2 {
		flipAce(r)
		defer revertAce(r)
	}
	return highest(l, r)
}

// flipAce MUST be followed by a call to revertAce
func flipAce(cards []card) {
	for i := range cards {
		if cards[i].Value == models.ACE {
			cards[i].Value -= 13
		}
	}
}

func revertAce(cards []card) {
	for i := range cards {
		if cards[i].Value == models.ACE-13 {
			cards[i].Value = models.ACE
		}
	}
}

func compareThreeOfAKind(l, r []card) verdict {
	l3 := threeValue(l)
	r3 := threeValue(r)
	if l3 < r3 {
		return rightWins
	}
	if r3 < l3 {
		return leftWins
	}
	return highest(l, r)
}

func threeValue(cards []card) int {

	for i := 0; i < len(cards)-2; i++ {
		if cards[i].Value == cards[i+1].Value && cards[i].Value == cards[i+2].Value {
			return int(cards[i].Value)
		}
	}
	return 0
}
