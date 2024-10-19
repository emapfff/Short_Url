package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Url struct {
	ID       string `json:"id"`
	Original string `json:"original_url"`
	Short    string `json:"short_url"`
}

var url Url

func GetOriginalUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(url)
	if err != nil {
		log.Fatal(err)
	}
}
