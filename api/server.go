package api

import (
	"context"
	"fmt"
	"net/http"

	db "main/db/sqlc"
	"main/token"
	"main/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	mux        *http.Server
}

// NewServer creates a new HTTP server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("con not create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter(config.HTTPServerAddress)
	return server, nil
}

func (server *Server) setupRouter(address string) {
	router := gin.Default()

	router.Static("/public", "./public")

	router.GET("/", server.indexHandler)
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.login)
	router.GET("/users/login", server.getLogin)
	router.GET("/users/signup", server.signup)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/posts", server.getPosts)
	authRoutes.POST("/posts", server.createPost)
	authRoutes.GET("/home", server.homeHandler)

	mux := &http.Server{
		Addr:    address,
		Handler: router.Handler(),
	}

	server.router = router
	server.mux = mux
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start() error {
	return server.mux.ListenAndServe()
}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.mux.Shutdown(ctx)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
