package transport

import (
	"fmt"
	"log"
	"net/http"

	"github.com/developer-profile/devmetr/internal/models"
)

type HTTPClient struct {
	client  *http.Client
	hostURL string
}

func NewHTTPClient(url string, client *http.Client) HTTPClient {
	return HTTPClient{
		client:  client,
		hostURL: url}
}

func (c HTTPClient) SendMetric(m models.Metric) error {
	url := fmt.Sprintf("http://%s/update/%s/%s/%s", c.hostURL, m.Type, m.Name, m.Value)
	log.Println(url)
	response, err := c.client.Post(url, "text/plain", nil)
	if err != nil {
		return fmt.Errorf("SendMetric: %s", err)
	}
	response.Body.Close()
	return nil
}
