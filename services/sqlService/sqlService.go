package sqlservice

import (
	"database/sql"
	"fmt"

	"crud.com/crud/infrastructures/database"
	"crud.com/crud/models"
)

type Dto struct {
}

func (m *Dto) GetUser(in string) (*models.User, error) {
	msSql := database.NewMsSql()
	err := msSql.CheckConnect()
	if err != nil {
		return nil, err
	}

	tsql := `SELECT *
			 FROM UserProfile
			 WHERE Id = @Id
	`

	var result models.User
	err = msSql.Db.Get(&result, tsql, sql.Named("Id", in))
	if err != nil {
		return nil, err
	}

	msSql.CloseConnection()

	return &result, nil
}

func (m *Dto) CreateUser(model *models.User) (int64, error) {
	msSql := database.NewMsSql()
	err := msSql.CheckConnect()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}

	tsql := `
      INSERT INTO UserProfile VALUES (@Id, @Username, @Email, @Phone);   	      
    `

	result, err := msSql.Db.ExecContext(msSql.Ctx, tsql,
		sql.Named("Id", model.ID),
		sql.Named("Username", model.Username),
		sql.Named("Email", model.Email),
		sql.Named("Phone", model.Phone))
	if err != nil {
		return -1, err
	}

	msSql.CloseConnection()

	return result.RowsAffected()
}

func (m *Dto) UpdateUser(model *models.User) (int64, error) {
	msSql := database.NewMsSql()
	err := msSql.CheckConnect()
	if err != nil {
		return -1, err
	}

	tsql := `UPDATE UserProfile 
			 SET Email = @Email, 
			 Phone = @Phone,
			 Username = @Username 
			 WHERE Id = @Id;`

	result, err := msSql.Db.ExecContext(msSql.Ctx, tsql,
		sql.Named("Id", model.ID),
		sql.Named("Username", model.Username),
		sql.Named("Email", model.Email),
		sql.Named("Phone", model.Phone))
	if err != nil {
		return -1, err
	}

	msSql.CloseConnection()

	return result.RowsAffected()
}

func (m *Dto) DeleteUser(model *models.User) (int64, error) {
	msSql := database.NewMsSql()
	err := msSql.CheckConnect()
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}

	tsql := `DELETE FROM UserProfile
			 WHERE Id = @Id;`

	result, err := msSql.Db.ExecContext(msSql.Ctx, tsql,
		sql.Named("Id", model.ID))
	if err != nil {
		return -1, err
	}

	msSql.CloseConnection()

	return result.RowsAffected()
}
