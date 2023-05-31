package router

import (
	"net/http"
	"receipt-processor/internal/process_receipt"

	"github.com/gorilla/mux"
)

const defaultPort string = ":8000"

func SetupRouter() (router *mux.Router) {
	r := mux.NewRouter()
	paths(r)
	http.ListenAndServe(defaultPort, r)
	return router
}

func paths(r *mux.Router) {
	r.HandleFunc("/receipts/process", process_receipt.Process)
	r.HandleFunc("/receipts/{id}/points", process_receipt.Score)
}
