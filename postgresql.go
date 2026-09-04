package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

// NON-QUERIES
func DBExecute(sql string) error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

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

func Test() {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
