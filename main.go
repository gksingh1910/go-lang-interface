package main

import (
	"fmt"
	"net/http"

	"github.com/gksingh1910/go-lang-interface/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Printf("Let's test patience")
	config.Init()
	fmt.Printf("Test for configuration host is %v", config.Get().ServerConfig.Host)
	rc := chi.NewRouter()

	srv := http.Server{
		Addr:    fmt.Sprintf(config.Get().ServerConfig.Host, ':', config.Get().ServerConfig.Port),
		Handler: rc,
	}

	rc.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	srv.ListenAndServe()

}
