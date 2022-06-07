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
rows, err := db.QueryContext(ctx, "SELECT name FROM users WHERE age = $1", age)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	var name string
	if err := rows.Scan(&name); err != nil {
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
err := db.QueryRowContext(ctx, "SELECT age FROM users WHERE name = $1", name).Scan(&age)
//END_4 OMIT

//START_5 OMIT
age := 27
stmt, err := db.PrepareContext(ctx, "SELECT name FROM users WHERE age = $1")
if err != nil {
	log.Fatal(err)
}
rows, err := stmt.Query(age)
// process rows
//END_5 OMIT