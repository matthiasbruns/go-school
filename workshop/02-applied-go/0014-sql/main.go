package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/go-school")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = db.Close()
	}()

	ctx := context.Background()
	if err := createTable(ctx, db); err != nil {
		panic(err)
	}

	// insert

	_, err = db.ExecContext(ctx, "insert into recipes(name, ingredients, instructions) values(?, ?, ?)", "Spaghetti Carbonara", "Spaghetti, Guanciale, Eggs, Pecorino Cheese, Black Pepper", "Cook the spaghetti, fry the guanciale, mix the eggs with the cheese, mix everything together")
	if err != nil {
		panic(err)
	}

	// query

	rows, err := db.QueryContext(ctx, "select * from recipes")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var id int
		var name, ingredients, instructions string
		if err := rows.Scan(&id, &name, &ingredients, &instructions); err != nil {
			panic(err)
		}

		fmt.Printf("ID: %d\nName: %s\nIngredients: %s\nInstructions: %s\n", id, name, ingredients, instructions)
	}

	// query row
	var id int
	var name, ingredients, instructions string
	if err := db.QueryRowContext(ctx, "select * from recipes where id = ?", 1).Scan(&id, &name, &ingredients, &instructions); err != nil {
		panic(err)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	_, err = tx.ExecContext(ctx, "insert into recipes(name, ingredients, instructions) values(?, ?, ?)", "Spaghetti Carbonara", "Spaghetti, Guanciale, Eggs, Pecorino Cheese, Black Pepper", "Cook the spaghetti, fry the guanciale, mix the eggs with the cheese, mix everything together")
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	_, err = tx.ExecContext(ctx, "insert into recipes(name, ingredients, instructions) values(?, ?, ?)", "Spaghetti Bolognese", "Spaghetti, Minced Meat, Tomato Sauce, Onion, Garlic", "Cook the spaghetti, fry the minced meat, add the tomato sauce, add the onion and garlic, mix everything together")
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

func createTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, `create table if not exists recipes(
    		id int primary key auto_increment,
			name varchar(255) not null,
			ingredients text not null,
			instructions text not null
	)`)

	return err
}
