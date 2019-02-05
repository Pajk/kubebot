package main

import (
	"os"
	"net/http"
	"log"
)

func frontendHandler(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Kubebot"))
	log.Print("Request processed")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	var b []byte
	b = append([]byte(""), os.Getenv(appVersion)...)
	w.Write(b)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Alive!"))
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Ready!"))

}
