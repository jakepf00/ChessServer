package main

import (
	"fmt"
	"net/http"
)

const PORT = 8090

// NOTE: main page
func GetHTMLLinks() string {
	return ` <html>
			<meta charset=\"UTF-8\">
			<a href="/ViewGame">View Game</a>
			<a href="/StartGame">Start Game</a>
			<a href="/GetGames">Get Games</a>
			<a href="/GetGame">Get Game</a>
	</html>`
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", GetHTMLLinks())
}

func main() {

	// Expecting that all requests have the users' username in the header , --header "Username: <username>"

	// http.HandleFunc("/", Chess)

	http.HandleFunc("/", MainPage)

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

	fmt.Printf("Chess server running on PORT: %d\n", PORT)

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
