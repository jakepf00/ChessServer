package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"strings"
)

var game_map = make(map[string]*GameState)

type ColorIdMap struct {
	White string `json:"white"` //maps to user ids
	Black string `json:"black"` //maps to user ids
}

// TODO: better name... the whole state of a game... who is playing the game
// To be stored in the database
type WholeGameState struct {
	GameId    uuid.UUID  `json:"game_id"`
	GameState GameState  `json:"game_state"`
	Users     ColorIdMap `json:"users"`
}

type StartingGameReq struct {
	OpponentUsername string `json:"opponent_username"`
}

type MakeMoveReq struct {
	GameId   uuid.UUID `json:"game_id"`
	StartRow int       `json:"start_row"`
	StartCol int       `json:"start_col"`
	EndRow   int       `json:"end_row"`
	EndCol   int       `json:"end_col"`
}

func GetStartingGameState(client_user, opp_user string) WholeGameState {
	index := rand.Intn(2)

	return WholeGameState{
		Users: ColorIdMap{
			White: [2]string{client_user, opp_user}[index],
			Black: [2]string{client_user, opp_user}[1-index],
		},
		GameId: uuid.New(),
		GameState: GameState{
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
		},
	}
}

// TODO: move game creation to separate method, takes two user ids as args
func GetGame(w http.ResponseWriter, r *http.Request) {
	state := GameState{
		Board: [8][8]string{
			{"r", "n", "b", "q", "k", "b", "n", "r"},
			{"p", "p", "p", "p", "p", "p", "p", "p"},
			{" ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " "},
			{"P", "P", "P", "P", "P", "P", "P", "P"},
			{"R", "N", "B", "Q", "K", "B", "N", "R"},
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

	starting_game_state := GetStartingGameState(username, opp_username)

	game_map[starting_game_state.GameId.String()] = &starting_game_state.GameState

	json.NewEncoder(w).Encode(starting_game_state)
}

func JoinGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chess")
}

func MakeMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Bad request method: " + r.Method)
		// TODO: common error middleware
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "Bad request method: " + r.Method})
		return
	}

	var req MakeMoveReq

	// NOTE: unmarshalling does not alwyas fit the req body format...
	// err := json.Unmarshal(reqBody, &reqString)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&req)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "Error unmarshalling: " + err.Error()})
		return
	}

	game_state, ok := game_map[req.GameId.String()]

	if !ok {
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "Could not find game state"})
		return
	}

	game_state.Board[req.EndRow][req.EndCol] = game_state.Board[req.StartRow][req.StartCol]
	game_state.Board[req.StartRow][req.StartCol] = " "
	game_state.WhiteTurn = game_state.WhiteTurn != true // flips it

	json.NewEncoder(w).Encode(game_state)
}

func ViewGame(w http.ResponseWriter, r *http.Request) {
	game_id := strings.TrimPrefix(r.URL.Path, "/ViewGame/")

	w.Header().Add("Content-Type", "text/html")

	// fmt.Println(state.GetHtml())

	state, ok := game_map[game_id]

	if !ok {
		json.NewEncoder(w).Encode(map[string]string{"ERROR": "Could not find game id"})
		return
	}

	fmt.Fprintf(w, "%s", state.GetHtml())
	// http.ServeFile(w, r, "static/chess.html")
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chess")
}
