package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/developer-profile/devmetr/internal/server/handlers"
	"github.com/developer-profile/devmetr/internal/server/repository"
	"github.com/developer-profile/devmetr/internal/server/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gorilla/mux"
)

func TestRootPath(t *testing.T) {

	repo := repository.NewRepoMem()
	bl := usecase.NewMetricBusinessLogic(repo)
	handlers := handlers.NewMetricHandler(bl, "127.0.0.1")

	mux := mux.NewRouter()
	mux.HandleFunc("/", handlers.GetAll).Methods("GET")

	srv := httptest.NewServer(mux)
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/", srv.URL))

	assert.Nil(t, err, "Get error should be nil")
	assert.Equal(t, http.StatusOK, res.StatusCode, "Status code should by 200")

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	assert.Nil(t, err, "Read body error should be nil")
	assert.Equal(t, "<html>\n<body>\n\t<h1>Metrics</h1>\n\t\n</body>\n</html>\n", string(body), "Check body")
}

func TestSetMetric(t *testing.T) {

	repo := repository.NewRepoMem()
	bl := usecase.NewMetricBusinessLogic(repo)
	handlers := handlers.NewMetricHandler(bl, "127.0.0.1")

	mux := mux.NewRouter()
	mux.HandleFunc("/update/{mtype}/{name}/{value}", handlers.SetMetric).Methods("POST")

	srv := httptest.NewServer(mux)
	defer srv.Close()

	tests := []struct {
		name       string
		url        string
		statusCode int
	}{
		{"Positive gauge metric", "/update/gauge/metricGauge/12", http.StatusOK},
		{"Positive counter metric", "/update/counter/metricCounter/12", http.StatusOK},
		{"Wrong metric type", "/update/wrong/metricGauge/12", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s%s", srv.URL, tt.url)
			res, err := http.Post(url, "text/plain", nil)
			if err != nil {
				//Handle the error here.
				return
			}
			defer res.Body.Close()
			//Read and parse response body here
			assert.Nil(t, err, "Get error should be nil")
			assert.Equal(t, tt.statusCode, res.StatusCode, "Check status code")
		})

	}
}

func TestGetMetric(t *testing.T) {

	repo := repository.NewRepoMem()
	bl := usecase.NewMetricBusinessLogic(repo)
	handlers := handlers.NewMetricHandler(bl, "127.0.0.1")

	mux := mux.NewRouter()
	mux.HandleFunc("/update/{mtype}/{name}/{value}", handlers.SetMetric).Methods("POST")
	mux.HandleFunc("/value/{mtype}/{name}", handlers.GetMetric).Methods("GET")

	srv := httptest.NewServer(mux)
	defer srv.Close()

	tests := []struct {
		name       string
		postURL    string
		getURL     string
		value      string
		statusCode int
	}{
		{"Positive gauge metric", "/update/gauge/metricGauge/", "/value/gauge/metricGauge", "12", http.StatusOK},
		{"Positive counter metric", "/update/counter/metricCounter/", "/value/counter/metricCounter", "12", http.StatusOK},
		{"Negative counter metric", "", "/value/counter/NotExist", "\n", http.StatusNotFound},
		{"Negative notExist type", "", "/value/NotExist/metricCounter", "\n", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.statusCode == http.StatusOK {
				url := fmt.Sprintf("%s%s%s", srv.URL, tt.postURL, tt.value)
				res, err := http.Post(url, "text/plain", nil)
				if err != nil {
					//Handle the error here.
					return
				}
				defer res.Body.Close()
				require.Nil(t, err, "Get error should be nil")
				require.Equal(t, tt.statusCode, res.StatusCode, "Check status code")
			}
			res, err := http.Get(fmt.Sprintf("%s%s", srv.URL, tt.getURL))
			if err != nil {
				//Handle the error here.
				return
			}
			defer res.Body.Close()
			//Read and parse response body here

			assert.Nil(t, err, "Get error should be nil")
			assert.Equal(t, tt.statusCode, res.StatusCode, "Check status code")
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			assert.Nil(t, err, "Read body error should be nil")
			assert.Equal(t, tt.value, string(body), "Check body")
		})

	}
}
