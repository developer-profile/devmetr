package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(requestDump))

	switch r.Method {
	case "GET":

	case "POST":

		//map[string]float64
		log.Println(r.RequestURI)
		params := strings.Split(r.RequestURI, "/")
		for _, value := range params {
			log.Printf("value %v", value)
		}

	default:
		_, err := fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		if err != nil {
			log.Printf("Error %v", err)
		}
	}
}

func main() {
	http.HandleFunc("/UPDATE/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
