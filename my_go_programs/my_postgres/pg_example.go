package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	urlExample := "postgres://mfm45656:@localhost:5432/mfm45656"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var c1 string
	var c2 string
	err = conn.QueryRow(context.Background(), "SELECT col1, col2 FROM example WHERE col1 = 'one'").Scan(&c1, &c2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(c1, c2)
}
