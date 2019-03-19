package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"poker/models"
	"strings"
	"testing"
)

const (
	SPADE = 1 + iota
	HEART
	CLUB
	DIAMOND
)

const (
	ACE = models.ACE - iota
	KING
	QUEEN
	JACK
	TEN
	NINE
	EIGHT
	SEVEN
	SIX
	FIVE
	FOUR
	THREE
	TWO
)

func TestServer(t *testing.T) {
	expected := []byte(`["Brent"]`)

	JSON, err := json.Marshal(models.Table{
		Players: []models.TablePlayer{
			{
				Name: "Brent", Cards: []card{card{TWO, HEART}, card{TWO, CLUB}}, Folded: false,
			},
			models.TablePlayer{
				Name: "Devin", Cards: []card{card{THREE, SPADE}, card{FOUR, HEART}}, Folded: false,
			},
		},
		FaceUp: []card{card{FIVE, SPADE}, card{JACK, HEART}, card{KING, CLUB}, card{SEVEN, DIAMOND}, card{NINE, DIAMOND}},
	})

	if err != nil {
		t.Error("failed while marshelling the json")
	}

	req, err := http.NewRequest("POST", "http://localhost:4002/", strings.NewReader(string(JSON)))
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	recieveTable(res, req)

	if res.Code != http.StatusOK {
		t.Error("response code was not desired", res.Code)
	}

	if bytes.Compare(expected, res.Body.Bytes()) != 0 {
		t.Error("reponse was wrong. Expected\n", string(expected), "\ngot", string(res.Body.Bytes()))
	}
}

func TestWinner(t *testing.T) {
	// At the moment, this is a dummy test. Started working, realized
	// this was too big to be a single unit
	tests := []struct {
		table    models.Table
		expected models.TablePlayer
	}{
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						Name: "Brent", Cards: []card{card{TWO, HEART}, card{TWO, CLUB}}, Folded: true,
					},
					models.TablePlayer{
						Name: "Devin", Cards: []card{card{THREE, SPADE}, card{FOUR, HEART}}, Folded: false,
					},
				},
				FaceUp: []card{card{FIVE, SPADE}, card{JACK, HEART}, card{KING, CLUB}, card{SEVEN, DIAMOND}, card{NINE, DIAMOND}},
			},
			models.TablePlayer{
				Name: "Devin", Cards: []card{card{THREE, SPADE}, card{FOUR, HEART}}, Folded: false,
			},
		},
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						Name: "Brent", Cards: []card{card{TWO, HEART}, card{TWO, CLUB}}, Folded: false,
					},
					models.TablePlayer{
						Name: "Devin", Cards: []card{card{THREE, SPADE}, card{FOUR, HEART}}, Folded: false,
					},
				},
				FaceUp: []card{card{FIVE, SPADE}, card{JACK, HEART}, card{KING, CLUB}, card{SEVEN, DIAMOND}, card{NINE, DIAMOND}},
			},
			models.TablePlayer{
				Name: "Brent", Cards: []card{card{TWO, HEART}, card{TWO, CLUB}}, Folded: false,
			},
		},
	}
	for _, test := range tests {
		actual, err := evaluateWinner(test.table)
		if err != nil {
			t.Error(err)
		}
		if actual.Name != test.expected.Name {
			t.Error("wrong person won. Expected", test.expected.Name, "got", actual.Name, "\n", test.table)
		}
	}
}

func TestRankHand(t *testing.T) {
	tests := []struct {
		hand []card
		rank handRank
	}{
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{FIVE, SPADE},
				card{SEVEN, SPADE},
				card{JACK, SPADE},
				card{QUEEN, HEART},
			}, pair,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{FIVE, SPADE},
				card{QUEEN, SPADE},
				card{JACK, SPADE},
				card{TWO, DIAMOND},
			}, threeOfAKind,
		},
		{
			[]card{
				card{TWO, SPADE},
				card{TWO, HEART},
				card{THREE, SPADE},
				card{FIVE, SPADE},
				card{QUEEN, SPADE},
				card{JACK, SPADE},
				card{TWO, DIAMOND},
			}, flush,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{FIVE, SPADE},
				card{FIVE, CLUB},
				card{JACK, SPADE},
				card{JACK, HEART},
			}, twoPair,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{FOUR, HEART},
				card{THREE, DIAMOND},
				card{SIX, SPADE},
				card{FIVE, CLUB},
				card{JACK, SPADE},
				card{JACK, HEART},
			}, straight,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{FOUR, HEART},
				card{THREE, DIAMOND},
				card{SIX, SPADE},
				card{FIVE, CLUB},
				card{JACK, SPADE},
				card{FIVE, HEART},
			}, straight,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{FOUR, CLUB},
				card{THREE, CLUB},
				card{SIX, CLUB},
				card{FIVE, CLUB},
				card{JACK, SPADE},
				card{FIVE, DIAMOND},
			}, straightFlush,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{FOUR, CLUB},
				card{THREE, CLUB},
				card{SIX, SPADE},
				card{FIVE, CLUB},
				card{JACK, CLUB},
				card{FIVE, DIAMOND},
			}, flush,
		},
		{
			[]card{
				card{TWO, CLUB},
				card{FOUR, HEART},
				card{THREE, DIAMOND},
				card{models.ACE, SPADE}, //testing that Ace's will switch to a one when needed
				card{FIVE, CLUB},
				card{JACK, SPADE},
				card{FIVE, HEART},
			}, straight,
		},
	}

	for _, test := range tests {
		actual := rankHand(test.hand)
		if actual != test.rank {
			t.Errorf("Hand ranked incorrectly. %v\nranked as %v should be %v", test.hand, actual, test.rank)
		}
	}
}

