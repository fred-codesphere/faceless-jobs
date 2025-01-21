package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func HandleThisJob(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing looooong request")

	time.Sleep(3 * time.Second)

	data := map[string]string{
		"duration": "17842379 seconds",
		"status":   "done",
		"result":   "Uff, that was a long one!",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/thisJob", HandleThisJob)

	log.Println("Starting server on :8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
