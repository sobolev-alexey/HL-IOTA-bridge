package main

import (
	// "strconv"
	// "fmt"
	. "iota/iota"
)

type Response struct {
	Name       		string `json:"name"`
	Model       	string `json:"model"`
	Manufacturer    string `json:"manufacturer"`
}

func main() {
	Store()

	const jsonData1 = `
	    {"Name": "Alice", "Age": 25}
	    {"Name": "Bob", "Age": 22}
	`

	const jsonData2 = `
			{"Name": "Charlie", "Age": 35}
			{"Name": "Dave", "Age": 42}
	`
	// initiate IOTA transaction
	// TransferTokens()
	// Fetch()
	// trannsmiter := Init(jsonData1, nil)
	// Init(jsonData2, trannsmiter)

	// var randomNumber = Random()
	//
	// fmt.Println("randomNumber", randomNumber)
	//
	// rsp := &Response{}
	// if err := MakeRequest1("https://swapi.co/api/vehicles/" + strconv.Itoa(randomNumber), rsp); err != nil {
	// 	fmt.Println(666, err)
	// }
	// // b := []byte("My string " + strconv.Itoa(randomNumber))
	//
	// result := rsp.Name + " | " + strconv.Itoa(randomNumber)
	// fmt.Println("result", result)
}

func Main() string {
  return "Hello, world."
}
