package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleRequest(path string, data map[string]interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == path {
			json.NewEncoder(w).Encode(data)
			w.Header().Set("Content-Type", "application/json")
		} else {
			http.NotFound(w, r)
		}
	}
}

func StartServer(port string, path string, data map[string]interface{}) error {
	handler := handleRequest(path, data)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: handler,
	}
	return srv.ListenAndServe()
}
