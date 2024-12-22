package server

import (
	v1 "calc-lms/internal/api/v1"
	"log"
	"net/http"
)


func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", v1.Calculate)
	if err := http.ListenAndServe("", mux); err != nil {
		log.Fatal(err)
	}
}
