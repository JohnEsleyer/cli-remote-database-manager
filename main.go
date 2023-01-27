package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var username, password, name, host, port, tableName string

	fmt.Print("db_username:")
	fmt.Scan(&username)

	fmt.Print("db_password:")
	fmt.Scan(&password)

	fmt.Print("db_name:")
	fmt.Scan(&name)

	fmt.Print("db_host:")
	fmt.Scan(&host)

	fmt.Print("db_port:")
	fmt.Scan(&port)

	fmt.Print("db_table_name (table to display rows):")
	fmt.Scan(&tableName)
	fmt.Println("----------------------------------- ")

	// Connect to the RDS database
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// Execute the query to select all rows from the table
	rows, err := db.Query("SELECT * FROM " + tableName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	// Get the column names
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the column names
	fmt.Println(columns)

	// Get the values of each row
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Iterate through the rows and print the values
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var value string
		for i, col := range values {

			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
}
