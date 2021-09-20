package web

import (
	"net/http"
	"fmt"
	"dkv-db/db"
)

// Server contains HTTP method handlers to be used for the database.
type Server struct {
	db *db.Database
}

// NewServer creates a new instance with HTTP handlers to get and set values.
func NewServer(db *db.Database) *Server {
	return &Server{
		db: db,
	}
}

// GetHandler handles read requests
func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	value, err := s.db.GetKey(key)
	fmt.Fprintf(w, "Value = %q, error = %v\n", value, err)
}

// GetHandler handles write requests
func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	value := r.Form.Get("value")
	err := s.db.SetKey(key, []byte(value))
	fmt.Fprintf(w, "Error = %v\n", err)
}
