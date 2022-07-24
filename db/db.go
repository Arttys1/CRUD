/**
 * Package who manage sql database.
 */
package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type DB struct {
	_handle *sql.DB
}

func (db *DB) Init() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "bestiole",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db._handle, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return
	}

	pingErr := db._handle.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return
	}
	fmt.Println("Connected!")
}

func (db *DB) Ping() bool {
	err := db._handle.Ping()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

/**
 * function to Query the database.
 * This function is typically used with SELECT for example
 * @param query string the query to be executed.
 * @return sql.Rows the cursor of the query. and an error if any.
 */
func (db *DB) Query(query string) (*sql.Rows, error) {
	rows, err := db._handle.Query(query)
	return rows, err
}

/**
 * function to Execute a query.
 * This function is typically used with INSERT, UPDATE, DELETE for example
 * @param query string the query to be executed.
 * @return sql.Result the result of the query. and an error if any.
 */
func (db *DB) Exec(query string) (sql.Result, error) {
	result, err := db._handle.Exec(query)
	return result, err
}

func (db *DB) Close() {
	db._handle.Close()
	fmt.Println("database closed")
}

func (db *DB) Handle() *sql.DB {
	return db._handle
}
