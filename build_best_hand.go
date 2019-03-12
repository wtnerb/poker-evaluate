package main

type best_hand struct {
	cards [5]card
	rank  handRank
}

func buildBestHand(cards []card) best_hand {
	r := rankHand(cards)
	var best [5]card
	switch r {
	case pair:
		best = buildPairHand(cards)
	}
	return best_hand{best, r}
}

func buildPairHand(cards []card) [5]card {
	var best [5]card
	b := best[:0]
	for i := range cards {
		if i == 0 {
			continue
		}
		if cards[i].Value == cards[i-1].Value {
			b = append(b, cards[i-1:i+1]...)
		}
	}
	fillOutHand(cards, &best, len(b))
	return best
}

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
// use pairs approach, which I currently am using

// build an "occurences [13]int"
