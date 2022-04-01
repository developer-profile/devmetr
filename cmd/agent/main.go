package main

import (
	"expvar"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
	"time"
)

type Gauge struct {
	Alloc,
	BuckHashSys,
	Frees,
	GCCPUFraction,
	GCSys,
	HeapAlloc,
	HeapIdle,
	HeapInuse,
	HeapObjects,
	HeapReleased,
	HeapSys,
	LastGC,
	Lookups,
	MCacheInuse,
	MCacheSys,
	MSpanInuse,
	MSpanSys,
	Mallocs,
	NextGC,
	NumForcedGC,
	NumGC,
	OtherSys,
	PauseTotalNs,
	RandomValue,
	StackInuse,
	StackSys,
	Sys,
	TotalAlloc float64
	NumGoroutine int
	Counter      int64
}

func main() {

	GetMetrics(2)

}

func GetMetrics(duration time.Duration) {
	// url to update
	baseUrl := "http://127.0.0.1:8080/UPDATE/"

	s := 0  // steps counter
	cs := 0 // send to server counter
	i := 0  // for increment base value
	// The next line goes at the start of NewMonitor()
	var goroutines = expvar.NewInt("num_goroutine")
	var rtm runtime.MemStats
	var g Gauge
	var urlString []string

	var interval = duration * time.Millisecond
	for {
		<-time.After(interval)
		// Read full mem stats
		runtime.ReadMemStats(&rtm)
		// Number of goroutines
		g.NumGoroutine = runtime.NumGoroutine()
		// The next line goes after the runtime.NumGoroutine() call
		goroutines.Set(int64(g.NumGoroutine))

		// agent metrics

		urlString = append(urlString, fmt.Sprintf("%vGAUGE/ALLOC/%v", baseUrl, float64(rtm.Alloc)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/BuckHashSys/%v", baseUrl, float64(rtm.BuckHashSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/Frees/%v", baseUrl, float64(rtm.Frees)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/GCCPUFraction/%v", baseUrl, rtm.GCCPUFraction))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/GCSys/%v", baseUrl, float64(rtm.GCSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapAlloc/%v", baseUrl, float64(rtm.HeapAlloc)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapIdle/%v", baseUrl, float64(rtm.HeapIdle)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapInuse/%v", baseUrl, float64(rtm.HeapInuse)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapObjects/%v", baseUrl, float64(rtm.HeapObjects)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapReleased/%v", baseUrl, float64(rtm.HeapReleased)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapSys/%v", baseUrl, float64(rtm.HeapSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/LastGC/%v", baseUrl, float64(rtm.LastGC)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/Lookups/%v", baseUrl, float64(rtm.Lookups)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/MCacheInuse/%v", baseUrl, float64(rtm.MCacheInuse)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/MCacheSys/%v", baseUrl, float64(rtm.MCacheSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/MSpanInuse/%v", baseUrl, float64(rtm.MSpanInuse)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/MSpanSys/%v", baseUrl, float64(rtm.MSpanSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/Mallocs/%v", baseUrl, float64(rtm.Mallocs)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/NextGC/%v", baseUrl, float64(rtm.NextGC)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/NumForcedGC/%v", baseUrl, float64(rtm.NumForcedGC)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/NumGC/%v", baseUrl, float64(rtm.NumGC)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/OtherSys/%v", baseUrl, float64(rtm.OtherSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/PauseTotalNs/%v", baseUrl, float64(rtm.PauseTotalNs)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/RandomValue/%v", baseUrl, rand.ExpFloat64()))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/StackInuse/%v", baseUrl, float64(rtm.StackInuse)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/StackSys/%v", baseUrl, float64(rtm.StackSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/Sys/%v", baseUrl, float64(rtm.Sys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/TotalAlloc/%v", baseUrl, float64(rtm.TotalAlloc)))

		// Transport config
		tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}
		client := &http.Client{Transport: tr}

		s += 1
		if i < 1 {
			i += 1
			//s := fmt.Sprintf("%s is %d years old.\n", name, age)
			log.Printf("Step #%d collecting data with 2 seconds interval %d \n", s, i)
		} else {
			i = 1
			cs += 1

			for _, value := range urlString {
				fmt.Printf("%v\n", value)
				//fmt.Sprintf("%v", value)
				_, err := client.Post(value, "text/plain; utf-8", strings.NewReader(""))
				if err != nil {
					log.Fatal(err)
				}

			}

			fmt.Printf("Step #%d sending data to 127.0.0.1:8080/update: #%d \n", s, cs)
			fmt.Printf("Step #%d collecting data with 2 seconds interval %d \n", s, i)

		}
		if s > 5 {
			fmt.Printf("Total steps: %d \nTotal server update: %d \nExiting.. \n", s, cs)
			return
		}

	}
}
