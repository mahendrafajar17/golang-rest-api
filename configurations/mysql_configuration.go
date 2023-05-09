package configurations

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func InitMySQL(config *Config) *sql.DB {
	db, err := sql.Open("mysql", config.DB.User+":"+config.DB.Password+"@tcp("+config.DB.Host+":"+strconv.Itoa(config.DB.Port)+")/"+config.DB.Database)
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}
	return db
}
