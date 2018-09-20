package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func decode(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "%s %s is %d years old!\n", user.Firstname, user.Lastname, user.Age)
}

func encode(w http.ResponseWriter, r *http.Request) {
	peter := User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}
	json.NewEncoder(w).Encode(peter)
}

func main() {
	http.HandleFunc("/decode", decode)
	http.HandleFunc("/encode", encode)

	http.ListenAndServe(":8888", nil)
}
