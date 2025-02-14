package es

import (
	"crypto/tls"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticsearchClient() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		APIKey: "V1B3Rjk1UUItM2h2dUsxUGQzOFQ6LTF1LWhxUDVSY3lGSm04bEJscTVCZw",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Disable TLS verification
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic("failed to connect to Elasticsearch")
	}

	return es
}
