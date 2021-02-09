package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
_
func init() {
	sql.Register("sql", &MySQLDriver{})
}

func Test(input string) string {
	val := fmt.Sprintf("%v", input)
	return val
}

