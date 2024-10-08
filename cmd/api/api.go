package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohitsolanki026/econ-go/service/cart"
	"github.com/mohitsolanki026/econ-go/service/order"
	"github.com/mohitsolanki026/econ-go/service/product"
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

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(apiRouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(orderStore,productStore,userStore)
	cartHandler.RegisterRoutes(apiRouter)

	fmt.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
