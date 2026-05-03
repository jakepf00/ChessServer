package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type ColorIdMap struct {
    white string //maps to user ids
    black string
}

func Chess (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Chess")
}

// TODO: move game creation to separate method, takes two user ids as args
func GetGame (w http.ResponseWriter, r *http.Request) {
    state := GameState{
        Board: [8][8]string{
           [8]string{ "r", "b", "n", "q", "k", "n", "b", "r" },
           [8]string{ "p", "p", "p", "p", "p", "p", "p", "p" },
           [8]string{ " ", " ", " ", " ", " ", " ", " ", " " },
           [8]string{ " ", " ", " ", " ", " ", " ", " ", " " },
           [8]string{ " ", " ", " ", " ", " ", " ", " ", " " },
           [8]string{ " ", " ", " ", " ", " ", " ", " ", " " },
           [8]string{ "P", "P", "P", "P", "P", "P", "P", "P" },
           [8]string{ "R", "B", "N", "Q", "K", "N", "B", "R" },
        },
        WhiteTurn: true, CastlingRights: 4, EnPassantSquare: 0,
    }

    user := r.Header.Get( "USERNAME" )

    fmt.Println(user)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(state)
}

func StartGame (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Chess")
}

func JoinGame (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Chess")
}

func MakeMove (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Chess")
}

func ViewGame (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Chess")
}

func GetGames (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Chess")
}
