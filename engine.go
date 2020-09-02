package micro

import "net/http"

type Engine interface {
	Register() error
	RunHTTP(addr string) *Service
	RunRPC(addr string) *Service
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
