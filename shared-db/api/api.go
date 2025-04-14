package api

import (
	"fmt"
	"log"
	"net/http"
	"shared-db/conf"
)

type API struct {
	apiPort string
	server *http.Server
}

	func New(conf *conf.Config, mux *http.ServeMux) *API {
	return &API{
		apiPort: conf.API_PORT,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%s", conf.API_PORT),
			Handler: mux,
		},
	}
}

func (a *API) Start() error {
	log.Printf("Starting HTTP server on %s", a.apiPort)
	return a.server.ListenAndServe()
}
/* 
func (a *API) InitRoute(path string, handler http.HandlerFunc) {
	// Get the current mux from the server
	mux, ok := a.server.Handler.(*http.ServeMux)
	if !ok {
		// If it's not a ServeMux, create a new one
		mux = http.NewServeMux()
	}
	
	// Register the handler with the mux
	mux.HandleFunc(path, handler)
	
	// Update the server's handler
	a.server.Handler = mux
}
 */

func (a *API) Mux() *http.ServeMux {	
	return a.server.Handler.(*http.ServeMux)
}
