// http://localhost:15000/fetch
// http://localhost:15000/transfer?amount=6700&sender=1&receiver=2
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"../../db"
	_ "github.com/go-sql-driver/mysql"
)

// NewDbConnection with an sql.db arg
func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:   "payment_system",
		User:     "ihor",
		Password: "123",
		Host:     "localhost:3306",
	}
	return conn.Open()
}

var dbCon, err = NewDbConnection()

func rowsToString(rows *sql.Rows) string {
	result := ""
	col := make([]string, 0)
	col, err = rows.Columns()
	for i := 0; i < len(col); i++ {
		result += col[i] + "\t"
	}
	result += "\n"
	for rows.Next() {
		var id, balance, lastOperationTime string
		rows.Scan(&id, &balance, &lastOperationTime)
		result += fmt.Sprintf("%s\t%s\t%s\n", id, balance, lastOperationTime)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

// StartServer func
func StartServer() error {
	log.Print("Server started on\n127.0.0.1:", PORT)
	http.HandleFunc("/fetch", fetchAccounts)
	http.HandleFunc("/transfer", transferMoney)
	return http.ListenAndServe(":15000", nil)
}

func fetchAccounts(rw http.ResponseWriter, r *http.Request) {
	rows, err := db.FetchAccountsdb(dbCon)

	if err != nil {
		log.Fatal(err)
	}

	masterData := rowsToString(rows)
	rw.Write([]byte(masterData))
}

func transferMoney(rw http.ResponseWriter, r *http.Request) {
	amount, ok1 := r.URL.Query()["amount"]
	sender, ok2 := r.URL.Query()["sender"]
	receiver, ok3 := r.URL.Query()["receiver"]
	if !ok1 || !ok2 || !ok3 {
		log.Fatal("not ok")
	}

	date := time.Now().Format("2006.01.02 15:04:05")

	err := db.TransferMoneydb(sender[0], receiver[0], amount[0], date, dbCon)

	if err != nil {
		rw.Write([]byte(err.Error()))
		log.Println(err)
		return
	}
	rw.Write([]byte("Money transfered successfully"))
}
