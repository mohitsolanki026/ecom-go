package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohitsolanki026/econ-go/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	router.PathPrefix("/api/v1").Subrouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	// userHandler.RegisterRoutes(router)

    // Register routes on the apiRouter (subrouter)
    userHandler.RegisterRoutes(apiRouter)

    fmt.Println("Starting server on", s.addr)

    return http.ListenAndServe(s.addr, router)
}
