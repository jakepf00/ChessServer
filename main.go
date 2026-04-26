package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
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
        WhiteTurn: true, CastlingRights: 4, EnPassantSquare: 0
    }

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Chess")
    })

    http.HandleFunc("/GetBoard", func (w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(state)
    })

    http.ListenAndServe(":8090", nil)
}
