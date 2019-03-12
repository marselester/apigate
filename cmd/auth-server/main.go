// auth-server is an HTTP server API that authenticates all API requests of the travel project.
// It performs HTTP Basic Auth for all URLs except /healthz.
package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	apiAddr := flag.String("http", ":8000", "HTTP API address")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		username, password, ok := req.BasicAuth()
		if !ok || username != password {
			w.Header().Set("WWW-Authenticate", "Basic realm=\"travel\"")
			w.WriteHeader(http.StatusUnauthorized)
			log.Printf("%s 401", req.RequestURI)
			return
		}

		w.Header().Set("X-Travel-User", username)
		log.Printf("%s 200", req.RequestURI)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok"))
	})
	log.Fatal(http.ListenAndServe(*apiAddr, nil))
}
