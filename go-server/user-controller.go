package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var DataBase []User

type User struct {
	Username string `json:username`
	Idade    string `json:idade`
}

func dataHandler(req *http.Request) User {
	var tempUser User
	err := json.NewDecoder(req.Body).Decode(&tempUser)

	if err != nil {
		fmt.Println("Error during JSON decoding")
	}

	return tempUser
}

func insert(user User) {
	DataBase = append(DataBase, user)
	fmt.Println(DataBase)
}
