package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

var server = "localhost"
var port = 1433
var user = ""
var password = ""
var database = "GoDb"

var Db *sqlx.DB

func OpenConnection() error {
	// Build connection string
	// connString := fmt.Sprintf("server=%s;database=%s;trusted_connection=yes", server, database)
	connString := "server=localhost;database=GoDb;integrated security=SSPI;"
	var err error

	// Create connection pool

	Db, err = sqlx.Connect("sqlserver", connString) //sql.Open("sqlserver", connString)
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
}

func CloseConnection() error {
	err := Db.Close()
	if err != nil {
		fmt.Printf("Error close db connection %s \n", err.Error())
	}
	return err
}
