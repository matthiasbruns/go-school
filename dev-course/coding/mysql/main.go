package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:mypassword@tcp(localhost:52000)/goschool")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reset(context.Background(), db)
	count(context.Background(), db)
	insert(context.Background(), db)
	count(context.Background(), db)
}

func reset(ctx context.Context, db *sql.DB) {
	db.ExecContext(ctx, "DELETE FROM gophers")
}

func count(ctx context.Context, db *sql.DB) {
	rows, err := db.QueryContext(ctx, "SELECT COUNT(*) FROM gophers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Count: %d\n", count)
}

func insert(ctx context.Context, db *sql.DB) {
	tx, _ := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO gophers (name, age) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	stmt.ExecContext(ctx, "Gopher's Name", 30)

	tx.Commit()
}
