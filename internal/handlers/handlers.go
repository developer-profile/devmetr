package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func GaugeUpdate(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(requestDump))

	switch r.Method {
	case "GET":
		// сначала устанавливаем заголовок Content-Type
		// для передачи клиенту информации, кодированной в JSON
		w.Header().Set("content-type", "application/json")
		// устанавливаем статус-код 200
		w.WriteHeader(http.StatusNotAcceptable)

		log.Println("GET method not acceptable")

	case "POST":
		params := strings.Split(r.RequestURI, "/")
		log.Printf("request parsed: \n %v \n", params)

		switch params[1] {
		case "gauge":
			log.Println("Gauge")

		case "counter":
			log.Println("counter")

		}

	default:
		_, err := fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		if err != nil {
			log.Printf("Error %v", err)
		}

	}

}

func CounterUpdate(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(requestDump))

	switch r.Method {
	case "GET":
		// сначала устанавливаем заголовок Content-Type
		// для передачи клиенту информации, кодированной в JSON
		w.Header().Set("content-type", "application/json")
		// устанавливаем статус-код 200
		w.WriteHeader(http.StatusNotAcceptable)

		w.Write([]byte("Get method not acceptable"))

		log.Println("GET method not acceptable")

	case "POST":
		params := strings.Split(r.RequestURI, "/")
		log.Printf("request parsed: \n %v \n", params)

		switch params[1] {
		case "gauge":
			log.Println("Gauge")

		case "counter":
			log.Println("counter")

		}

	default:
		_, err := fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		if err != nil {
			log.Printf("Error %v", err)
		}

	}
}

func SendNotFound(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(requestDump))

	// сначала устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	w.Header().Set("content-type", "text/plain")
	// устанавливаем статус-код 200
	w.WriteHeader(http.StatusNotFound)

	w.Write([]byte("Page not fount code 404"))

	log.Println("Page not fount code 404")
}
