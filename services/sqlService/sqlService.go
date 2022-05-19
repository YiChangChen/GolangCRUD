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

func (svc *MsSql) GetUser(in *model.User) (*model.User, error) {
	var err error
	err = database.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	ctx := context.Background()

	if database.Db == nil {
		err = errors.New("GoDb: db is null")
		return nil, err
	}

	err = database.Db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := `SELECT *
			 FROM UserProfile
			 WHERE Id = @Id
	`

	var result model.User
	err = database.Db.Get(&result, tsql, sql.Named("Id", in.ID))
	if err != nil {
		return nil, err
	}

	database.CloseConnection()

	return &result, nil
}

func (svc *MsSql) CreateUser(model *model.User) (int64, error) {
	var err error
	err = database.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
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

	result, err := database.Db.ExecContext(ctx, tsql,
		sql.Named("Id", model.ID),
		sql.Named("Username", model.Username),
		sql.Named("Email", model.Email),
		sql.Named("Phone", model.Phone))
	if err != nil {
		return -1, err
	}

	database.CloseConnection()

	return result.RowsAffected()
}

func (svc *MsSql) UpdateUser(model *model.User) (int64, error) {
	var err error
	err = database.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
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

	tsql := `UPDATE UserProfile 
			 SET Email = @Email, 
			 Phone = @Phone,
			 Username = @Username 
			 WHERE Id = @Id;`

	result, err := database.Db.ExecContext(ctx, tsql,
		sql.Named("Id", model.ID),
		sql.Named("Username", model.Username),
		sql.Named("Email", model.Email),
		sql.Named("Phone", model.Phone))
	if err != nil {
		return -1, err
	}

	database.CloseConnection()

	return result.RowsAffected()
}

func (svc *MsSql) DeleteUser(model *model.User) (int64, error) {
	var err error
	err = database.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
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

	tsql := `DELETE FROM UserProfile
			 WHERE Id = @Id;`

	result, err := database.Db.ExecContext(ctx, tsql,
		sql.Named("Id", model.ID))
	if err != nil {
		return -1, err
	}

	database.CloseConnection()

	return result.RowsAffected()
}
