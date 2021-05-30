package db_handling

import (
	"backend_server/entities"
	"backend_server/logger"
	"errors"
	"fmt"
)

type User = entities.User

func GetUserByEmail(userEmail string) (user *User, err error) {

	db := ConnectDBORM()

	var user_ User

	err = db.Model(&user_).Where("email = ?", userEmail).Select()

	return &user_, err
}

func DeleteUserByEmail(userEmail string) (err error) {

	db := ConnectDBORM()

	if exists := CheckUserExistence(userEmail); exists {
		var user_ User
		_, err = db.Model(&user_).Where("email = ?", userEmail).Delete()

	} else {
		err = errors.New("User Not Found")
	}
	return err
}

func CheckUserExistence(userEmail string) (exists bool) {

	db := ConnectDBORM()

	var user_ User

	err := db.Model(&user_).Where("email = ?", userEmail).Select()

	if err != nil {
		exists = false
	} else {
		exists = true
	}
	return exists
}

func WriteUserToDb(user *User) (err error) {

	if !CheckUserExistence(user.Email) {
		db := ConnectDbSQL()

		sqlStatement := `INSERT INTO users (firstname, lastname, address, email)
		VALUES ($1, $2, $3, $4)
		`

		err := db.QueryRow(sqlStatement, user.Firstname, user.Lastname, user.Address, user.Email)
		fmt.Println(err)
		return nil
	} else {
		logger.InfoLogger.Println("User with email ", user.Email, "already in db and will not be written into the db.")
		err := errors.New("problem occured")
		return err
	}

}
