package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Router struct{}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	path := req.URL.Path

	if path == "/cadaster" && req.Method == "POST" {
		user := dataHandler(req)
		insert(user)
		response, err := json.Marshal(user)
		if err != nil {
			res.Write([]byte("Found an isue during JSON deploy"))
		}
		res.Write(response)
	}
}

func RunServer() {
	s := http.Server{
		Addr:         ":8080",
		Handler:      Router{},
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
