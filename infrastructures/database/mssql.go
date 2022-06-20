package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

var server = "localhost"
var port = 1433
var user = ""
var password = ""
var database = "GoDb"

// var Db *sqlx.DB

type MsSql struct {
	Db  *sqlx.DB
	Ctx context.Context
}

func NewMsSql() *MsSql {
	return &MsSql{}
}

func (c *MsSql) OpenConnection() error {
	// Build connection string
	// connString := fmt.Sprintf("server=%s;database=%s;trusted_connection=yes", server, database)
	connString := "server=localhost;database=GoDb;integrated security=SSPI;"
	var err error

	// Create connection pool

	c.Db, err = sqlx.Connect("sqlserver", connString) //sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Printf("Error creating connection pool: %s \n", err.Error())
	}

	c.Ctx = context.Background()
	err = c.Db.PingContext(c.Ctx)
	if err != nil {
		fmt.Printf("Error connection: %s \n", err.Error())
	}
	fmt.Printf("Connected!\n")
	return err
}

func (c *MsSql) CloseConnection() error {
	err := c.Db.Close()
	if err != nil {
		fmt.Printf("Error close db connection %s \n", err.Error())
	}
	return err
}

func (c *MsSql) CheckConnect() error {
	var err error
	err = c.OpenConnection()
	if err != nil {
		return err
	}

	if c.Db == nil {
		err = errors.New("GoDb: db is null")
		return err
	}

	c.Ctx = context.Background()

	err = c.Db.PingContext(c.Ctx)
	if err != nil {
		return err
	}
	return nil
}
