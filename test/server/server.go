package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mux"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	idUser, err := strconv.Atoi(id)
	fmt.Println(idUser)
	if err != nil {
		fmt.Println(err)
	}
	jsonFile, err := os.Open("user.json")
	fmt.Println("jsonFile: ", jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &users)
	fmt.Println("users: ", users)
	for _, v := range users {
		if v.Id == idUser {
			fmt.Fprintln(w, v)
		}
	}

}

func addUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Vo day roi ne")
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(body)
	if err != nil {
		fmt.Println(err)
	}
	jsonFile, err := os.OpenFile("user.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonFile.WriteString(string(body))
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/post", addUser).Methods("POST")
	r.HandleFunc("/user/{id}/get", getUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))

}
