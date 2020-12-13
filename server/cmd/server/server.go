// http://localhost:8080/fetch?id=5&balance=1000&lastOperationTime=2006.01.02%2015:04:05
// http://localhost:8080/transfer?amount=6700&sender=1&receiver=2
package main

import (
	"database/sql"
  "log"
	"net/http"
	"../../db"
	"strconv"
	"encoding/json"
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

type DbDao struct{
	db *sql.DB
}

type JSONData struct {
	Values []float64
	Dates []string
}

func (d *DbDao) sendJSON(sqlString string, w http.ResponseWriter) (error) {

	stmt, err := d.db.Prepare(sqlString)
	if err != nil { return err }
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {	return err }
	defer rows.Close()

	values := make([]interface{}, 2)
	scanArgs := make([]interface{}, 2)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {	return err }

		var tempDate string 
		var tempValue float64
		var myjson JSONData

		d, dok := values[0].([]byte)
		v, vok := values[1].(float64)

		if dok {
			tempDate = string(d)
			if err != nil {	return err	}
			myjson.Dates = append(myjson.Dates, tempDate)
		}

		if vok {      
			tempValue = v 
			myjson.Values = append(myjson.Values, tempValue)
			log.Println(v)
			log.Println(tempValue)

		}    

		err = json.NewEncoder(w).Encode(&myjson)
		if err != nil {	return err }
	}

	return nil 
}

func StartServer() error {
	log.Print("Server started on\n127.0.0.1:", PORT)
	http.HandleFunc("/fetch", fetchAccounts)
	http.HandleFunc("/transfer", transferMoney)
	return http.ListenAndServe(":15000", nil)
}

func fetchAccounts (rw http.ResponseWriter, r *http.Request) {
	dbc := DbDao{db: dbCon}
	dbc.sendJSON("SELECT * FROM accounts;", rw)
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

	// We should do it in one querry so that money doesn't disappear due to atomary principle
	_, err = dbCon.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?;" +
											"UPDATE accounts SET balance = balance + ? WHERE id = ?",
											amount[0], sender[0], amount[0], receiver[0])
	if err != nil { 
		rw.Write([]byte("couldn't update you balance"))
		log.Fatal(err)
	}
	rw.Write([]byte("Money transfered successfully"))
}