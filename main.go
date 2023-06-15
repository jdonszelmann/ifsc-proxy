package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Handler struct{}

func (*Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	resp, err := http.Get(fmt.Sprintf("https://components.ifsc-climbing.org/results-api.php?api=event_full_results&result_url=/api/v1/%s", path))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(buf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &Handler{},
	}
	log.Fatal(server.ListenAndServe())
}