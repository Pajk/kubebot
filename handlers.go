package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

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
