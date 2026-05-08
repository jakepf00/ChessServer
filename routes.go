package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
)

type ColorIdMap struct {
	White string `json:"white"` //maps to user ids
	Black string `json:"black"` //maps to user ids
}

type StartingGameStateStruct struct {
	StartingGameState ColorIdMap `json:"starting_game_state"`
	NewGameId         uuid.UUID  `json:"new_game_id"`
}

type StartingGameReq struct {
	OpponentUsername string `json:"opponent_username"`
}

func GetStartingGameState(client_user, opp_user string) StartingGameStateStruct {
	index := rand.Intn(2)

	return StartingGameStateStruct{
		StartingGameState: ColorIdMap{
			White: [2]string{client_user, opp_user}[index],
			Black: [2]string{client_user, opp_user}[1-index],
		},
		NewGameId: uuid.New(),
	}
}

func Chess(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chess")
}

// TODO: move game creation to separate method, takes two user ids as args
func GetGame(w http.ResponseWriter, r *http.Request) {
	state := GameState{
		Board: [8][8]string{
			{"r", "b", "n", "q", "k", "n", "b", "r"},
			{"p", "p", "p", "p", "p", "p", "p", "p"},
			{" ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " "},
			{"P", "P", "P", "P", "P", "P", "P", "P"},
			{"R", "B", "N", "Q", "K", "N", "B", "R"},
		},
		WhiteTurn: true, CastlingRights: 4, EnPassantSquare: 0,
	}

	user := r.Header.Get("USERNAME")

	fmt.Println(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(state)
}

func GetUsername(r *http.Request) string {
	return r.Header.Get("username")
}

func StartGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Bad request method: " + r.Method)
		// TODO: common error middleware
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "Bad request method: " + r.Method})
		return
	}

	username := GetUsername(r)

	if username == "" {
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "No username header"})
		return
	}

	var req StartingGameReq

	// NOTE: unmarshalling does not alwyas fit the req body format...
	// err := json.Unmarshal(reqBody, &reqString)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&req)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "Error unmarshalling: " + err.Error()})
		return
	}

	opp_username := req.OpponentUsername

	fmt.Println("Opp username: " + opp_username)

	json.NewEncoder(w).Encode(GetStartingGameState(username, opp_username))
}

func JoinGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chess")
}

func MakeMove(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chess")
}

func ViewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chess")
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chess")
}
