package configurations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "root"
const dbpass = ""
const dbname = "test"

func InitMySQL() *sql.DB {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}
	return db
}
