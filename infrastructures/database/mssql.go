package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

var server = "localhost"
var port = 1433
var user = "test"
var password = "test123"
var database = "GoDb"

var Db *sql.DB

func OpenConnection() error {
	// Build connection string
	// connString := fmt.Sprintf("server=%s;database=%s;trusted_connection=yes", server, database)
	connString := "server=localhost;database=GoDb;integrated security=SSPI;"
	var err error

	// Create connection pool
	Db, err = sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Printf("Error creating connection pool: %s \n", err.Error())
	}
	ctx := context.Background()
	err = Db.PingContext(ctx)
	if err != nil {
		fmt.Printf("Error connection: %s \n", err.Error())
	}
	fmt.Printf("Connected!\n")
	return err

	// // Read employees
	// count, err := ReadEmployees()
	// if err != nil {
	// 	log.Fatal("Error reading Employees: ", err.Error())
	// }
	// fmt.Printf("Read %d row(s) successfully.\n", count)

	// // Update from database
	// updatedRows, err := UpdateEmployee("Jake", "Poland")
	// if err != nil {
	// 	log.Fatal("Error updating Employee: ", err.Error())
	// }
	// fmt.Printf("Updated %d row(s) successfully.\n", updatedRows)

	// // Delete from database
	// deletedRows, err := DeleteEmployee("Jake")
	// if err != nil {
	// 	log.Fatal("Error deleting Employee: ", err.Error())
	// }
	// fmt.Printf("Deleted %d row(s) successfully.\n", deletedRows)
}

func CloseConnection() error {
	err := Db.Close()
	if err != nil {
		fmt.Printf("Error close db connection %s \n", err.Error())
	}
	return err
}

// ReadEmployees reads all employee records
func ReadEmployees() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := Db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees;")

	// Execute query
	rows, err := Db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var name, location string
		var id int

		// Get values from row.
		err := rows.Scan(&id, &name, &location)
		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
		count++
	}

	return count, nil
}

// UpdateEmployee updates an employee's information
func UpdateEmployee(name string, location string) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := Db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("UPDATE TestSchema.Employees SET Location = @Location WHERE Name = @Name")

	// Execute non-query with named parameters
	result, err := Db.ExecContext(
		ctx,
		tsql,
		sql.Named("Location", location),
		sql.Named("Name", name))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

// DeleteEmployee deletes an employee from the database
func DeleteEmployee(name string) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := Db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("DELETE FROM TestSchema.Employees WHERE Name = @Name;")

	// Execute non-query with named parameters
	result, err := Db.ExecContext(ctx, tsql, sql.Named("Name", name))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
