// http://localhost:15000/fetch
// http://localhost:15000/transfer?amount=6700&sender=1&receiver=2
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func makeStructJSON(rows *sql.Rows) (map[string][]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	masterData := make(map[string][]interface{})

	for rows.Next() {

		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		for i, v := range values {

			x := v.([]byte)

			if nx, ok := strconv.ParseFloat(string(x), 64); ok == nil {
				masterData[columns[i]] = append(masterData[columns[i]], nx)
			} else if b, ok := strconv.ParseBool(string(x)); ok == nil {
				masterData[columns[i]] = append(masterData[columns[i]], b)
			} else if "string" == fmt.Sprintf("%T", string(x)) {
				masterData[columns[i]] = append(masterData[columns[i]], string(x))
			} else {
				fmt.Printf("Failed on if for type %T of %v\n", x, x)
			}

		}
	}
	return masterData, nil
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
	masterData, err := makeStructJSON(rows)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rows here!!!")

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(masterData)
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
		rw.Write([]byte("couldn't update you balance"))
		log.Fatal(err)
	}
	rw.Write([]byte("Money transfered successfully"))
}
