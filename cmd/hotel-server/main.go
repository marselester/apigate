// Program hotel-server is an HTTP server API for hotel reservation.
package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Max API response latency in ms added to emulate random response time.
const maxLatency = 200

func main() {
	apiAddr := flag.String("http", ":8000", "HTTP API address")
	// Initialize the default source of uniformly-distributed pseudo-random ints
	// to add latencies to API responses.
	rand.Seed(time.Now().UnixNano())

	r := mux.NewRouter().SkipClean(true)
	r.Use(loggingMiddleware)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.RequestURI)
	})
	r.HandleFunc("/v1/hotels", hotelsHandler).Methods("Get")
	r.HandleFunc("/v1/hotels/bookings/{booking_id}", bookingHandler).Methods("Get")
	r.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok"))
	})
	log.Fatal(http.ListenAndServe(*apiAddr, r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("user %q %s %s", req.Header.Get("X-Travel-User"), req.Method, req.RequestURI)
		next.ServeHTTP(w, req)
	})
}

// hotelsHandler shows the list of available hotels.
func hotelsHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Duration(rand.Intn(maxLatency)) * time.Millisecond)
	w.Write([]byte(`[
    {
        "id": "046d471d-70c7-4595-80cc-266d3e6e07fa",
        "name": "Holiday Inn",
        "price": "35"
    },
    {
        "id": "f183e115-7efb-49c1-b338-ed5265bc8431",
        "name": "Four Seasons",
        "price": "45"
    }
]`))
}

// bookingsHandler shows details of a hotel reservation with id 7b4fc183-ee67-494d-9715-3510c6d8f2ef.
// The remaining responses are 404.
func bookingHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Duration(rand.Intn(maxLatency)) * time.Millisecond)
	vars := mux.Vars(req)
	if vars["booking_id"] != "7b4fc183-ee67-494d-9715-3510c6d8f2ef" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte(`{
    "id": "7b4fc183-ee67-494d-9715-3510c6d8f2ef",
    "hotel_id": "046d471d-70c7-4595-80cc-266d3e6e07fa",
    "status": "confirmed"
}`))
}
