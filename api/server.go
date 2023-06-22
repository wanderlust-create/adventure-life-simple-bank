package api

import (
	db "github.com/wanderlust-create/adventure-life-simple-bank/db/simple_bank_db"

	"github.com/gin-gonic/gin"
)

// serves all HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router

	return server
}

// run the HTTP server on the input address to start listenng for API requests
// TODO: add graceful shutdown logic
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
