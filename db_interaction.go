package main

import (
	"fmt"
)

func UsersToDB(user_id string) error {
	// NOTE: possible UTF-8 issue here
	if len(user_id) > 50 {
		return fmt.Errorf("User id greater than 50")
	}

	err := DBExecute("CREATE TABLE IF NOT EXISTS users(user_id VARCHAR(50));")

	if err != nil {
		return err
	}

	err = DBExecute(fmt.Sprintf("INSERT INTO users (user_id) VALUES ('%s');", user_id))

	if err != nil {
		return err
	}

	return nil
}

func WholeGameStateToDB(wholeGameState *WholeGameState) error {

	err := DBExecute(
		`CREATE TABLE IF NOT EXISTS games(
			game_id VARCHAR(36),
			user_white VARCHAR(50),
			user_black VARCHAR(50),
			board VARCHAR(64),
			white_turn BOOLEAN,
			castling_rights SMALLINT,
			en_passant_square SMALLINT
		);`,
	)

	if err != nil {
		return err
	}

	var white_turn string
	if wholeGameState.GameState.WhiteTurn {
		white_turn = "TRUE"
	} else {
		white_turn = "FALSE"
	}

	err = DBExecute(
		fmt.Sprintf(
			`INSERT INTO games (game_id, user_white, user_black, board, white_turn, castling_rights, en_passant_square) VALUES ('%s', '%s', '%s', '%s', %s, %d, %d);`,
			wholeGameState.GameId.String(),
			wholeGameState.Users.White,
			wholeGameState.Users.Black,
			wholeGameState.GameState.GetBoardString(),
			white_turn,
			wholeGameState.GameState.CastlingRights,
			wholeGameState.GameState.EnPassantSquare,
		),
	)

	if err != nil {
		return err
	}

	return nil
}
