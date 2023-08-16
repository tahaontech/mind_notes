package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gonanoid "github.com/matoous/go-nanoid/v2"
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
	// setup logger middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_unix_milli}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	// setup CORS middleware
	e.Use(middleware.CORS())

	// frontend
	e.Static("/", "UI")
	// images
	e.Static("/images", "data/images")

	// API's
	api := e.Group("/api")
	// mindmap
	api.GET("/roots", s.handleGetRoots)
	api.GET("/mindmap/:rootId", s.handleGetMindMap)

	api.POST("/root", s.handleCreateRootNode)
	api.POST("/node", s.handleAddNode)

	api.PATCH("/nodelabel", s.handleUpdateNodeLabel)

	api.DELETE("/node/:nodeId", s.handleDeleteNode)
	api.DELETE("/root/:rootId", s.handleDeleteRoot) // TODO

	// document
	api.GET("/document/:nodeId", s.handleGetDocument)
	api.POST("/document", s.handleCreateDocument)
	api.PATCH("/document", s.handleUpdateDocument)

	// start server
	e.Logger.Fatal(e.Start(s.addr))
}

// GET

func (s *Server) handleGetRoots(c echo.Context) error {
	resp, err := s.db.NodeGetRoots()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResp{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) handleGetMindMap(c echo.Context) error {
	rootId := c.Param("rootId")

	root, err := s.db.NodeGetOne(rootId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResp{Error: err.Error()})
	}

	nodes, err := s.db.NodeGetMany(rootId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResp{Error: err.Error()})
	}

	edges, err := s.db.EdgeGetMany(rootId)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResp{Error: err.Error()})
	}

	resp := types.MindMapResp{
		Category: root.Label,
		Nodes:    nodes,
		Edges:    edges,
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) handleGetDocument(c echo.Context) error {
	nodeId := c.Param("nodeId")

	doc, err := s.db.DocumentGetOne(nodeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResp{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, doc)
}

// POSTS

func (s *Server) handleCreateRootNode(c echo.Context) error {
	var body types.CreateNoderootReq
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResp{Error: err.Error()})
	}

	err := s.db.NodeAddRoot(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResp{Error: err.Error()})
	}

	// create docs
	id, _ := gonanoid.New()
	s.db.DocumentAdd(&types.CreateDocumentReq{ID: id, NodeID: body.ID, Data: ""})

	return c.JSON(http.StatusOK, types.OkResp{Msg: "node & edge added successfully"})
}

func (s *Server) handleAddNode(c echo.Context) error {
	var body types.AddNodeReq
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResp{Error: err.Error()})
	}

	err := s.db.NodeAdd(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResp{Error: err.Error()})
	}

	// create docs
	id, _ := gonanoid.New()
	s.db.DocumentAdd(&types.CreateDocumentReq{ID: id, NodeID: body.ID, Data: ""})

	return c.JSON(http.StatusOK, types.OkResp{Msg: "root node created successfully"})
}

func (s *Server) handleCreateDocument(c echo.Context) error {
	var body types.CreateDocumentReq
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResp{Error: err.Error()})
	}

	err := s.db.DocumentAdd(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResp{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, types.OkResp{Msg: "docs created successfully"})
}

// PATCH

func (s *Server) handleUpdateNodeLabel(c echo.Context) error {
	var body types.UpdateNodeLabelReq
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResp{Error: err.Error()})
	}

	err := s.db.NodeLabelUpdate(body.ID, body.Label)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResp{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, types.OkResp{Msg: "node updated successfully"})
}

func (s *Server) handleUpdateDocument(c echo.Context) error {
	var body types.UpdateDocumentReq
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResp{Error: err.Error()})
	}

	err := s.db.DocumentUpdate(&body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResp{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, types.OkResp{Msg: "document updated successfully"})
}

// DELETE
func (s *Server) handleDeleteNode(c echo.Context) error {
	nodeId := c.Param("nodeId")

	// delete node
	if err := s.db.NodeDelete(nodeId); err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResp{Error: fmt.Sprintf("delete node: %s", err.Error())})
	}
	// delete edge
	if err := s.db.EdgeDeleteByNode(nodeId); err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResp{Error: fmt.Sprintf("delete edge: %s", err.Error())})
	}
	// delete duc
	if err := s.db.DocumentDelete(nodeId); err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResp{Error: fmt.Sprintf("delete docs: %s", err.Error())})
	}

	return c.JSON(http.StatusOK, types.OkResp{Msg: "node deletd successfully"})
}

func (s *Server) handleDeleteRoot(c echo.Context) error {
	// TODO
	return c.JSON(http.StatusOK, types.OkResp{Msg: "root deletd successfully"})
}
