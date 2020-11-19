package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	arguments := os.Args[1:]
	lengthOfArguments := len(arguments)
	if lengthOfArguments != 6 {
		log.Fatal("arguments is invalid.\n" +
			"usage: $ go checkDBConnection username password url port database table")
	}
	username, password, url, port, database := arguments[0], arguments[1], arguments[2], arguments[3], arguments[4]
	dataSourceName := username + ":" + password + "@tcp(" + url + ":" + port + ")/" + database
	table := arguments[5]

	connection, error1 := sql.Open("mysql", dataSourceName)
	if error1 != nil {
		log.Fatal(error1)
	} else {
		connection.SetMaxOpenConns(1)
		rows, error2 := connection.Query("SELECT * FROM " + table + " LIMIT 1")
		if error2 != nil {
			log.Fatal(error2)
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
