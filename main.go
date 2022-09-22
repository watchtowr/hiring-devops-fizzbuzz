package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

var globalDataStore = map[string]string{}

func init() {
	rand.NewSource(time.Now().Unix())
}

func main() {
	// endpoint routing configuration

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
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("not ok"))
	})

	handler.HandleFunc("/stubbed-process-2", func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(100) > 40 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("not ok"))
	})

	handler.HandleFunc("/stubbed-process-3", func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(100) > 60 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("not ok"))
	})

	handler.HandleFunc("/stubbed-process-4", func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(100) > 80 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("not ok"))
	})

	handler.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("use post"))
			return
		}
		bodyData, _ := ioutil.ReadAll(r.Body)
		requestId := uuid.New().String()
		globalDataStore[requestId] = string(bodyData)
		w.Write([]byte(requestId))
	})

	handler.HandleFunc("/load/", func(w http.ResponseWriter, r *http.Request) {
		// get the part after /load/, which should be the uuid returned from /save-me
		dataId := strings.TrimPrefix(r.URL.Path, "/load/")
		if globalDataInstance, ok := globalDataStore[dataId]; ok {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(globalDataInstance))
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
	})

	// healthcheck function

	go func() {
		for {
			// default to pinging ourselves
			targetService := "http://127.0.0.1:3000"
			fmt.Printf("pinging %s...\n", targetService)
			req, _ := http.NewRequest("GET", targetService, nil)
			res, _ := http.DefaultClient.Do(req)
			if res.StatusCode != http.StatusOK {
				fmt.Printf("failed to complete healthcheck (status code: %v)\n", res.StatusCode)
			}
			responseBody, _ := ioutil.ReadAll(res.Body)
			fmt.Printf("body data: %s\n", string(responseBody))
			<-time.After(time.Second)
		}
	}()

	// start the server

	if err := http.ListenAndServe("0.0.0.0:3000", handler); err != nil {
		fmt.Printf("failed to start server: %s\n", err)
	}
}
