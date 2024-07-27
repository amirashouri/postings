package api

import (
	"net/http"

	db "main/db/sqlc"
	"main/util"
)

type Server struct {
	config  util.Config
	store   db.Store
	router  *http.ServeMux
	Service *http.Server
}

// NewServer creates a new HTTP server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("public"))

	router.Handle("/public/*", http.StripPrefix("/public/", fileServer))

	router.HandleFunc("/", server.createUser)

	router.HandleFunc("POST /users", server.createUser)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	service := &http.Server{
		Addr:    address,
		Handler: server.router,
	}
	server.Service = service
	return service.ListenAndServe()
}
