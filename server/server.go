package server

import (
	"github.com/labstack/echo/v4"
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
	e := echo.New()

	// frontend
	e.Static("/", "UI")
	// images
	e.Static("/images", "data/images")

	// API's
	e.Logger.Fatal(e.Start(s.addr))
}
