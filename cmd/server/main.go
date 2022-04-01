package main

import (
	"fmt"
	"github.com/developer-profile/devmetr.git/internal/handlers"
	"log"
	"net/http"
)

func main() {

	//http.Handle("/update/", middleware.Conveyor(http.HandlerFunc(handlers.LoginHandler1), handlers.Hello, handlers.StatusHandler, handlers.UpdateHandler, handlers.LoginHandler))
	http.HandleFunc("/update/gauge/", handlers.GaugeUpdate)
	http.HandleFunc("/update/counter/", handlers.CounterUpdate)
	http.HandleFunc("/", handlers.SendNotFound)

	//http.HandleFunc("/update/COUNTER/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}
