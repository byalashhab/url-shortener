package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

/*
   two endpoints
     1) POST /api/v1/shorten -> generte
     2) GET /api/v1/{hash} -> redirects 301
*/

type Server struct {
	port    string
	storage Storage
}

func NewServer(port string, storage Storage) *Server {
	return &Server{port: port, storage: storage}
}

func (s *Server) Run() {
	// start the server with the endpoints
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/shorten", s.HandleShortURLs)
	router.HandleFunc("/api/v1/{id}", s.HandleGetLongURLs)

	log.Println("JSON API server running on port: ", s.port)

	http.ListenAndServe(s.port, router)
}

func (s *Server) HandleShortURLs(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		WriteJSON(w, 405, "Method not supported")
		return
	}

	WriteJSON(w, 200, "Hello from generation route")
}

func (s *Server) HandleGetLongURLs(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		WriteJSON(w, 405, "Method not supported")
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	WriteJSON(w, 200, "Hello from returning route, "+id)
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
