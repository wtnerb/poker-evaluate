package main

type best_hand struct {
	cards [5]card
	rank  handRank
}

func buildBestHand(cards []card) best_hand {
	r := rankHand(cards)
	var best [5]card
	numCards := 0
	switch r {
	case pair:
		numCards += buildPair(cards, &best)
	case twoPair:
		numCards += buildPair(cards, &best)
	case threeOfAKind:
		numCards += buildThreeOfAKind(cards, &best)
	case fourOfAKind:
		numCards += buildFourOfAKind(cards, &best)
	}
	fillOutHand(cards, &best, numCards)
	return best_hand{best, r}
}

func buildThreeOfAKind(cards []card, best *[5]card) int {
	b := best[:0]
	for i := len(cards) - 1; i > 1 && len(b) < 3; i-- {
		if cards[i].Value == cards[i-1].Value && cards[i].Value == cards[i-2].Value {
			b = append(b, cards[i-2:i+1]...)
		}
	}
	return len(b)
}

func buildFourOfAKind(cards []card, best *[5]card) int {
	b := best[:0]
	for i := len(cards) - 1; i > 2 && len(b) < 3; i-- {
		if cards[i].Value == cards[i-1].Value && cards[i].Value == cards[i-2].Value && cards[i].Value == cards[i-3].Value {
			b = append(b, cards[i-3:i+1]...)
		}
	}
	return len(b)
}

func buildPair(cards []card, best *[5]card) int {
	b := best[:0]
	for i := len(cards) - 1; i > 0 && len(b) < 4; i-- {
		if cards[i].Value == cards[i-1].Value {
			b = append(b, cards[i-1:i+1]...)
		}
	}
	return len(b)
}

// TODO: use map to keep in linear time, stay away from 7*5 time complexity
// Is this actually a significant slowdown? perhaps should profile
func fillOutHand(cards []card, best *[5]card, place int) {
	b := best[:place]
	for i := len(cards) - 1; i > 0 && len(b) < 5; i-- {
		if !containsCard(b, cards[i]) {
			b = append(b, cards[i])
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

// approaches for ranking hands:
// - use pairs approach, which I currently am using
// - build an "occurences [13][]card"
//
// Reasons for chosing each: pairs approach is already half done
//
// occurence approach uses easier logic.
