package main

import (
	"encoding/json"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Item struct {
	ShortDescription string
	Price            string
}

type Receipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item
	Total        string
}

type HTTPSessionStorage struct {
	receiptIds map[string]Receipt
	pointsIds  map[string]Points
}

type Id struct {
	Id string `json:"id"`
}

type Points struct {
	Points int `json:"points"`
}

// register receipt with an Id
func (s *HTTPSessionStorage) process(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Should Use Post Method", http.StatusMethodNotAllowed)
		return
	}
	receipt := &Receipt{}
	err := json.NewDecoder(req.Body).Decode(receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUUID := uuid.New().String()
	s.receiptIds[newUUID] = *receipt

	returnId := Id{
		Id: newUUID,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(returnId); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// compute points based on formula in description for receipt
func (s *HTTPSessionStorage) points(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Should Use GET Method", http.StatusMethodNotAllowed)
		return
	}

	receiptId := req.PathValue("id")

	// instead of recomputing points, store in session and return immediately
	pointsId, ok := s.pointsIds[receiptId]
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(pointsId); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	receipt, ok := s.receiptIds[receiptId]

	if !ok {
		http.Error(w, "Id Not Found", http.StatusBadRequest)
		return
	}

	totalPoints := 0

	// add alphanumeric characters in retailer name to point toal
	retailer := receipt.Retailer
	alphaNumericRetailer := regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(retailer, "")
	totalPoints += len(alphaNumericRetailer)

	total := receipt.Total
	cents := total[len(total)-2:]

	// add points if cents = 0 or is a multiple of .25
	switch cents {
	case "00":
		totalPoints += 75
	case "25":
	case "50":
	case "75":
		totalPoints += 25
	}

	// add number of items to point total
	numItems := len(receipt.Items)
	totalPoints += (numItems >> 1) * 5

	// check trimmed strings of length multiple of 3
	for _, item := range receipt.Items {
		trimmedName := strings.TrimSpace(item.ShortDescription)
		if len(trimmedName)%3 == 0 {
			if floatPrice, err := strconv.ParseFloat(item.Price, 64); err == nil {
				totalPoints += int(math.Ceil(floatPrice * 0.2))
			}
		}
	}

	// check if date is odd
	if int(receipt.PurchaseDate[len(receipt.PurchaseDate)-1])%2 != 0 {
		totalPoints += 6
	}

	// convert time to number and check if between 2pm and 4pm
	numericTime, err := strconv.Atoi(receipt.PurchaseTime[:2] + receipt.PurchaseTime[3:])
	if err != nil {
		http.Error(w, "Purchase Time is improperly formatted", http.StatusBadRequest)
		return
	}

	if numericTime > 1400 && numericTime < 1600 {
		totalPoints += 10
	}

	retPoints := Points{
		Points: totalPoints,
	}

	s.pointsIds[receiptId] = retPoints

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(retPoints); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}

func main() {
	_HTTPSessionStorage := HTTPSessionStorage{
		receiptIds: map[string]Receipt{},
		pointsIds:  map[string]Points{},
	}

	http.HandleFunc("/receipts/process", http.HandlerFunc(_HTTPSessionStorage.process))
	http.HandleFunc("/receipts/{id}/points", http.HandlerFunc(_HTTPSessionStorage.points))
	http.ListenAndServe(":8080", nil)
}
