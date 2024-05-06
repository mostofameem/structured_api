package handlers

import (
	"encoding/json"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		userJson, err := json.Marshal(users)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(userJson)
			return
		}
		http.Error(w, "Error converting users to JSON", http.StatusInternalServerError)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
