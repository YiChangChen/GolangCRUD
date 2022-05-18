package sqlservice

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"crud.com/crud/infrastructures/database"
	"crud.com/crud/model"
)

type MsSql struct {
}

func (svc *MsSql) CreateUser(model *model.User) (int64, error) {
	var err error
	err = database.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}

	ctx := context.Background()

	if database.Db == nil {
		err = errors.New("GoDb: db is null")
		return -1, err
	}

	err = database.Db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := `
      INSERT INTO UserProfile VALUES (@Id, @Username, @Email, @Phone);   	      
    `

	stmt, err := database.Db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	result, err := database.Db.ExecContext(ctx, tsql,
		sql.Named("Id", model.ID),
		sql.Named("Username", model.Username),
		sql.Named("Email", model.Email),
		sql.Named("Phone", model.Phone))
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
