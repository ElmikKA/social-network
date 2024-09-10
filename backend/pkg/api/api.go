package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"social-network/db"
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

func (api *APIServer) Run() error {
	handlerStore := db.NewStore(api.db)
	handler := NewHandler(handlerStore)
	fmt.Printf("listening on: http://localhost%v\n", api.addr)
	return http.ListenAndServe(api.addr, handler.Routes())
}
