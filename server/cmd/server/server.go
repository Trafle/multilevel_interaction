// http://localhost:8080/fetch?id=5&balance=1000&lastOperationTime=2006.01.02%2015:04:05
// http://localhost:8080/transfer?amount=6700&sender=1&receiver=2
package main

import (
    "log"
		"net/http"
)

const PORT = ":8080"

func main() {
	log.Print("Server started on\n127.0.0.1", PORT)
	http.HandleFunc("/fetch", fetchAccount)
	http.HandleFunc("/transfer", transferMoney)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func fetchAccount (rw http.ResponseWriter, r *http.Request) {
	id, ok1 := r.URL.Query() ["id"]
	balance, ok2 := r.URL.Query() ["balance"]
	lastOperationTime, ok3 := r.URL.Query() ["lastOperationTime"]
	if(!ok1 || !ok2 || !ok3) {
		log.Fatal("not ok")
	}
	log.Print(id[0])
	log.Print(balance[0])
	log.Print(lastOperationTime[0])
}

func transferMoney (rw http.ResponseWriter, r *http.Request) {
	amount, ok1 := r.URL.Query() ["amount"]
	sender, ok2 := r.URL.Query() ["sender"]
	receiver, ok3 := r.URL.Query() ["receiver"]
	if(!ok1 || !ok2 || !ok3) {
		log.Fatal("not ok")
	}
	log.Print(amount[0])
	log.Print(sender[0])
	log.Print(receiver[0])
}