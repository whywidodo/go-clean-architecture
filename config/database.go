package config

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Open Conection
func OpenConnection() error {
	var err error
	db, err = setupConnection()

	return err
}

func setupConnection() (*sql.DB, error) {
	connection := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", DBUser, DBPass, DBName, DBHost, DBPort, SSLMode)

	fmt.Println("Connection info: ", DBDriver, connection)

	db, err := sql.Open(DBDriver, connection)
	if err != nil {
		return db, errors.New("connection closed: failed to connect to database")
	}

	return db, nil
}

// Close Connection
func CloseConnection() {
	db.Close()
}

// Database Connection
func DBConnection() *sql.DB {
	return db
}
