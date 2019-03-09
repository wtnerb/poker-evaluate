package main

import (
	"sort"
)

type handRank int
type h []card

const (
	highcard handRank = iota
	pair
	twoPair
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
)

func (r handRank) String() string {
	switch r {
	case pair:
		return "pair"
	case twoPair:
		return "two pair"
	case threeOfAKind:
		return "three of a kind"
	case fourOfAKind:
		return "three of a kind"
	case straight:
		return "straight"
	case flush:
		return "flush"
	case straightFlush:
		return "straight flush"
	case fullHouse:
		return "full house"
	case highcard:
		return "highcard"
	default:
		panic("invalid hand rank")
	}
}

func rankHand(hand []card) handRank {
	pairs := numPairs(hand)
	switch pairs {
	case 1:
		return pair
	}
	return highcard
}

// numPairs checks for the number of unique pairs it is possible to make
// with a given set of cards. When there are 7 or fewer cards to choose
// from (like texas hold'em) this enables a way to distinguish two pair,
// pair, three of a kind, full house and four of a kind for the maximum
// scoring set of 5 cards selected from the 7.
func numPairs(hand []card) int {
	sort.Sort(h(hand))
	pairs := 0
	for i := range hand {
		for j := i + 1; j < len(hand) && hand[j].Value == hand[i].Value; j++ {
			pairs++
		}
	}
	return pairs
}

//these three methods are required to fulfil the sort.Interface interface
func (hand h) Less(i, j int) bool {
	return hand[i].Value < hand[j].Value
}

func (hand h) Swap(i, j int) {
	hand[i], hand[j] = hand[j], hand[i]
}

func (hand h) Len() int {
	return len(hand)
}
