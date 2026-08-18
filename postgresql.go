package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func CreateTable(table_name string) error {
	println("Start of table creation")
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	println("Beginning transaction")
	tx, err := conn.Begin(context.Background())

	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	println("Beginning execution")
	_, err = tx.Exec(context.Background(), fmt.Sprintf("create table if not exists %s();", table_name))

	if err != nil {
		return err
	}

	println("Beginning commit")
	err = tx.Commit(context.Background())

	println("After commit")
	if err == nil {
		fmt.Fprintf(os.Stderr, "Created Table!")
	} else {
		fmt.Fprintf(os.Stderr, "Did not create Table!")
	}

	return nil
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
