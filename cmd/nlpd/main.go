package main

import (
	"encoding/json"
	"fmt"
	nlp "github.com/Ferum-Bot/GoPracticeNLP"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func healthHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "OK")
}

func tokenizeHandler(response http.ResponseWriter, request *http.Request) {
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	stringBody := string(requestBody)
	tokens := nlp.Tokenize(stringBody)

	responseBody := struct {
		Tokens []string `json:"tokens"`
	}{
		Tokens: tokens,
	}

	responseData, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(responseData)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	//http.HandleFunc("/health", healthHandler)
	//http.HandleFunc("/tokenize", tokenizeHandler)

	router := mux.NewRouter()
	router.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	router.HandleFunc("/tokenize", tokenizeHandler).Methods(http.MethodPost)
	http.Handle("/", router)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
