package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/adarsh-jaiss/microservice-toll-calculator/aggregator/client"
	"github.com/adarsh-jaiss/microservice-toll-calculator/types"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

func main() {
	httpListenAddr := flag.String("httplistenaddr", ":3000", "server listen address of http transport server")
	// httpListenAddr := flag.String("httplistenaddr", "127.0.0.1:3000", "server listen address of http transport server")
	// httpListenAddr := flag.String("httplistenaddr", "[::1]:3000", "server listen address of http transport server")
	grpcListenAddr := flag.String("grpclistenaddr", ":8080", "server listen address of http transport server")
	flag.Parse()

	var (
		store = NewMemoryStore()
		svc   = NewInvoiceAggregator(store)
	)
	svc = NewMetricsMiddleware(svc)
	svc = NewLogMiddleware(svc)

	go func() {
		log.Fatal(makeGRPCTransport(*grpcListenAddr, svc))
	}()

	time.Sleep(time.Second * 2)
	c, err := client.NewGRPCClient(*grpcListenAddr)
	if err != nil {
		log.Fatal(err)

	}
	if err := c.Aggregate(context.Background(), &types.AggregateRequest{
		ObuID: 1098,
		Value: 6472.23,
		Unix:  time.Now().UnixNano(),
	}); err != nil {
		log.Fatal(err)
	}

	// starting my HTTP Server
	log.Fatal(makeHTTPTransport(*httpListenAddr, svc))

}

func makeHTTPTransport(listenAddr string, svc Aggregator) error {
	// var (
	// 	aggMetricHandler = ewHTTPMetricsHandler("aggregate")
	// 	invMetricHandler = ewHTTPMetricsHandler("invoice")
	// 	aggregateHandler = makeHTTPHandlerFunc(aggMetricHandler.instrument(handleAggregate(svc)))
	// 	invoiceHandler   = makeHTTPHandlerFunc(invMetricHandler.instrument(handleGetInvoice(svc)))
	// )
	fmt.Printf("HTTP Transport running at port %s...\n", listenAddr)
	http.HandleFunc("/aggregate", HandleAggregate(svc))
	http.HandleFunc("/invoice", HandleGetInvoice(svc))
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(listenAddr, nil)
}

func makeGRPCTransport(listenAddr string, svc Aggregator) error {
	fmt.Printf("gRPC Transport running at port %s...\n", listenAddr)
	// Make the TCP Listener
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer listen.Close()
	// Create a new gRPC native server with oprtions
	server := grpc.NewServer([]grpc.ServerOption{}...)
	// register our GRPC server implememtation to the gRPC package
	types.RegisterAggregatorServer(server, NewGRPCAggregatorServer(svc))
	return server.Serve(listen)
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

		WriteJSON(w, http.StatusOK, invoice)

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
