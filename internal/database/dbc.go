package dbc

import (
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

// DB connection data
const (
    host     = "localhost"
    port     = 5432
    user     = "jmanuel"
    password = "jmanuel"
    dbname   = "crud1"
)

var db *sql.DB
var err error

// DB connection function, returns a connection db.SQL (db) from pq driver
func ConnectDB() *sql.DB {
  // Db connection string
  psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

  // Opens the database using the 'postgres' driver and the 'psqlconn' connection string
  db, err := sql.Open("postgres", psqlconn)
  CheckError(err)

  // Check DB level errors
  err = db.Ping()
  CheckError(err)

  // Returns the connection to the DB to be used
  return db
}

// General purpose error checking function
func CheckError(err error) {
  if err != nil {
    panic(err.Error())
  }
}
