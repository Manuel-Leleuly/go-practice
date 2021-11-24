package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('Leleuly','Leleuly')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		/*
			Use Nullable type from database package in order to handle null value
		*/

		var id, name string
		var email sql.NullString // will return a struct {String string, Valid bool}
		var balance int32
		var rating float64
		var birthDate sql.NullTime // will return a struct {Time time.Time, Valid bool}
		var createdAt time.Time
		var married bool

		/*
			by default, driver MySQL for golang can query DATE, DATETIME, TIMESTAMP and convert them into []byte/[]uint8,
			where they can be converted into string and then parsed into time.Time

			if you want to prevent parsing it manually, you can use "parseTime=true" from Driver MySQL to do so.
		*/

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("===================================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created at:", createdAt)
	}
}

/*
	SQL injection is done by manipulating the SQL query using the input from user.
	This is hazzardous and should be fix immediately, otherwise the user will be able to breach the database easily.
*/
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	/*
		When username is set to "admin'; #", The resulting query will be:
		"SELECT username FROM user WHERE username = 'admin'; # AND password = 'salah' LIMIT 1"

		# tag will turn the right side of the query into a comment.

		This means the actual query is "SELECT username FROM user WHERE username ='admin'", and the password is ingored.
		User can then just input the username without having to input the password by manipulating the username input.

		The best way to avoid this is to not create a manual query where a user input is needed.
		If, however, user input is required, we can use Execute or Query method.
	*/
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "manuel' DROP TABLE user; #"
	password := "manuel"

	script := "INSERT INTO user(username, password) VALUES(?,?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "manuel@gmail.com"
	comment := "comment testing"

	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId();
	if err != nil{
		panic(err)
	}

	
	fmt.Println("Success insert new comment with id", insertId)
}

/*
	prepare statement makes sure that we don't have to connect to the db more than once
	when we have to store multiple data
*/
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?,?)"

	statement, err := db.PrepareContext(ctx, script)
	if err != nil{
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "manuel" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen ke " + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	// do transaction
	for i := 0; i < 10; i++ {
		email := "manuel" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen ke " + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id", id)
	}

	// transactionError := tx.Commit()
	transactionError := tx.Rollback() // even though the test says that it's successful, the data will not be written in the db
	if transactionError != nil{
		panic(transactionError)
	}
}