package dbops

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
)

var (
	dbConn *sql.DB
	err error
)

func init()  {
	dbConn, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/lelel?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

}
