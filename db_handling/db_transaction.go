package db_handling

import (
	"backend_server/entities"
)

type Transaction = entities.Transaction

// func GetTransactionsByUser(userEmail string) (user *User, err error) {

// 	db := ConnectDBORM()

// 	var user_ User

// 	err = db.Model(&user_).Where("email = ?", userEmail).Select()

// 	return &user_, err
// }

// func DeleteTransactionById(userEmail string) (err error) {

// 	db := ConnectDBORM()

// 	if exists := CheckUserExistence(userEmail); exists {
// 		var user_ User
// 		_, err = db.Model(&user_).Where("email = ?", userEmail).Delete()

// 	} else {
// 		err = errors.New("User Not Found")
// 	}
// 	return err
// }

func WriteTransactionToDb(transaction *Transaction) {

	db := ConnectDbSQL()

	sqlStatement := `INSERT INTO transactions (tr_userid, tr_amount, tr_category, tr_currency, tr_date, tr_group)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	db.QueryRow(sqlStatement, transaction.TransactionUserId, transaction.TransactionAmount, transaction.TransactionCategory, transaction.TransactionCurrency, transaction.TransactionDate, transaction.TransactionGroup)

}
