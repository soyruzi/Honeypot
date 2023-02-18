package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	logger *log.Logger
)

type Record struct {
	RemoteAddr   string      `json:"IP"`
	Method       string      `json:"method"`
	RequestURI   string      `json:"URL"`
}

func GenerateRecord(r *http.Request) Record {
	data := Record{}
	data.RemoteAddr = r.RemoteAddr
	data.Method = r.Method
	data.RequestURI = r.RequestURI
	r.ParseForm()

	return data
}

func LogRecord(r Record) error {
	r_json, err := json.Marshal(r)
	if err != nil {
		return err
	}
	logger.Println(string(r_json))

	return nil
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	record := GenerateRecord(r)
	if err := LogRecord(record); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "CRSI honeypot")
}

func main() {
	// setup logging
	if logfile, err := os.OpenFile("/var/log/http.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666); err == nil {
		logger = log.New(logfile, "", 0)
	} else {
		log.Fatal(err)
	}

	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}
