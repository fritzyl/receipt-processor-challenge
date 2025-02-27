package routes

import (
	"encoding/json"
	"net/http"

	"github.com/fritzyl/receipt-processor-challenge/api/receipts"
	"github.com/fritzyl/receipt-processor-challenge/api/types"
)

func Register(server *http.ServeMux) {
	server.HandleFunc("POST /receipts/process", processReceipt)
	server.HandleFunc("GET /receipts/{id}/points", getPoints)
}

func writeResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt *types.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate Receipt
	errs := receipts.Validate(receipt)

	if errs != nil {
		http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	uuid, prErr := receipts.Process(receipt)

	if prErr != nil {
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}

	writeResponse(w, types.ProcessResponse{Id: uuid.String()})
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	lookupId := r.PathValue("id")
	points, err := receipts.Lookup(lookupId)

	if err != nil || points < 0 {
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	writeResponse(w, types.GetPointsResponse{Points: points})
}
