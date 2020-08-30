package micro

import "net/http"

type Engine interface {
	Run()
	RunTLS()
	Register()
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
