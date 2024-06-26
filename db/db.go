package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var server = "kbs-data.database.windows.net"
var port = 1433
var password = "Cahgreedy79."
var user = "kbs-data"
var database = "kbs-data"

func NewDriver() (*sql.DB, error) {

	con := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	db, err := sql.Open("sqlserver", con)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return db, err
	}
	return db, err
}
