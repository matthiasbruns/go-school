import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...

db, err := sql.Open("mysql", "root:mypassword@tcp(localhost:3306)/goschool") // HL1
if err != nil {
	panic(err)
}