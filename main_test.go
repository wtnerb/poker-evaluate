package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	models "github.com/wtnerb/poker-models"
	"gopkg.in/mgo.v2/bson"
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
	r1, _ := json.Marshal(RespObj{
		[][]byte{
			[]byte(bson.ObjectIdHex("5c6482805508c93011b4e332")),
			[]byte(bson.ObjectIdHex("5c6482805508c93011b4e375")),
		},
	})

	tests := []struct {
		table    models.Table
		expected []byte
		status   int
	}{
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e375"),
						Name: "Brent", Cards: makeCards(cards{{TWO, HEART}, {TWO, CLUB}}), Folded: false,
					},
					// {
					// 	Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: false,
					// },
					models.TablePlayer{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e333"),
						Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
					},
				},
				FaceUp: []card{newCard(FIVE, SPADE), newCard(JACK, HEART), newCard(KING, CLUB), newCard(SEVEN, DIAMOND), newCard(NINE, DIAMOND)},
			},
			[]byte(bson.ObjectIdHex("5c6482805508c93011b4e375")),
			http.StatusOK,
		},
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e375"),
						Name: "Brent", Cards: makeCards(cards{{TWO, HEART}, {TWO, CLUB}}), Folded: false,
					},
					// {
					// 	Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: false,
					// },
					models.TablePlayer{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e333"),
						Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
					},
					models.TablePlayer{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e332"),
						Name: "Jason", Cards: []card{newCard(TWO, DIAMOND), newCard(TWO, SPADE)}, Folded: false,
					},
				},
				FaceUp: []card{newCard(FIVE, SPADE), newCard(JACK, HEART), newCard(KING, CLUB), newCard(SEVEN, DIAMOND), newCard(NINE, DIAMOND)},
			},
			r1,
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

		var resp RespObj
		err = json.Unmarshal(res.Body.Bytes(), &resp)
		// resposes will be in the respObj form. Getting something else
		// should mean an error message (likely in plain text) is the response
		if err != nil {
			if 0 != bytes.Compare(test.expected, res.Body.Bytes()) {
				t.Error("Error response test is not what was expected. Expected:", string(test.expected), "\nrecieved:", string(res.Body.Bytes()))
			}
			return
		}
		if bytes.Compare(test.expected, resp.Ids[0]) != 0 {
			t.Error("reponse was wrong. Expected\n", test.expected, "\ngot", resp)
		}
	}
}
