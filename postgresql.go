package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

var _connection *pgx.Conn

// TODO: future if necessary can make this into a connection map should we have to deal with multiple databases
// TODO: error...
func getConnection() *pgx.Conn {
	if _connection != nil {
		return _connection
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// TODO: how to close connection
	// defer conn.Close(context.Background())

	return conn
}

// NON-QUERIES
func DBExecute(sql string) error {
	conn := getConnection()

	tx, err := conn.Begin(context.Background())

	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), sql)

	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())

	return err
}

func QueryRow(sql string, results ...any) error {
	conn := getConnection()

	err := conn.QueryRow(context.Background(), sql).Scan(results...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}

	return nil
}
