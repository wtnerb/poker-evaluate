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

// TODO: This can probably be combined with buildBestHand to be more
// efficient. This is easier to reason as a human. Is there enough
// performance overhead to make a refactor worthwhile?
func rankHand(hand []card) handRank {
	// Working with a list sorted by value (high to low) is easier
	sort.Sort(h(hand))
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

func isStraight(hand []card) bool {
	// checking for a straight is much easier without worrying about
	// duplicate values in the middle of the sorted straight.
	noDups := pruneDuplicateValues(hand)

	// if the slice without duplicate values every has a value 4 slots
	// ahead with a value +4 than the current, there must be 3 intermediate
	// values x+1, x+2, and x+3. Combined with x and x+4, that makes a
	// five card straight
	for i := 0; i < len(noDups)-4; i++ {
		if noDups[i].Value == noDups[i+4].Value+4 {
			return true
		}
	}

	//special logic for ace switching to be a 'one' value
	if noDups[0].Value == models.ACE && noDups[len(noDups)-4].Value == 5 {
		return true
	}
	return false
}

// isFlush checks if a slice of cards has five or more cards of a single suit.
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

// isStraightFlush checks if a slice of cards contains a straight flush
func isStraightFlush(hand []card) bool {
	suits := bySuit(hand)

	for _, suit := range suits {
		if len(suit) >= 5 && isStraight(suit) {
			return true
		}
	}
	return false
}

// bySuit will take a source slice of cards and return four slices of
// cards organized by suit. No promise about order of suits.
func bySuit(source []card) (suits [4][]card) {
	for _, c := range source {
		suits[c.Suit-1] = append(suits[c.Suit-1], c)
	}
	return
}
