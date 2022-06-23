import (
	"database/sql"

	// ..

	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/go-sql-driver/mysql" // HL1
)