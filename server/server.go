package server

import "github.com/gorilla/mux"

type Router struct {
	Router *mux.Router
}

func NewServer(r *Router) mux.Router {
	return *r.Router
}

func SubRouterauth(r *mux.Router) *mux.Router {
	return r.NewRoute().Subrouter()
}
