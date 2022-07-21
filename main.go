package main

import (
	"fmt"
	"log"
	"net/http"
)

const serverPort = ":8080"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}

func main() {
	handleRequests()
}
