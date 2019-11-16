package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getMes) // store data to mysql and push sametime
	// http.HandleFunc("/getMessage", handle(){ get data from mysql})// get data from mysql by set fillter

	err := http.ListenAndServe(portPub, nil)
	if err != nil {
		log.Fatalf("listenandserver error: %s", err)
		return
	}
}
