package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var usrInput int
	if len(os.Args) < 6 {
		fmt.Println("Error: Not enough arguments provided. Usage: go run main.go <dbUser> <dbPass> <dbName> <dbHost> <dbPort>")
		return
	}

	dbUser := os.Args[1]
	dbPass := os.Args[2]
	dbName := os.Args[3]
	dbHost := os.Args[4]
	dbPort := os.Args[5]

	// Connect to the RDS database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	fmt.Println("Select Action")
	fmt.Println("[1] Create Table")
	fmt.Println("[2] Delete Table")
	fmt.Println("[3] Display Table Rows")
	fmt.Println("[4] Add Table Row")
	fmt.Println("[5] Delete Table Row")
	fmt.Print("Input:")
	fmt.Scan(&usrInput)

	switch usrInput {
	case 1:
		//Create Table

		// Ask for table name
		var tableName string
		fmt.Print("Enter the name of the table to create:")
		fmt.Scan(&tableName)
		_, err = db.Exec(fmt.Sprintf("CREATE TABLE %s (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50))", tableName))
		if err != nil {
			fmt.Println("Error creating table:", err)
			return
		}
		fmt.Println("Table created successfully")

	case 2:
		//Delete the table
		var tableName string
		fmt.Print("Enter the name of the table to delete:")
		fmt.Scan(&tableName)
		_, err = db.Exec(fmt.Sprintf("DROP TABLE %s", tableName))
		if err != nil {
			fmt.Println("Error deleting table:", err)
			return
		}
		fmt.Println("Table deleted successfully")
	case 3:
		//Display all rows

		// Ask for table name
		var tableName string
		fmt.Print("Enter the name of the table:")
		fmt.Scan(&tableName)

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
	case 4:
		// Create table row

		// Ask for table name
		var tableName string
		fmt.Print("Enter the name of the table:")
		fmt.Scan(&tableName)

		var rowName string
		fmt.Print("Enter the name of the row to insert:")
		fmt.Scan(&rowName)
		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (name) VALUES ('%s')", tableName, rowName))
		if err != nil {
			fmt.Println("Error inserting row:", err)
			return
		}
		fmt.Println("Row inserted successfully")

	case 5:
		// Delete table row

		// Ask for table name
		var tableName string
		fmt.Print("Enter the name of the table row to delete:")
		fmt.Scan(&tableName)

		var rowID int
		fmt.Print("Enter the ID of the row to delete:")
		fmt.Scan(&rowID)
		_, err = db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id=%d", tableName, rowID))
		if err != nil {
			fmt.Println("Error deleting row:", err)
			return
		}
		fmt.Println("Row deleted successfully")

	}

}
