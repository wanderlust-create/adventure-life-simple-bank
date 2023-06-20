package api

import (
	db "github.com/wanderlust-create/adventure-life-simple-bank/db/simple_bank_db"

	"github.com/gin-gonic/gin"
)

// serves all HTTP requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	server.router = router

	return server
}
