package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", getHealhStatus)
	http.HandleFunc("/payment", createPayment)

	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal("Sorry the server is down as of now")
	}
}

func getHealhStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("health is completely fine")
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("payment happen")
	w.WriteHeader(200)
}
