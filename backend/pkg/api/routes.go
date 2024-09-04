package api

import "net/http"

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/register", Register)
	mux.HandleFunc("/api/login", Login)

	return mux
}
