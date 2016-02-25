package main

import (
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"encoding/json"
)

const (
	DB_USER = "postgres"
	DB_PASSWORD = "root"
	DB_NAME = "clients"
)

type Client struct {
	Name    string
	Balance float32
}

type Deposit struct {
	User   int
	Amount int
}

type Transfer struct {
	From   int
	To     int
	Amount int
}

var database sql.DB

func connect() sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return *db
}

func insert(db sql.DB, id int) {
	rows, err := db.Exec("INSERT INTO client(id, name, balance) VALUES($1, $2, $3)", id, "", 0);
	if (rows == nil) {
		checkErr(err)
	}
}

func addToDeposit(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var deposit Deposit
	err := decoder.Decode(&deposit)
	checkErr(err)
	if (!isExists(deposit.User)) {
		insert(database, deposit.User);
	}
	s, err1 := database.Query("update client set balance = balance + $1 where id = $2", deposit.Amount, deposit.User);
	if (s == nil) {
		checkErr(err1)
	}
}

func transfer(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var transfer Transfer
	err := decoder.Decode(&transfer)
	checkErr(err)
	if (!isExists(transfer.From)) {
		insert(database, transfer.From);
	}
	if (!isExists(transfer.To)) {
		insert(database, transfer.To);
	}
	to, err1 := database.Query("update client set balance = balance + $1 where id = $2", transfer.Amount, transfer.To);
	if (to != nil) {
		from, err1 := database.Query("update client set balance = balance - $1 where id = $2", transfer.Amount, transfer.From);
		if (from == nil) {
			checkErr(err1)
		}
	}else{
		checkErr(err1)
	}
}

func withdrawFromDep(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var deposit Deposit
	err := decoder.Decode(&deposit)
	checkErr(err)
	if (!isExists(deposit.User)) {
		insert(database, deposit.User);
	}
	res, err1 := database.Query("update client set balance = balance - $1 where id = $2", deposit.Amount, deposit.User);
	if (res == nil) {
		checkErr(err1)
	}
	checkErr(err1)
}

func isExists(id int) bool {
	fmt.Println(id)
	var idd string
	err := database.QueryRow("select id from client where id = $1", id).Scan(&idd);
	if (err == nil) {
		return true;
	}
	return false;
}

func getBalance(rw http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("user")

	var client Client
	err := database.QueryRow("select name, balance from client where id = $1", id).Scan(&client.Name, &client.Balance)

	m, err := json.Marshal(client)

	code := checkErr(err)
	fmt.Println(code)
	rw.Write(m)
}

func main() {
	database = connect()
	http.HandleFunc("/balance", getBalance)
	http.HandleFunc("/deposit", addToDeposit)
	http.HandleFunc("/withdraw", withdrawFromDep)
	http.HandleFunc("/transfer", transfer)
	http.ListenAndServe(":9000", nil)
	defer database.Close()
}

func checkErr(err error) int {
	if err != nil {
		fmt.Println(err)
		return 422
	}
	return 200
}