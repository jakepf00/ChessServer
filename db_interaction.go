package main

import "fmt"

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

func WholeGameStateToDB(wholeGameState *WholeGameState) {
	// wholeGameState.GameId

	// InsertRow()
}
