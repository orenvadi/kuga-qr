package handler

import "net/http"

type Server struct{}

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(
		"/api",
		func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("shit")) },
	)

	return mux
}
