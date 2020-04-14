package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string
	Email string
}

func main() {
	user := User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	// initialize http client
	client := &http.Client{}

	// marshal User to json
	json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, "http://api.example.com/v1/user", bytes.NewBuffer(json))
	if err != nil {
		//Xem co conflic khong nhe
		panic(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//Co tinh tao conflic nhe
	fmt.Println(resp.StatusCode)
}
