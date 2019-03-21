package main

import (
	"sort"

	models "github.com/wtnerb/poker-models"
)

type bestHand struct {
	cards [5]card
	rank  handRank
}

func buildBestHand(source []card) bestHand {
	r := rankHand(source)
	var best [5]card

	// rank hand already does this, but since the builds all
	// depend upon getting things sorted by value it was considered
	// wise to explicitly sort here
	sort.Sort(h(source))
	var numCards int
	switch r {
	case pair:
		numCards = buildPair(source, &best)
	case twoPair:
		numCards = buildPair(source, &best)
	case threeOfAKind:
		numCards = buildThreeOfAKind(source, &best)
	case fourOfAKind:
		numCards = buildFourOfAKind(source, &best)
	case straight:
		numCards = buildStraight(source, &best)
	case flush:
		numCards = buildFlush(source, &best)
	case straightFlush:
		numCards = buildStraightFlush(source, &best)
	case fullHouse:
		numCards = buildFullHouse(source, &best)
	}
	fillOutHand(source, &best, numCards)
	return bestHand{best, r}
}

func buildFullHouse(source []card, best *[5]card) int {
	_ = buildThreeOfAKind(source, best)
	r := best[:3]
	for i := len(source) - 1; i > 0; i-- {
		a, b := source[i], source[i-1]
		if a.Value == b.Value && !(containsCard(r, a) && !containsCard(r, b)) {
			_ = append(r, source[i-1:i+1]...)
			return 5
		}
	}
	return 0
}

// TODO: this is not robust to being incorrectly called.
func buildStraight(source []card, best *[5]card) int {
	noDups := pruneDuplicateValues(source)
	for i := range noDups {
		if len(noDups)-i < 5 {
			if noDups[0].Value == models.ACE && noDups[len(noDups)-4].Value == 5 {
				s := append(noDups[0:1], noDups[len(noDups)-4:]...)
				_ = append(best[:0], s...)
				return 5
			}
			panic("there isn't a straight")
		}
		if noDups[i].Value == noDups[i+4].Value+4 {
			_ = append(best[:0], source[i:i+5]...)
			return 5
		}
	}
	return 0
}

func buildStraightFlush(source []card, best *[5]card) int {
	suits := bySuit(source)

	for _, suit := range suits {
		if len(suit) >= 5 {
			return buildStraight(suit, best)
		}
	}
	panic("Building a straight flush when there is not a flush")
}

func buildFlush(source []card, best *[5]card) int {
	suits := bySuit(source)

	for _, suit := range suits {
		if len(suit) >= 5 {
			_ = append(best[:0], suit[0:5]...)
			return 5
		}
	}
	panic("Building a flush when there is not a flush")
}

func pruneDuplicateValues(source []card) (noPairs []card) {
	vals := make(map[int8]bool)
	for _, c := range source {
		if _, ok := vals[int8(c.Value)]; ok {
			continue
		}

		vals[int8(c.Value)] = true
		noPairs = append(noPairs, c)
	}
	return
}

func buildThreeOfAKind(source []card, best *[5]card) int {
	b := best[:0]
	for i := len(source) - 1; i > 1 && len(b) < 3; i-- {
		if source[i].Value == source[i-1].Value && source[i].Value == source[i-2].Value {
			b = append(b, source[i-2:i+1]...)
		}
	}
	return len(b)
}

func buildFourOfAKind(source []card, best *[5]card) int {
	b := best[:0]
	for i := len(source) - 1; i > 2 && len(b) < 3; i-- {
		if source[i].Value == source[i-1].Value && source[i].Value == source[i-2].Value && source[i].Value == source[i-3].Value {
			b = append(b, source[i-3:i+1]...)
		}
	}
	return len(b)
}

func buildPair(source []card, best *[5]card) int {
	b := best[:0]
	for i := range source {
		if len(b) >= 4 || i == len(source)-1 {
			break
		}
		if source[i].Value == source[i+1].Value {
			b = append(b, source[i:i+2]...)
		}
	}
	return len(b)
}

// TODO: use map to keep in linear time, stay away from 7*5 time complexity.
// RESP: Is this actually a significant slowdown? profile before refactor
func fillOutHand(source []card, best *[5]card, place int) {
	b := best[:place]
	for _, c := range source {
		if len(b) >= 5 {
			break
		}
		if !containsCard(b, c) {
			b = append(b, c)
		}
	}
}

func containsCard(cards []card, target card) bool {
	for _, c := range cards {
		if c.Value == target.Value && c.Suit == target.Suit {
			return true
		}
	}
	return false
}
