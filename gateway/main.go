package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

type apifunc func(w http.ResponseWriter, r *http.Request) error

func makeApifunc(fn apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "HTTP server listen address")
	flag.Parse()
	http.HandleFunc("/invoice", makeApifunc(HandleGetInvoice))
	logrus.Infof("HTTP Gateway server is running on port %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))

}

func HandleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, map[string]string{"invoice": "Generated successfully!"})
	
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
