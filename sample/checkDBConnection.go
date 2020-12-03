package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	const ValidLengthOfArgs = 6
	const StartIndexOfProgramArgs = 1
	const DriverName = "mysql"

	const (
		IndexOfUsername = iota
		IndexOfPassword
		IndexOfUrl
		IndexOfPort
		IndexOfDatabase
		IndexOfTableName
	)

	arguments := os.Args[StartIndexOfProgramArgs:]
	lengthOfArgs := len(arguments)
	if lengthOfArgs != ValidLengthOfArgs {
		log.Fatal("arguments is invalid.\n" +
			"usage: $ go checkDBConnection username password url port database table")
	}
	username, password, url, port, database := arguments[IndexOfUsername], arguments[IndexOfPassword], arguments[IndexOfUrl], arguments[IndexOfPort], arguments[IndexOfDatabase]
	dataSourceName := username + ":" + password + "@tcp(" + url + ":" + port + ")/" + database
	table := arguments[IndexOfTableName]

	connection, errOfConnection := sql.Open(DriverName, dataSourceName)
	if errOfConnection != nil {
		log.Fatal(errOfConnection)
	} else {
		const MaxConnectionCount = 1
		connection.SetMaxOpenConns(MaxConnectionCount)
		rows, errOfQuery := connection.Query("SELECT * FROM " + table + " LIMIT 1")
		if errOfQuery != nil {
			log.Fatal(errOfQuery)
		} else {
			defer rows.Close()
			fmt.Println("Row list: ")
			for rows.Next() {
				fmt.Println(rows.Columns())
			}
			fmt.Println("Done.")
		}
	}
	connection.Close()
}
