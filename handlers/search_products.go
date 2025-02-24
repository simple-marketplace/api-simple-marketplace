package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type SearchResult struct {
	ID   uint
	Name string
}

type SearchProductsHandler struct{}

func (h SearchProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Should Use Get Method", http.StatusMethodNotAllowed)
		return
	}
	queryParam := r.URL.Query().Get("name")
	if queryParam == "" {
		fmt.Fprintf(w, "Query parameter 'name' is missing!")
		return
	}

	esClient := r.Context().Value("es").(*elasticsearch.Client)

	query := fmt.Sprintf(`{"query":{"match":{"name":{"query":"%s", "fuzziness": "0", "fuzzy_transpositions": true}}},"_source":["name"]}`, queryParam)

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

	searchResults := []SearchResult{}

	// Extracting hits
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		hitID := hit.(map[string]interface{})["_id"]
		hitName := hit.(map[string]interface{})["_source"].(map[string]interface{})["name"]

		stringID, _ := hitID.(string)
		id64, _ := strconv.ParseUint(stringID, 10, 0)
		id := uint(id64)
		name, _ := hitName.(string)

		searchResults = append(searchResults, SearchResult{
			ID:   id,
			Name: name,
		})
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searchResults)

}
