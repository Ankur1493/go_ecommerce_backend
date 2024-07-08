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

func getHealhStatus(w http.ResponseWriter, _ *http.Request) {
	response := []byte("Server is up and running...")
	_, err := w.Write(response)
	if err != nil {
		fmt.Println("error in payment")
	}
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	response := []byte("payment happened")
	_, err := w.Write(response)
	if err != nil {
		fmt.Println("error in payment")
	}
}
