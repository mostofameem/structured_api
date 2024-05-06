package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var users []User

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err == nil {
			users = append(users, user)
			fmt.Fprintf(w, "Registered Successfully: %s", user.Name)
			return
		}
		http.Error(w, "Error converting user to JSON", http.StatusInternalServerError)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
