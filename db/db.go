package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

func NewDriver() (*sql.DB, error) {
	server := os.Getenv("server")
	user := os.Getenv("user")
	password := os.Getenv("password")
	port := os.Getenv("port")
	portint, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	database := os.Getenv("database")

	con := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, portint, database)
	db, err := sql.Open("sqlserver", con)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return db, err
	}
	return db, err
}
