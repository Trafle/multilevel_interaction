// http://localhost:8080/fetch?id=5&balance=1000&lastOperationTime=2006.01.02%2015:04:05
// http://localhost:8080/transfer?amount=6700&sender=1&receiver=2
package main

import (
	"fmt"
	"database/sql"
  "log"
	"net/http"
	"../../db"
	"strconv"
	"time"
)

func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection {
		DbName:     "payment_system",
		User:       "ihor",
		Password:	"123",
		Host:       "localhost:3306",
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
	if err := rows.Err(); err != nil { log.Fatal(err) }
	return result
}

func StartServer() error {
	log.Print("Server started on\n127.0.0.1:", PORT)
	http.HandleFunc("/fetch", fetchAccounts)
	http.HandleFunc("/transfer", transferMoney)
	return http.ListenAndServe(":15000", nil)
}

func fetchAccounts (rw http.ResponseWriter, r *http.Request) {
	rows, err := dbCon.Query("SELECT * FROM accounts")
	if err != nil { log.Fatal(err) }
	defer rows.Close()
	responseSTR := rowsToString(rows)
	rw.Write([]byte(responseSTR))
}

func transferMoney (rw http.ResponseWriter, r *http.Request) {
	amount, ok1 := r.URL.Query() ["amount"]
	sender, ok2 := r.URL.Query() ["sender"]
	receiver, ok3 := r.URL.Query() ["receiver"]
	if(!ok1 || !ok2 || !ok3) {
		log.Fatal("not ok")
	}

	var balance string
	err := dbCon.QueryRow(sender[0]).Scan(&balance)
	if err != nil { log.Fatal(err) }

	amountInt, _ := strconv.ParseInt(amount[0], 10, 64)
	balanceInt, _ := strconv.ParseInt(balance, 10, 64)
	if ( amountInt > balanceInt) {
		rw.Write([]byte("insufficient funds"))
		return
	}
	date := time.Now()
	// We should do it in one querry so that money doesn't disappear due to atomary principle
	_, err = dbCon.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?;" +
											"UPDATE accounts SET balance = balance + ? WHERE id = ?" +
											"UPDATE accounts SET lastOperationTime = ? WHERE id = ? OR id = ?",
											amount[0], sender[0], amount[0], receiver[0], date.Format("2006.01.02 15:04:05"), sender[0], receiver[0])
	if err != nil { 
		rw.Write([]byte("couldn't update you balance"))
		log.Fatal(err)
	}
	rw.Write([]byte("Money transfered successfully"))
}