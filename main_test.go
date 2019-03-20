package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	models "github.com/wtnerb/poker-models"
)

const (
	SPADE = int8(1 + iota)
	HEART
	CLUB
	DIAMOND
)

const (
	ACE = int8(models.ACE - iota)
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

func newCard(v, s int8) card {
	return card(models.NewCard(v, s))
}

func makeCards(vals [][2]int8) (cs []card) {
	for _, c := range vals {
		cs = append(cs, newCard(c[0], c[1]))
	}
	return
}

type cards [][2]int8

func TestServer(t *testing.T) {
	tests := []struct {
		table    models.Table
		expected []byte
		status   int
	}{
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						Name: "Brent", Cards: makeCards(cards{{TWO, HEART}, {TWO, CLUB}}), Folded: false,
					},
					// {
					// 	Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: false,
					// },
					models.TablePlayer{
						Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
					},
				},
				FaceUp: []card{newCard(FIVE, SPADE), newCard(JACK, HEART), newCard(KING, CLUB), newCard(SEVEN, DIAMOND), newCard(NINE, DIAMOND)},
			},
			[]byte(`["Brent"]`),
			http.StatusOK,
		},
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: true,
					},
					models.TablePlayer{
						Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: true,
					},
				},
				FaceUp: []card{newCard(FIVE, SPADE), newCard(JACK, HEART), newCard(KING, CLUB), newCard(SEVEN, DIAMOND), newCard(NINE, DIAMOND)},
			},
			[]byte(`invalid game state there are no active players`),
			http.StatusBadRequest,
		},
	}

	for _, test := range tests {

		JSON, err := json.Marshal(test.table)

		if err != nil {
			t.Error("failed while marshelling the json")
		}

		req, err := http.NewRequest("POST", "http://localhost"+port+"/", strings.NewReader(string(JSON)))
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()

		recieveTable(res, req)

		if res.Code != test.status {
			t.Log(req)
			t.Log(res)
			t.Error("response code was not desired. Expected", test.status, "recieved", res.Code)
		}

		if bytes.Compare(test.expected, res.Body.Bytes()) != 0 {
			t.Error("reponse was wrong. Expected\n", string(test.expected), "\ngot", string(res.Body.Bytes()))
		}
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
						Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: true,
					},
					models.TablePlayer{
						Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
					},
				},
				FaceUp: []card{newCard(FIVE, SPADE), newCard(JACK, HEART), newCard(KING, CLUB), newCard(SEVEN, DIAMOND), newCard(NINE, DIAMOND)},
			},
			models.TablePlayer{
				Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
			},
		},
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: false,
					},
					models.TablePlayer{
						Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
					},
				},
				FaceUp: []card{newCard(FIVE, SPADE), newCard(JACK, HEART), newCard(KING, CLUB), newCard(SEVEN, DIAMOND), newCard(NINE, DIAMOND)},
			},
			models.TablePlayer{
				Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: false,
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
				newCard(TWO, CLUB),
				newCard(TWO, HEART),
				newCard(THREE, DIAMOND),
				newCard(FIVE, SPADE),
				newCard(SEVEN, SPADE),
				newCard(JACK, SPADE),
				newCard(QUEEN, HEART),
			}, pair,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(TWO, HEART),
				newCard(THREE, DIAMOND),
				newCard(FIVE, SPADE),
				newCard(QUEEN, SPADE),
				newCard(JACK, SPADE),
				newCard(TWO, DIAMOND),
			}, threeOfAKind,
		},
		{
			[]card{
				newCard(TWO, SPADE),
				newCard(TWO, HEART),
				newCard(THREE, SPADE),
				newCard(FIVE, SPADE),
				newCard(QUEEN, SPADE),
				newCard(JACK, SPADE),
				newCard(TWO, DIAMOND),
			}, flush,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(TWO, HEART),
				newCard(THREE, DIAMOND),
				newCard(FIVE, SPADE),
				newCard(FIVE, CLUB),
				newCard(JACK, SPADE),
				newCard(JACK, HEART),
			}, twoPair,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(FOUR, HEART),
				newCard(THREE, DIAMOND),
				newCard(SIX, SPADE),
				newCard(FIVE, CLUB),
				newCard(JACK, SPADE),
				newCard(JACK, HEART),
			}, straight,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(FOUR, HEART),
				newCard(THREE, DIAMOND),
				newCard(SIX, SPADE),
				newCard(FIVE, CLUB),
				newCard(JACK, SPADE),
				newCard(FIVE, HEART),
			}, straight,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(FOUR, CLUB),
				newCard(THREE, CLUB),
				newCard(SIX, CLUB),
				newCard(FIVE, CLUB),
				newCard(JACK, SPADE),
				newCard(FIVE, DIAMOND),
			}, straightFlush,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(FOUR, CLUB),
				newCard(THREE, CLUB),
				newCard(SIX, SPADE),
				newCard(FIVE, CLUB),
				newCard(JACK, CLUB),
				newCard(FIVE, DIAMOND),
			}, flush,
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(FOUR, HEART),
				newCard(THREE, DIAMOND),
				newCard(ACE, SPADE), //testing that Ace's will switch to a one when needed
				newCard(FIVE, CLUB),
				newCard(JACK, SPADE),
				newCard(FIVE, HEART),
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
				newCard(TWO, CLUB),
				newCard(TWO, HEART),
				newCard(THREE, DIAMOND),
				newCard(FIVE, SPADE),
				newCard(SEVEN, SPADE),
				newCard(JACK, SPADE),
				newCard(QUEEN, HEART),
			}, bestHand{
				[5]card{
					newCard(TWO, CLUB),
					newCard(TWO, HEART),
					newCard(QUEEN, HEART),
					newCard(JACK, SPADE),
					newCard(SEVEN, SPADE),
				},
				pair,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(TWO, HEART),
				newCard(THREE, DIAMOND),
				newCard(FIVE, SPADE),
				newCard(SEVEN, SPADE),
				newCard(JACK, SPADE),
				newCard(JACK, HEART),
			}, bestHand{
				[5]card{
					newCard(TWO, CLUB),
					newCard(TWO, HEART),
					newCard(JACK, HEART),
					newCard(JACK, SPADE),
					newCard(SEVEN, SPADE),
				},
				pair,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(TWO, HEART),
				newCard(THREE, DIAMOND),
				newCard(SEVEN, CLUB),
				newCard(SEVEN, SPADE),
				newCard(JACK, SPADE),
				newCard(JACK, HEART),
			}, bestHand{
				[5]card{
					newCard(THREE, DIAMOND),
					newCard(SEVEN, SPADE),
					newCard(SEVEN, CLUB),
					newCard(JACK, HEART),
					newCard(JACK, SPADE),
				},
				pair,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(EIGHT, HEART),
				newCard(THREE, DIAMOND),
				newCard(JACK, CLUB),
				newCard(SEVEN, SPADE),
				newCard(TWO, SPADE),
				newCard(TWO, HEART),
			}, bestHand{
				[5]card{
					newCard(EIGHT, HEART),
					newCard(JACK, CLUB),
					newCard(TWO, CLUB),
					newCard(TWO, HEART),
					newCard(TWO, SPADE),
				},
				threeOfAKind,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(EIGHT, HEART),
				newCard(THREE, DIAMOND),
				newCard(JACK, CLUB),
				newCard(THREE, SPADE),
				newCard(TWO, SPADE),
				newCard(TWO, HEART),
			}, bestHand{
				[5]card{
					newCard(THREE, DIAMOND),
					newCard(TWO, CLUB),
					newCard(THREE, SPADE),
					newCard(TWO, HEART),
					newCard(TWO, SPADE),
				},
				fullHouse,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(EIGHT, HEART),
				newCard(TWO, DIAMOND),
				newCard(JACK, CLUB),
				newCard(SEVEN, SPADE),
				newCard(TWO, SPADE),
				newCard(TWO, HEART),
			}, bestHand{
				[5]card{
					newCard(TWO, DIAMOND),
					newCard(JACK, CLUB),
					newCard(TWO, CLUB),
					newCard(TWO, HEART),
					newCard(TWO, SPADE),
				},
				threeOfAKind,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(THREE, HEART),
				newCard(FOUR, DIAMOND),
				newCard(JACK, CLUB),
				newCard(SEVEN, SPADE),
				newCard(FIVE, SPADE),
				newCard(SIX, HEART),
			}, bestHand{
				[5]card{
					newCard(SEVEN, SPADE),
					newCard(FIVE, SPADE),
					newCard(SIX, HEART),
					newCard(THREE, HEART),
					newCard(FOUR, DIAMOND),
				},
				straight,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(THREE, CLUB),
				newCard(FOUR, DIAMOND),
				newCard(JACK, CLUB),
				newCard(SEVEN, CLUB),
				newCard(FIVE, SPADE),
				newCard(SIX, CLUB),
			}, bestHand{
				[5]card{
					newCard(TWO, CLUB),
					newCard(THREE, CLUB),
					newCard(JACK, CLUB),
					newCard(SEVEN, CLUB),
					newCard(SIX, CLUB),
				},
				flush,
			},
		},
		{
			[]card{
				newCard(TWO, CLUB),
				newCard(THREE, CLUB),
				newCard(FOUR, DIAMOND),
				newCard(JACK, CLUB),
				newCard(SEVEN, CLUB),
				newCard(FIVE, CLUB),
				newCard(SIX, CLUB),
			}, bestHand{
				[5]card{
					newCard(FIVE, CLUB),
					newCard(THREE, CLUB),
					newCard(JACK, CLUB),
					newCard(SEVEN, CLUB),
					newCard(SIX, CLUB),
				},
				flush,
			},
		},
		{
			[]card{
				newCard(EIGHT, SPADE),
				newCard(THREE, CLUB),
				newCard(SEVEN, CLUB),
				newCard(JACK, DIAMOND),
				newCard(FOUR, CLUB),
				newCard(FIVE, CLUB),
				newCard(SIX, CLUB),
			}, bestHand{
				[5]card{
					newCard(FIVE, CLUB),
					newCard(THREE, CLUB),
					newCard(FOUR, CLUB),
					newCard(SEVEN, CLUB),
					newCard(SIX, CLUB),
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
