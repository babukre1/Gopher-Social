package api

import (
	"database/sql"
	"ecommerce-api/service/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// import (
// 	// "ecommerce-api/service/user"
// 	"log"
// 	// "log"
// )

// // "golang-api-ecommerce/cmd/api"


type APIServer struct {
	addr string
	db   *sql.DB
}


func NewApiServer(addr string, db *sql.DB)*APIServer{
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error{
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Printf("Starting on port %v ", s.addr)
	return http.ListenAndServe(s.addr, router)
}