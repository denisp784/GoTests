package main

import (
	"testing"
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	Name    string
	Balance float32
}

func TestReq1(t *testing.T)  {
	resp, err := http.Get("http://localhost:9000/balance?user1");
	checkErr(err)
	decoder := json.NewDecoder(resp.Body);
	var client Client;
	err1 := decoder.Decode(&client)
	checkErr(err1)
	if (err != nil || resp.StatusCode != 200){
		t.Error("code: ", resp.StatusCode)
	}
	fmt.Println("Name: ",client.Name, " balance: ", client.Balance);
}

func TestReq2(t *testing.T)  {
	resp1, err := http.PostForm("http://localhost:9000/deposit", url.Values{"user" : {"1"}, "amount" : {"100"}});
	checkErr(err)
	if (err != nil || resp1.StatusCode != 200){
		t.Error("code: ", resp1.StatusCode)
	}
	fmt.Println("code: ", resp1.StatusCode)
}

func TestReq3(t *testing.T) {
	resp2, err := http.PostForm("http://localhost:9000/withdraw", url.Values{"user" : {"1"}, "amount" : {"100"}});
	checkErr(err)
	if (err != nil || resp2.StatusCode != 200){
		t.Error("code: ", resp2.StatusCode)
	}
	fmt.Println("code: ", resp2.StatusCode)
}

func TestReq4(t *testing.T)  {
	resp3, err := http.PostForm("http://localhost:9000/transfer", url.Values{"From" : {"1"}, "To" : {"2"}, "amount" : {"100"}});
	checkErr(err)
	if (err != nil || resp3.StatusCode != 200){
		t.Error("code: ", resp3.StatusCode)
	}
	fmt.Println("code: ", resp3.StatusCode)
}

func checkErr(err error){
	if err != nil {
		fmt.Println(err)
	}
}