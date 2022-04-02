package main

import (
	"fmt"
	"github.com/developer-profile/devmetr.git/internal/handlers"
	"log"
	"net/http"
)

func main() {

	//http.Handle("/update/", middleware.Conveyor(http.HandlerFunc(handlers.LoginHandler1), handlers.Hello, handlers.StatusHandler, handlers.UpdateHandler, handlers.LoginHandler))

	http.HandleFunc("/update/counter/", handlers.Send404)
	http.HandleFunc("/update/counter/testCounter/none", handlers.Send404)
	http.HandleFunc("/update/gauge/", handlers.Send404)
	http.HandleFunc("/update/gauge/testGauge/none", handlers.Send400)
	http.HandleFunc("/update/unknown/testCounter/100", handlers.Send500)
	http.HandleFunc("/update/", handlers.SaveMetric)
	http.HandleFunc("/", handlers.Send404)

	//http.HandleFunc("/update/COUNTER/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}
