package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	DbName         string
	User, Password string
	Host           string
}

func (c *Connection) formatConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.User, c.Password, c.Host, c.DbName)
}

func (c *Connection) Open() (*sql.DB, error) {
	return sql.Open("mysql", c.formatConnection())
}

func TransferMoneydb(sender, receiver, amount, date string, dbCon *sql.DB) error {
	var balance string
	err := dbCon.QueryRow("SELECT balance FROM accounts WHERE id = " + sender + ";").Scan(&balance)
	if err != nil {
		log.Fatal(err)
	}

	amountInt, _ := strconv.ParseInt(amount, 10, 64)
	balanceInt, _ := strconv.ParseInt(balance, 10, 64)

	if amountInt > balanceInt {
		return errors.New("Transfer failed: insufficient funds")
	}
	// We should do it in one querry so that money doesn't disappear due to atomicity principle
	sqlcom := "CALL transferMoney (" + sender + ", " + receiver + ", " + amount + ", '" + date + "');"
	_, err = dbCon.Exec(sqlcom)

	return nil
}

func FetchAccountsdb(dbCon *sql.DB) (*sql.Rows, error) {
	rows, err := dbCon.Query("SELECT * FROM accounts;")
	if err != nil {
		return nil, err
	}
	return rows, nil
}
