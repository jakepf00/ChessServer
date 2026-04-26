package main

type GameState struct {
    Board [8][8]string `json:"Board"`
    WhiteTurn bool `json:"WhiteTurn"`
    CastlingRights uint8 `json:"CastlingRights"`
    EnPassantSquare uint8 `json:"EnPassantSquare"`
}

