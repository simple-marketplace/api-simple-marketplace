package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func SearchProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Should Use Get Method", http.StatusMethodNotAllowed)
		return
	}

	esClient := r.Context().Value("es").(*elasticsearch.Client)
	query := `{ "query": { "match_all": {} }, "_source": ["name"] }`
	res, err := esClient.Search(
		esClient.Search.WithIndex("products"),
		esClient.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close() // Always close the response body

	// Check if Elasticsearch returned an error
	if res.IsError() {
		log.Fatalf("Error response from Elasticsearch: %s", res.String())
	}

	// Parse the JSON response
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the response (for debugging)
	fmt.Printf("%v\n", result)

	// Extracting hits
	// hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	// for _, hit := range hits {
	// 	doc := hit.(map[string]interface{})["_source"]
	// 	fmt.Println("Document:", doc)
	// }

	w.WriteHeader(http.StatusOK)

}
