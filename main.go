package main

import (
	"fmt"
	"net/http"

	"github.com/gksingh1910/go-lang-interface/config"
	"github.com/gksingh1910/go-lang-interface/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	config.Init()
	r := chi.NewRouter()
	addr := fmt.Sprintf("%v:%v", config.Get().Server.Host, config.Get().Server.Port)
	fmt.Printf("\n Address is %v", addr)

	r.HandleFunc("/v1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello first api call")
		w.WriteHeader(200)
		w.Write([]byte("Hello first api call"))
	})

	r.HandleFunc("/createUser", handler.CreateUser)
	r.HandleFunc("/updateUser", handler.UpdateUser)

	serv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	err := serv.ListenAndServe()
	if err != nil {
		fmt.Printf("Server could not be started")
	}

}
