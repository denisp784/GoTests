package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	Name    string
	Balance float32
}

func main()  {
	testReq()
}

func testReq()  {
	resp, err := http.Get("http://localhost:9000/balance?user1");
	checkErr(err)
	decoder := json.NewDecoder(resp.Body);
	var client Client;
	err1 := decoder.Decode(&client)
	checkErr(err1)
	fmt.Println("Name: ",client.Name, " balance: ", client.Balance);
	resp1, err := http.PostForm("http://localhost:9000/deposit", url.Values{"user" : {"1"}, "amount" : {"100"}});
	fmt.Println("code: ", resp1.StatusCode)
	resp2, err := http.PostForm("http://localhost:9000/withdraw", url.Values{"user" : {"1"}, "amount" : {"100"}});
	fmt.Println("code: ", resp2.StatusCode)
	resp3, err := http.PostForm("http://localhost:9000/transfer", url.Values{"From" : {"1"}, "To" : {"2"}, "amount" : {"100"}});
	fmt.Println("code: ", resp3.StatusCode)
}

func checkErr(err error){
	if err != nil {
		fmt.Println(err)
	}
}