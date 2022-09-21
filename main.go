package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.NewSource(time.Now().Unix())
}

func main() {
	// routing configuration

	handler := http.NewServeMux()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		headers, _ := json.Marshal(r.Header)
		responseBody := map[string]interface{}{
			"headers": string(headers),
		}
		response, _ := json.Marshal(responseBody)
		w.Header().Add("Content-Type", "application/json")
		w.Write(response)
	})

	handler.HandleFunc("/stubbed-process-1", func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(100) > 20 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("not ok"))
		}
	})

	handler.HandleFunc("/stubbed-process-2", func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(100) > 40 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("not ok"))
		}
	})

	handler.HandleFunc("/stubbed-process-3", func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(100) > 60 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("not ok"))
		}
	})

	handler.HandleFunc("/stubbed-process-4", func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(100) > 80 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("not ok"))
		}
	})

	// healthcheck function

	go func() {
		for {
			targetService := "http://127.0.0.1:3000"
			log.Printf("pinging %s...", targetService)
			req, _ := http.NewRequest("GET", targetService, nil)
			res, _ := http.DefaultClient.Do(req)
			if res.StatusCode != http.StatusOK {
				log.Printf("failed to complete healthcheck (status code: %v)", res.StatusCode)
			}
			responseBody, _ := ioutil.ReadAll(res.Body)
			log.Println("body data: ", string(responseBody))
			<-time.After(time.Second)
		}
	}()

	// start the server

	if err := http.ListenAndServe("0.0.0.0:3000", handler); err != nil {
		log.Printf("failed to start server: %s", err)
	}
}
