package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/adarsh-jaiss/microservice-toll-calculator/types"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "server listen address of http transport server")
	flag.Parse()

	var (
		store = NewMemoryStore()
		svc   = NewInvoiceAggregator(store)
	)
	svc = NewLogMiddleware(svc)
	makeHTTPTransport(*listenAddr, svc)

}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Printf("HTTP Transport running at port %s...\n", listenAddr)
	http.HandleFunc("/aggregate", HandleAggregate(svc))
	http.HandleFunc("/invoice", HandleGetInvoice(svc))
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func HandleGetInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// fmt.Println(r.URL.Query())
		values, ok := r.URL.Query()["obu"]
		if !ok {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "missing OBU ID"})
			return
		}
		// obuId := r.URL.Query()["obu"][0]
		fmt.Println(values[0])
		obuId, err := strconv.Atoi(values[0])
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "inavlid OBU ID "})
		}
		fmt.Println(obuId)

		invoice, err := svc.CalculateInvoice(obuId)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		WriteJSON(w,http.StatusOK,invoice)


		// distance, err := svc.DistanceSum(obuId)
		// if err != nil {
		// 	WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		// }

		// WriteJSON(w, http.StatusOK, map[string]any{"distance"})
		// w.Write([]byte("need to return the invoice for the OBU ID"))
	}
}

func HandleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&Distance); err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			fmt.Fprintf(w, "Error decoding JSON: %v", err)
			return
		}

		if err := svc.AggregateDistance(Distance); err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}

func WriteJSON(w http.ResponseWriter, httpStatus int, v any) error {
	w.WriteHeader(httpStatus)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
