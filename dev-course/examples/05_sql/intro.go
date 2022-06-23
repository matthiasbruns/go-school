//START_1 OMIT
db, err := sql.Open(driver, dataSourceName)
//END_1 OMIT

//START_2 OMIT
result, err := db.ExecContext(ctx, // HL2
	"INSERT INTO users (name, age) VALUES ($1, $2)",
	"gopher",
	27,
)
//END_2 OMIT

//START_3 OMIT
rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age = $1", age) // HL3
if err != nil {
	log.Fatal(err)
}
defer rows.Close() // HL3

for rows.Next() { // HL3
	var name string
	if err := rows.Scan(&name); err != nil { // HL3
		log.Fatal(err)
	}
	fmt.Printf("%s is %d\n", name, age)
}

if err := rows.Err(); err != nil {
	log.Fatal(err)
}
//END_3 OMIT

//START_4 OMIT
var age int64
err := db.QueryRowContext(ctx, "SELECT age FROM users WHERE name = $1", name).Scan(&age) // HL4
//END_4 OMIT

//START_5 OMIT
age := 27
stmt, err := db.PrepareContext(ctx, "SELECT name FROM users WHERE age = $1") // HL5
if err != nil {
	log.Fatal(err)
}
rows, err := stmt.Query(age)
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