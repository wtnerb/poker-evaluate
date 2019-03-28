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

func makeTestRequest(t *testing.T, table models.Table) *httptest.ResponseRecorder {
	JSON, err := json.Marshal(table)

	if err != nil {
		t.Log("failed while marshelling the json")
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://localhost"+port+"/", strings.NewReader(string(JSON)))
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	recieveTable(res, req)
	return res
}
func TestServerErrors(t *testing.T) {
	tests := []struct {
		table    models.Table
		expected []byte
		status   int
	}{
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
		res := makeTestRequest(t, test.table)

		if res.Code != test.status {
			t.Log(res)
			t.Error("response code was not desired. Expected", test.status, "recieved", res.Code)
		}

		if 0 != bytes.Compare(test.expected, res.Body.Bytes()) {
			t.Error("Error response test is not what was expected. Expected:", string(test.expected), "\nrecieved:", string(res.Body.Bytes()))
		}

	}
}
func TestServer200s(t *testing.T) {

	tests := []struct {
		table    models.Table
		expected []bson.ObjectId
	}{
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e375"),
						Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: false,
					},
					models.TablePlayer{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e333"),
						Name: "Devin", Cards: []card{newCard(THREE, SPADE), newCard(FOUR, HEART)}, Folded: false,
					},
				},
				FaceUp: []card{newCard(FIVE, SPADE), newCard(JACK, HEART), newCard(KING, CLUB), newCard(SEVEN, DIAMOND), newCard(NINE, DIAMOND)},
			},
			[]bson.ObjectId{bson.ObjectIdHex("5c6482805508c93011b4e375")},
		},
		{
			models.Table{
				Players: []models.TablePlayer{
					{
						ID:   bson.ObjectIdHex("5c6482805508c93011b4e375"),
						Name: "Brent", Cards: []card{newCard(TWO, HEART), newCard(TWO, CLUB)}, Folded: false,
					},
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
			[]bson.ObjectId{
				bson.ObjectIdHex("5c6482805508c93011b4e332"),
				bson.ObjectIdHex("5c6482805508c93011b4e375"),
			},
		},
	}

	for _, test := range tests {
		res := makeTestRequest(t, test.table)

		if res.Code != http.StatusOK {
			t.Log(res)
			t.Error("response code was not desired. Expected 200, recieved", res.Code)
		}
		var expectation, actual RespObj
		for _, id := range test.expected {
			expectation.Ids = append(expectation.Ids, []byte(id))
		}
		err := json.Unmarshal(res.Body.Bytes(), &actual)
		// resposes will be in the respObj form. Getting something else
		// should mean an error message (likely in plain text) is the response
		if err != nil {
			t.Fatal("failed to unmarshel sever response")

		}
		if numWinners := len(actual.Ids); len(expectation.Ids) != numWinners {
			t.Error("wrong number of ids in response. Expected", numWinners, "recieved", len(expectation.Ids))
		}
		for _, id := range expectation.Ids {
			found := false
			for _, a := range actual.Ids {
				if bytes.Compare(id, a) == 0 {
					found = true
					break
				}
			}
			if !found {
				t.Error("An expected id was not found in the response! Expected to find", string(id))
			}
		}
	}
}
