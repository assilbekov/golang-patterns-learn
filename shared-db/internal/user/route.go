package user

import (
	"encoding/json"
	"net/http"
	"shared-db/db"
	"strconv"
)

func (u *User) InitRoute(mux *http.ServeMux) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			u.GetUsers(w, r)
		}
		if r.Method == http.MethodPost {
			u.CreateUser(w, r)
		}
	})

	mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			u.GetUser(w, r)
		}
		if r.Method == http.MethodPut {
			u.UpdateUser(w, r)
		}
	})
}

func (u *User) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.db.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	user, err := u.db.GetUser(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &db.User{}
	json.NewDecoder(r.Body).Decode(user)
	err := u.db.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u *User) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &db.User{}
	json.NewDecoder(r.Body).Decode(user)
	u.db.UpdateUser(user)
}
