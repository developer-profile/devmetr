package main

import (
	"bytes"
	"expvar"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strconv"
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
	baseURL := "http://127.0.0.1:8080/update/"
	//baseURL = strings.ToUpper(baseURL)

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

		urlString = append(urlString, fmt.Sprintf("%vGAUGE/ALLOC/%v/", baseURL, float64(rtm.Alloc)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/BuckHashSys/%v/", baseURL, float64(rtm.BuckHashSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/Frees/%v/", baseURL, float64(rtm.Frees)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/GCCPUFraction/%v/", baseURL, rtm.GCCPUFraction))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/GCSys/%v/", baseURL, float64(rtm.GCSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapAlloc/%v/", baseURL, float64(rtm.HeapAlloc)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapIdle/%v/", baseURL, float64(rtm.HeapIdle)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapInuse/%v/", baseURL, float64(rtm.HeapInuse)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapObjects/%v/", baseURL, float64(rtm.HeapObjects)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapReleased/%v/", baseURL, float64(rtm.HeapReleased)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/HeapSys/%v/", baseURL, float64(rtm.HeapSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/LastGC/%v/", baseURL, float64(rtm.LastGC)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/Lookups/%v/", baseURL, float64(rtm.Lookups)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/MCacheInuse/%v/", baseURL, float64(rtm.MCacheInuse)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/MCacheSys/%v/", baseURL, float64(rtm.MCacheSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/MSpanInuse/%v/", baseURL, float64(rtm.MSpanInuse)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/MSpanSys/%v/", baseURL, float64(rtm.MSpanSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/Mallocs/%v/", baseURL, float64(rtm.Mallocs)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/NextGC/%v/", baseURL, float64(rtm.NextGC)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/NumForcedGC/%v/", baseURL, float64(rtm.NumForcedGC)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/NumGC/%v/", baseURL, float64(rtm.NumGC)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/OtherSys/%v/", baseURL, float64(rtm.OtherSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/PauseTotalNs/%v/", baseURL, float64(rtm.PauseTotalNs)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/RandomValue/%v/", baseURL, rand.ExpFloat64()))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/StackInuse/%v/", baseURL, float64(rtm.StackInuse)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/StackSys/%v/", baseURL, float64(rtm.StackSys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/Sys/%v/", baseURL, float64(rtm.Sys)))
		urlString = append(urlString, fmt.Sprintf("%vGAUGE/TotalAlloc/%v/", baseURL, float64(rtm.TotalAlloc)))

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
				fmt.Printf("%v\n", strings.ToUpper(value))
				//fmt.Sprintf("%v", value)
				func() {
					request, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBufferString(strings.ToUpper(value)))
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					request.Header.Add("Content-Type", "text/plain")
					request.Header.Add("Content-Length", strconv.Itoa(len(strings.ToUpper(value))))
					// отправляем запрос и получаем ответ
					response, err := client.Do(request)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					// печатаем код ответа
					fmt.Println("Статус-код ", response.Status)
					defer func(Body io.ReadCloser) {
						err := Body.Close()
						if err != nil {
							log.Printf("Dropped error: %v", err)
						}
					}(response.Body)
					// читаем поток из тела ответа
					body, err1 := io.ReadAll(response.Body)
					if err1 != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					// и печатаем его
					fmt.Println(string(body))

					//resp, err := client.Post(strings.ToUpper(value), "text/plain; encoding=UTF-8", strings.NewReader(""))
					//if err != nil {
					//	log.Fatal(err)
					//}
					//defer func(Body io.ReadCloser) {
					//	err := Body.Close()
					//	if err != nil {
					//		fmt.Println(err)
					//	}
					//}(resp.Body)
					//_, err2 := io.Copy(io.Discard, resp.Body)
					//if err2 != nil {
					//	fmt.Println(err)
					//}

				}()
			}

			//fmt.Printf("Step #%d sending data to HTTP://127.0.0.1:8080/update/: #%d \n", s, cs)
			//fmt.Printf("Step #%d collecting data with 2 seconds interval %d \n", s, i)

		}
		if s > 5 {
			fmt.Printf("Total steps: %d \nTotal server update: %d \nExiting.. \n", s, cs)
			return
		}

	}
}
