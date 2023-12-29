package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	for idx, c := range buf {
		if c == '{' {
			buf = buf[idx:]
			break
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(buf)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("listening on %s\n", port)
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: &Handler{},
	}
	log.Fatal(server.ListenAndServe())
}
