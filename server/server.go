package server

import (
	"fmt"

	"github.com/tahaontech/mind_notes/db"
)

type Server struct {
	db   *db.DB
	addr string
}

func NewServer(db *db.DB, addr string) *Server {
	return &Server{
		db:   db,
		addr: addr,
	}
}

func (s *Server) Start() {
	fmt.Printf("server starting on %s", s.addr)
}
