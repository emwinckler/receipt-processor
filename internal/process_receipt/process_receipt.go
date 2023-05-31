package process_receipt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"receipt-processor/internal/receipt"
	"reflect"
	"strconv"
	"strings"
)

var idIndex = 2

func Process(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Path only supports POST.")
		w.WriteHeader(http.StatusBadRequest)
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Error reading body.")
			w.WriteHeader(http.StatusBadRequest)
		}
		new_receipt := new(receipt.Receipt)
		err = json.Unmarshal(body, &new_receipt)
		if err != nil {
			fmt.Fprintf(w, "The receipt is invalid")
			w.WriteHeader(http.StatusBadRequest)
		} else {
			id := receipt.NextId()

			receipt.SaveReceipt(id, *new_receipt)

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]int{"id": id})
		}
	}
}

func Score(w http.ResponseWriter, r *http.Request) {

	res := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(res[idIndex])
	if err != nil {
		panic(err)
	}

	if reflect.ValueOf(receipt.GetReceipt(id)).IsZero() {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No receipt found for that id\n")
	} else {
		points := receipt.ScoreReceipt(int(id))
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]int{"points": points})
		case "POST":
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Path only supports GET.")
		}
	}

}
