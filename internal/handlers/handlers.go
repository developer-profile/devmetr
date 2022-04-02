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
		w.Write([]byte("err"))
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

}

func Send500(w http.ResponseWriter, r *http.Request) {
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		log.Println(body)
	}
	http.Error(w, "Sorry, something wrong. Try later. #500", 500)
}

func Send400(w http.ResponseWriter, r *http.Request) {
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		log.Println(body)
	}
	http.Error(w, "Sorry, something wrong. Try later. #400", 400)
}

func Send404(w http.ResponseWriter, r *http.Request) {
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		log.Println(body)

	}

	http.Error(w, "Sorry, something wrong. Try later. #404", 404)
}
