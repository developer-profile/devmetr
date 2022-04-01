package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Gauge struct {
	Alloc         float64 `json:"Alloc"`
	BuckHashSys   float64 `json:"BuckHashSys"`
	Frees         float64 `json:"Frees"`
	GCCPUFraction float64 `json:"GCCPUFraction"`
	GCSys         float64 `json:"GCSys"`
	HeapAlloc     float64 `json:"HeapAlloc"`
	HeapIdle      float64 `json:"HeapIdle"`
	HeapInuse     float64 `json:"HeapInuse"`
	HeapObjects   float64 `json:"HeapObjects"`
	HeapReleased  float64 `json:"HeapReleased"`
	HeapSys       float64 `json:"HeapSys"`
	LastGC        float64 `json:"LastGC"`
	Lookups       float64 `json:"Lookups"`
	MCacheInuse   float64 `json:"MCacheInuse"`
	MCacheSys     float64 `json:"MCacheSys"`
	MSpanInuse    float64 `json:"MSpanInuse"`
	MSpanSys      float64 `json:"MSpanSys"`
	Mallocs       float64 `json:"Mallocs"`
	NextGC        float64 `json:"NextGC"`
	NumForcedGC   float64 `json:"NumForcedGC"`
	NumGC         float64 `json:"NumGC"`
	OtherSys      float64 `json:"OtherSys"`
	PauseTotalNs  float64 `json:"PauseTotalNs"`
	RandomValue   float64 `json:"RandomValue"`
	StackInuse    float64 `json:"StackInuse"`
	StackSys      float64 `json:"StackSys"`
	Sys           float64 `json:"Sys"`
	TotalAlloc    float64 `json:"TotalAlloc"`
	NumGoroutine  int
}

func hello(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/" {
	//	http.Error(w, "404 not found.", http.StatusNotFound)
	//	return
	//}

	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// намеренно сделана ошибка в JSON
		_, err := w.Write([]byte(`{"status":"ok"}`))
		if err != nil {
			log.Panicln(err)
			os.Exit(1)
		}
	case "POST":

		//map[string]float64
		log.Println(r.RequestURI)
		params := strings.Split(r.RequestURI, "/")
		for _, value := range params {
			log.Printf("value %v", value)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/UPDATE/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
