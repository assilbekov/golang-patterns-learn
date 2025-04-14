package tracker

import (
	"net/http"
)



func GetTracker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

func InitRoute(mux *http.ServeMux) {
	mux.HandleFunc("/tracker", GetTracker)
}
