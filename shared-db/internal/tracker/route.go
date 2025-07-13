package tracker

import (
	"encoding/json"
	"net/http"
	"shared-db/db"
)

func (t *Tracker) InitRoute(mux *http.ServeMux) {
	mux.HandleFunc("/tracker", t.GetTracker)
	mux.HandleFunc("/tracker/{userId}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			t.CreateUserWorkHours(w, r)
		}
		if r.Method == http.MethodPut {
			t.UpdateUserWorkHours(w, r)
		}
	})
}

func (t *Tracker) GetTracker(w http.ResponseWriter, r *http.Request) {
	userWorkHours, err := t.db.GetUserWorkHours(1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userWorkHours)
}

func (t *Tracker) CreateUserWorkHours(w http.ResponseWriter, r *http.Request) {
	userWorkHours := &db.UserWorkHours{}
	json.NewDecoder(r.Body).Decode(userWorkHours)
	err := t.db.CreateUserWorkHours(userWorkHours)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userWorkHours)
}

func (t *Tracker) UpdateUserWorkHours(w http.ResponseWriter, r *http.Request) {
	userWorkHours := &db.UserWorkHours{}
	json.NewDecoder(r.Body).Decode(userWorkHours)
	err := t.db.UpdateUserWorkHours(userWorkHours)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userWorkHours)
}