func TestBestHand(t *testing.T) {
	tests := []struct {
		hand []card
		best bestHand
	}{
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{FIVE, SPADE},
				card{SEVEN, SPADE},
				card{JACK, SPADE},
				card{QUEEN, HEART},
			}, bestHand{
				[5]card{
					card{TWO, CLUB},
					card{TWO, HEART},
					card{QUEEN, HEART},
					card{JACK, SPADE},
					card{SEVEN, SPADE},
				},
				pair,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{FIVE, SPADE},
				card{SEVEN, SPADE},
				card{JACK, SPADE},
				card{JACK, HEART},
			}, bestHand{
				[5]card{
					card{TWO, CLUB},
					card{TWO, HEART},
					card{JACK, HEART},
					card{JACK, SPADE},
					card{SEVEN, SPADE},
				},
				pair,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{TWO, HEART},
				card{THREE, DIAMOND},
				card{SEVEN, CLUB},
				card{SEVEN, SPADE},
				card{JACK, SPADE},
				card{JACK, HEART},
			}, bestHand{
				[5]card{
					card{THREE, DIAMOND},
					card{SEVEN, SPADE},
					card{SEVEN, CLUB},
					card{JACK, HEART},
					card{JACK, SPADE},
				},
				pair,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{EIGHT, HEART},
				card{THREE, DIAMOND},
				card{JACK, CLUB},
				card{SEVEN, SPADE},
				card{TWO, SPADE},
				card{TWO, HEART},
			}, bestHand{
				[5]card{
					card{EIGHT, HEART},
					card{JACK, CLUB},
					card{TWO, CLUB},
					card{TWO, HEART},
					card{TWO, SPADE},
				},
				threeOfAKind,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{EIGHT, HEART},
				card{THREE, DIAMOND},
				card{JACK, CLUB},
				card{THREE, SPADE},
				card{TWO, SPADE},
				card{TWO, HEART},
			}, bestHand{
				[5]card{
					card{THREE, DIAMOND},
					card{TWO, CLUB},
					card{THREE, SPADE},
					card{TWO, HEART},
					card{TWO, SPADE},
				},
				fullHouse,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{EIGHT, HEART},
				card{TWO, DIAMOND},
				card{JACK, CLUB},
				card{SEVEN, SPADE},
				card{TWO, SPADE},
				card{TWO, HEART},
			}, bestHand{
				[5]card{
					card{TWO, DIAMOND},
					card{JACK, CLUB},
					card{TWO, CLUB},
					card{TWO, HEART},
					card{TWO, SPADE},
				},
				threeOfAKind,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{THREE, HEART},
				card{FOUR, DIAMOND},
				card{JACK, CLUB},
				card{SEVEN, SPADE},
				card{FIVE, SPADE},
				card{SIX, HEART},
			}, bestHand{
				[5]card{
					card{SEVEN, SPADE},
					card{FIVE, SPADE},
					card{SIX, HEART},
					card{THREE, HEART},
					card{FOUR, DIAMOND},
				},
				straight,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{THREE, CLUB},
				card{FOUR, DIAMOND},
				card{JACK, CLUB},
				card{SEVEN, CLUB},
				card{FIVE, SPADE},
				card{SIX, CLUB},
			}, bestHand{
				[5]card{
					card{TWO, CLUB},
					card{THREE, CLUB},
					card{JACK, CLUB},
					card{SEVEN, CLUB},
					card{SIX, CLUB},
				},
				flush,
			},
		},
		{
			[]card{
				card{TWO, CLUB},
				card{THREE, CLUB},
				card{FOUR, DIAMOND},
				card{JACK, CLUB},
				card{SEVEN, CLUB},
				card{FIVE, CLUB},
				card{SIX, CLUB},
			}, bestHand{
				[5]card{
					card{FIVE, CLUB},
					card{THREE, CLUB},
					card{JACK, CLUB},
					card{SEVEN, CLUB},
					card{SIX, CLUB},
				},
				flush,
			},
		},
		{
			[]card{
				card{EIGHT, SPADE},
				card{THREE, CLUB},
				card{SEVEN, CLUB},
				card{JACK, DIAMOND},
				card{FOUR, CLUB},
				card{FIVE, CLUB},
				card{SIX, CLUB},
			}, bestHand{
				[5]card{
					card{FIVE, CLUB},
					card{THREE, CLUB},
					card{FOUR, CLUB},
					card{SEVEN, CLUB},
					card{SIX, CLUB},
				},
				straightFlush,
			},
		},
	}
	for _, test := range tests {
		actual := buildBestHand(test.hand)
		for _, c := range actual.cards {
			if !containsCard(test.best.cards[:], c) {
				t.Error("there is at least one missing/incorrect card!\n",
					test.best.rank, test.best.cards, "\n",
					actual.rank, actual.cards)
				break
			}
		}
	}
}
