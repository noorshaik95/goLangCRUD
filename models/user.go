package models

import (
	"errors"
	"goLandCRUD/config"
	"goLandCRUD/logger"
	"goLandCRUD/utils"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (u User) Save() error {
	query := utils.InsertUserQuery
	stmt, err := config.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Username, u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := utils.GetUserByEmailQuery
	row := config.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		logger.Info("error", err)
		return errors.New("Credentials invalid")
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
