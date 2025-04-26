package main

import (
	"database/sql"
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

	var req struct {
		LongURL string `json:"longURL"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, 500, "server error")
		return
	}

	if req.LongURL == "" {
		WriteJSON(w, 400, "bad request")
		return
	}

	hash, err := s.storage.AddShortURL(req.LongURL)

	if err != nil {
		log.Fatal(err)
	}

	WriteJSON(w, 201, hash)
}

func (s *Server) HandleGetLongURLs(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		WriteJSON(w, 405, "Method not supported")
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	longURL, err := s.storage.GetLongURL(id)

	if err != nil {

		if err == sql.ErrNoRows {
			WriteJSON(w, 404, "not found")
			return
		} else {
			log.Fatal(err)
		}

	}

	WriteJSON(w, 200, longURL)
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
