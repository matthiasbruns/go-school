//START_1 OMIT
db, err := sql.Open(driver, dataSourceName)
//END_1 OMIT

//START_2 OMIT
result, err := db.ExecContext(ctx, // HL2
	"INSERT INTO cars (doors, color) VALUES ($1, $2)", 2, "red",
)
//END_2 OMIT

//START_3 OMIT
rows, err := db.QueryContext(ctx, "SELECT color FROM cars WHERE doors = $1", count) // HL3
if err != nil {
	log.Fatal(err)
}
defer rows.Close() // HL3

for rows.Next() { // HL3
	var color string
	if err := rows.Scan(&color); err != nil { // HL3
		log.Fatal(err)
	}
	fmt.Printf("%s is %d\n", color, count)
}

if err := rows.Err(); err != nil {
	log.Fatal(err)
}
//END_3 OMIT

//START_4 OMIT
var doors int64
err := db.QueryRowContext(ctx, "SELECT doors FROM cars WHERE color = $1", color).Scan(&doors) // HL4
//END_4 OMIT

//START_5 OMIT
doors := 27
stmt, err := db.PrepareContext(ctx, "SELECT color FROM cars WHERE doors = $1") // HL5
if err != nil {
	log.Fatal(err)
}
rows, err := stmt.Query(doors)
// process rows
//END_5 OMIT

//START_6 OMIT
tx, err := db.BeginTx(ctx, nil) // HL6
if err != nil {
	log.Fatal(err)
}
//... do work
err = tx.Commit() // HL6
//END_6 OMIT
