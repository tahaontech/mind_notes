package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tahaontech/mind_notes/db"
	"github.com/tahaontech/mind_notes/types"
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
	api := e.Group("/api")
	api.GET("roots", s.handleGetRoots)
	e.Logger.Fatal(e.Start(s.addr))
}

func (s *Server) handleGetRoots(c echo.Context) error {

	return c.JSON(http.StatusOK, types.OkResp{Msg: "ok babe :)"})
}
