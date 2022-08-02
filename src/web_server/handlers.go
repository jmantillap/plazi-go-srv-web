package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Word Platzi!")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Word desde API HOME!")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decorer := json.NewDecoder(r.Body)
	var metaData MetaData
	err := decorer.Decode(&metaData)
	if err != nil {
		fmt.Fprintf(w, "error : %v", err)
		return
	}
	//fmt.Println(metaData["name"])
	fmt.Fprintf(w, "Payload %v\n", metaData)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decorer := json.NewDecoder(r.Body)
	var user User
	err := decorer.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error : %v", err)
		return
	}
	fmt.Println(user.Name)
	fmt.Fprintf(w, "Payload %v\n", user)
}
