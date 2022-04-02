package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func SaveMetric(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Request Dump: %v", requestDump)
	}

	//We Read the response body on the line below.
	body, err1 := ioutil.ReadAll(r.Body)
	if err1 != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	//log.Printf("Request body: %v", sb)

	rDumpLines := strings.Split(sb, "/")

	rLines := map[string]string{rDumpLines[5]: rDumpLines[6]}

	for key, element := range rLines {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
	switch r.Method {
	case "GET":
		// сначала устанавливаем заголовок Content-Type
		// для передачи клиенту информации, кодированной в JSON
		//w.Header().Set("content-type", "application/json")
		// устанавливаем статус-код 200
		//w.WriteHeader(http.StatusNotAcceptable)

		log.Println("GET method not acceptable")

	case "POST":
		params := strings.Split(r.RequestURI, "/")
		//log.Printf("Params: \n %v \n", params)

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

func Send500(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Sorry, something wrong. Try later", 500)
}

func Send400(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Sorry, something wrong. Try later", 400)
}

func Send404(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Sorry, something wrong. Try later", 404)
}
