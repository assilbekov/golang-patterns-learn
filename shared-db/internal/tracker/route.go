package tracker

import (
	"encoding/json"
	"net/http"
)

func (t *Tracker) GetTracker(w http.ResponseWriter, r *http.Request) {
	user, err := t.db.GetUser(1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (t *Tracker) InitRoute(mux *http.ServeMux) {
	mux.HandleFunc("/tracker", t.GetTracker)
}
