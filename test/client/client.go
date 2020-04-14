package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"../../lib/models"
)

type User struct {
	Id    int ``
	Name  string
	Email string
}

func get() {
	resp, err := http.Get("http://localhost:8081/user/1/get")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func post() {
	a = models.SmfRegistraion{}
	user := &User{
		Id:    5,
		Name:  "hau",
		Email: "hau@viettel.com.vn",
	}
	// requestBody, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	// idUser := strconv.Itoa(user.Id)
	// fmt.Println(idUser)
	// htpp1 := "http://localhost:8081/user/" + idUser + "/post"
	// fmt.Println(htpp1)
	resp, err := http.Post("http://localhost:8081/user/post", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func main() {
	// get()
	post()
}
