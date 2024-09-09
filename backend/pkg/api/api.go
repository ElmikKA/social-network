package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

type APIServer struct {
	addr     string
	db       *sql.DB
	username string
	id       int
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (api *APIServer) Run() error {
	fmt.Printf("listening on: http://localhost%v\n", api.addr)
	return http.ListenAndServe(api.addr, api.Routes())
}
