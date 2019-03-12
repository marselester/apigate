// Program car-server is an HTTP server API for renting cars.
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
	r.HandleFunc("/v1/cars", carsHandler).Methods("Get")
	r.HandleFunc("/v1/cars/bookings/{booking_id}", bookingHandler).Methods("Get")
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

// carsHandler shows the list of available cars.
func carsHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Duration(rand.Intn(maxLatency)) * time.Millisecond)
	w.Write([]byte(`[
    {
        "id": "cfb6f7a5-4591-4f5c-8b17-9a1b10f98ada",
        "name": "Toyota Yaris",
        "price": "30"
    },
    {
        "id": "afad6e6c-ef7f-4dc9-bc0f-ce74d5392175",
        "name": "Honda Civic",
        "price": "40"
    }
]`))
}

// bookingsHandler shows details of a car booking with id 9e0d65f5-9de2-4428-9bee-1f3967f05129.
// The remaining responses are 404.
func bookingHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Duration(rand.Intn(maxLatency)) * time.Millisecond)
	vars := mux.Vars(req)
	if vars["booking_id"] != "9e0d65f5-9de2-4428-9bee-1f3967f05129" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte(`{
    "id": "9e0d65f5-9de2-4428-9bee-1f3967f05129",
    "car_id": "cfb6f7a5-4591-4f5c-8b17-9a1b10f98ada",
    "status": "confirmed"
}`))
}
