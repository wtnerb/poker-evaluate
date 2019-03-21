package main

import (
	"sort"

	models "github.com/wtnerb/poker-models"
)

type handRank int
type h []card

const ( // This list and the ranks in the String method _MUST_ match
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
	ranks := []string{ // This list _MUST_ match the constants above
		"highcard",
		"pair",
		"two pair",
		"three of a kind",
		"straight",
		"flush",
		"full house",
		"four of a kind",
		"straight flush",
	}
	return ranks[r]
}

func rankHand(hand []card) handRank {
	pairs := numPairs(hand)
	fl := isFlush(hand)
	if fl && isStraightFlush(hand) {
		return straightFlush
	}
	switch pairs {
	case 6:
		return fourOfAKind
	case 4:
		return fullHouse
	}
	if fl {
		return flush
	}
	if isStraight(hand) {
		return straight
	}
	switch pairs {
	case 3:
		// with seven cards, it is possible to have three pairs OR a three of a kind
		for i := 0; i < len(hand)-2; i++ {
			if hand[i].Value == hand[i+2].Value {
				return threeOfAKind
			}
		}
		return twoPair
	case 2:
		return twoPair
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
	return hand[i].Value > hand[j].Value
}

func (hand h) Swap(i, j int) {
	hand[i], hand[j] = hand[j], hand[i]
}

func (hand h) Len() int {
	return len(hand)
}

// TODO: Too much logic here. Should be broken up.
func isStraight(hand []card) bool {
	// checking for a straight is much easier without worrying about
	// duplicate values in the middle of the sorted straight.
	noDups := pruneDuplicateValues(hand)
	// if there are not at least 5 distinct values, a straight is
	// impossible
	if len(noDups) < 5 {
		return false
	}

	// double for loop is inefficient, but hand size should be <= 7, so
	// efficiency at the asymtote is not relevant.
	for i := 0; i < len(noDups)-4; i++ {
		maybe := true
		for j := 1; j < 5; j++ {
			if int(noDups[i+j].Value) != int(noDups[i].Value)-j {
				maybe = false
				break
			}
		}
		if maybe {
			return true
		}
	}

	//special logic for ace switching to be a 'one' value
	if noDups[0].Value == models.ACE && noDups[len(noDups)-4].Value == 5 {
		return true
	}
	return false
}

func isFlush(hand []card) bool {
	suits := bySuit(hand)
	// check counts of suit
	for _, suit := range suits {
		if len(suit) >= 5 {
			return true
		}
	}

	// if no flushes have been found, return false
	return false
}

func isStraightFlush(hand []card) bool {
	suits := bySuit(hand)

	for _, suit := range suits {
		if len(suit) >= 5 && isStraight(suit) {
			return true
		}
	}
	return false
}

func bySuit(source []card) (suits [5][]card) {
	for _, c := range source {
		suits[c.Suit] = append(suits[c.Suit], c)
	}
	if len(suits[0]) != 0 {
		panic("card with invalid suit has been found.")
	}
	return
}
