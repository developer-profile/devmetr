package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	// сначала устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	w.Header().Set("content-type", "application/json")
	// устанавливаем статус-код 200
	w.WriteHeader(http.StatusOK)
	// собираем данные
	subj := Gauge{}
	// кодируем JSON
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// пишем тело ответа
	w.Write(resp)
	fmt.Println(string(resp))
}

func main() {
	http.HandleFunc("/post", JSONHandler)
	http.HandleFunc("/update", UpdateHandler)
	http.ListenAndServe(":8080", nil)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	// сначала устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	w.Header().Set("content-type", "text/plain")
	// устанавливаем статус-код 200
	resp, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Читаем данные
	//subj := Gauge{}
	// пишем тело ответа
	w.Write(resp)
	fmt.Println(string(resp))
}
