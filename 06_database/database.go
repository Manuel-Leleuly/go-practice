package belajargolangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	// this is for learning purposes only. In the future, please set a strong password for root access
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                 // set the maximum amount of idle connections
	db.SetMaxOpenConns(100)                // set the maximum amount of connections that are open
	db.SetConnMaxIdleTime(5 * time.Minute) // set how long a connection can be idle. If the max is reached, the connection will close
	db.SetConnMaxLifetime(1 * time.Hour)   // set how long a connection can be alive. If the max is reached, the connection will close

	return db
}
