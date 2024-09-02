package godatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:codingsluv@tcp(localhost:3306)/go-database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
