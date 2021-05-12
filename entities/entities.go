package entities

// wenn man bestandteile kleine schreibt werden diese nicht exportiert und daher nicht bei
// in json encoded because they are not exported when in small cases
// das entspricht allerdings nicht json convention
type User struct {
	UserId    string `json:"userId"`
	Email     string `json:"emailAddress"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Address   string `json:"address"`
}

type Users struct {
	Users []User
}

type Transaction struct {
	TransactionId     string `json:"transactionId"`
	TransactionUserId int    `json:"transactionUserId"`
	//expense or income
	TransactionCategory string `json:"transactionCategory"`
	TransactionGroup    string `json:"transactionGroup"`
	TransactionAmount   int    `json:"transactionAmount"`
	TransactionDate     int    `json:"transactionDate"`
	TransactionCurrency string `json:"transactionCurrency"`
}

type Transactions struct {
	Transactions []Transaction
}

type Account struct {
	AccountId       string `json:"accountId"`
	AccountUserId   int    `json:"accountUserId"`
	AccountBalance  int    `json:"accountBalance"`
	AccountCurrency string `json:"accountCurrency"`
}

type Accounts struct {
	Accounts []Account
}
