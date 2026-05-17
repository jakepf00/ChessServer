package main

import (
	"fmt"
	"strings"
)

type GameState struct {
	Board           [8][8]string `json:"Board"`
	WhiteTurn       bool         `json:"WhiteTurn"`
	CastlingRights  uint8        `json:"CastlingRights"`
	EnPassantSquare uint8        `json:"EnPassantSquare"`
}

func getStyle() string {
	return `<style>
            .chess-board { border-spacing: 0; border-collapse: collapse; }
            .chess-board td { width: 1.25em; height: 1.25em; text-align: center; font-size: 48px; line-height: 0;}
            .chess-board .light { background: #eee; }
            .chess-board .dark { background: #aaa; }
     </style>
	`
}

func getFull(rows string) string {
	return fmt.Sprintf(
		"<html>"+
			"<meta charset=\"UTF-8\">"+
			"<head>"+
			"%s"+
			"</head>"+
			"<body>"+
			"<table class=\"chess-board\">"+
			"<tbody>"+
			"%s"+
			"</tbody>"+
			"</table>"+
			"</body>"+
			"</html>",
		getStyle(),
		rows,
	)
}

func lightDark(i, j int) string {
	if (i+j)%2 == 0 {
		return "light"
	}
	return "dark"
}

func getPiece(rep string) string {
	switch rep {
	case " ":
		return " "
	case "K":
		return "\u2654"
	case "Q":
		return "\u2655"
	case "R":
		return "\u2656"
	case "B":
		return "\u2657"
	case "N":
		return "\u2658"
	case "P":
		return "\u2659"
	case "k":
		return "\u265A"
	case "q":
		return "\u265B"
	case "r":
		return "\u265C"
	case "b":
		return "\u265D"
	case "n":
		return "\u265E"
	case "p":
		return "\u265F"
	}

	return "ERR" // should never be reached
}

func (gs *GameState) GetHtml() string {
	var to_return strings.Builder
	to_return.Grow(2048)

	for i, row := range gs.Board {
		to_return.WriteString("<tr>\n")
		for j, square := range row {
			fmt.Fprintf(&to_return, "<td class=\"%s\">%s</td>", lightDark(i, j), getPiece(square))
		}
		to_return.WriteString("</tr>\n")
	}

	return getFull(to_return.String())
}
