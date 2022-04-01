package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/developer-profile/devmetr.git/internal/login"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func StatusHandler1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// намеренно сделана ошибка в JSON
	_, err := w.Write([]byte(`{"status":"ok"}`))
	if err != nil {
		log.Panicln(err)
		os.Exit(1)
	}

}

func StatusHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// намеренно сделана ошибка в JSON
		_, err := w.Write([]byte(`{"status":"ok"}`))
		if err != nil {
			log.Panicln(err)
			os.Exit(1)
		}
		next.ServeHTTP(w, r)
	})
}

func LoginHandler1(w http.ResponseWriter, r *http.Request) {
	login.Login(w, r)

}

func LoginHandler(next http.Handler) http.Handler {
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		login.Login(w, r)
		next.ServeHTTP(w, r)
	}),
	)
}

func UpdateHandler1(w http.ResponseWriter, r *http.Request) {
	// сначала устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	w.Header().Set("content-type", "application/json")
	// устанавливаем статус-код 200
	w.WriteHeader(http.StatusOK)

	resp, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// If you want a good line of code to get both query or form parameters
	// you can do the following:
	param1 := r.FormValue("Sys")
	fmt.Fprintf(w, "Parameter1:  %s ", param1)

	//to get a path parameter using the standard library simply
	param2 := strings.Split(r.URL.Path, "/")

	// make sure you handle the lack of path parameters
	if len(param2) > 4 {
		fmt.Fprintf(w, " Parameter2:  %s", param2[5])
	}
	//subj := Gauge{}
	// пишем тело ответа
	w.Write(resp)
	fmt.Println(string(resp))

}

func UpdateHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// сначала устанавливаем заголовок Content-Type
		// для передачи клиенту информации, кодированной в JSON
		w.Header().Set("content-type", "application/json")
		// устанавливаем статус-код 200
		w.WriteHeader(http.StatusOK)

		resp, err := json.Marshal(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		// If you want a good line of code to get both query or form parameters
		// you can do the following:
		param1 := r.FormValue("Sys")
		fmt.Fprintf(w, "Parameter1:  %s ", param1)

		//to get a path parameter using the standard library simply
		param2 := strings.Split(r.URL.Path, "/")

		// make sure you handle the lack of path parameters
		if len(param2) > 4 {
			fmt.Fprintf(w, " Parameter2:  %s", param2[5])
		}
		//subj := Gauge{}
		// пишем тело ответа
		w.Write(resp)
		fmt.Println(string(resp))
		next.ServeHTTP(w, r)
	})
}

func Hello1(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(requestDump))

	switch r.Method {
	case "GET":

	case "POST":

		//log.Println(r.RequestURI)
		params := strings.Split(r.RequestURI, "/")
		valueCount := true
		valueName := ""
		valueData := ""
		for _, value := range params {

			if value != "" && value != "update" && value != "GAUGE" {
				if valueCount {
					valueName = value
					valueCount = false
				} else {
					valueData = value
					valueCount = true
				}

			}
			if valueData != "" {

				log.Printf("%v: %v", valueName, valueData)
				valueName = ""
				valueData = ""
			}
		}

	default:
		_, err := fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		if err != nil {
			log.Printf("Error %v", err)
		}
	}
}

func Hello(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(requestDump))

		switch r.Method {
		case "GET":

		case "POST":

			//log.Println(r.RequestURI)
			params := strings.Split(r.RequestURI, "/")
			valueCount := true
			valueName := ""
			valueData := ""
			for _, value := range params {

				if value != "" && value != "update" && value != "GAUGE" {
					if valueCount {
						valueName = value
						valueCount = false
					} else {
						valueData = value
						valueCount = true
					}

				}
				if valueData != "" {

					log.Printf("%v: %v", valueName, valueData)
					valueName = ""
					valueData = ""
				}
			}

		default:
			_, err := fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
			if err != nil {
				log.Printf("Error %v", err)
			}
		}
		next.ServeHTTP(w, r)
	})
}
