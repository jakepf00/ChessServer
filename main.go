package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {

	if err := godotenv.Load(); err != nil {
		println("Failure to load!")
	}

	err := UsersToDB("test_user")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Added user to db")
	}

	// Expecting that all requests have the users' username in the header , --header "Username: <username>"

	// http.HandleFunc("/", Chess)

	// StartGame - POST
	// Pass in ID of opponent (username)
	// We allow user combinations to have multiple concurrent games -- starting a game always creates a new game id and a list is also returned with game ids for this particualr user pair
	// Client: {
	//   "username" : "<username>"
	// }
	// Response: {
	//  "starting_game_state": {
	//      "white_user": <id>,
	//      "black_user": <id>
	// },
	//  "new_game_id": "1234"
	//  "existing_ids": ["5678"]
	// }
	http.HandleFunc("/StartGame", StartGame)

	// GetGames - GET
	// Response: {
	//      <username>: ["<game_id_1>", "<game_id_2>"]
	// }
	http.HandleFunc("/GetGames", GetGames)

	// REQ:  {
	//  "game_id": <id>
	// }
	// RES: {
	//  "game_state": <Game_State>
	// }
	http.HandleFunc("/GetGame", GetGame)
	http.HandleFunc("/MakeMove", MakeMove)
	http.HandleFunc("/ViewGame/", ViewGame)

	// http.Handle("/ViewGame", http.StripPrefix("/ViewGame/", http.FileServer(http.Dir("./static/chess.html"))))

	http.ListenAndServe(":8090", nil)
}
